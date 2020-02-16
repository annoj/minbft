// Copyright (c) 2018 NEC Laboratories Europe GmbH.
//
// Authors: Wenting Li <wenting.li@neclab.eu>
//          Sergey Fedorov <sergey.fedorov@neclab.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package minbft

import (
	"fmt"
	"sync"

	"github.com/hyperledger-labs/minbft/core/internal/requestlist"
	"github.com/hyperledger-labs/minbft/messages"
)

// newViewValidator validates a NewView message.
//
// It authenticates and checks the supplied message for internal
// consistency. It does not use replica's current state and has no
// side-effect. It is safe to invoke concurrently.
type newViewValidator func(newView messages.NewView) error

// newViewApplier applies NewView message to current replica state.
//
// The supplied message is applied to the current replica state by
// changing the state accordingly and producing any required side
// effects. The supplied message is assumed to be authentic and
// internally consistent. Parameter active indicates if the message
// refers to the active view. It is safe to invoke concurrently.
type newViewApplier func(newView messages.NewView, active bool) error

// newViewmentCollector collects newViewment on prepared Request.
//
// The supplied Prepare message is assumed to be valid and should have
// a UI assigned. Once the threshold of matching newViewments from
// distinct replicas has been reached, it triggers further required
// actions to complete the prepared Request. It is safe to invoke
// concurrently.
type newViewmentCollector func(replicaID uint32, prepare messages.Prepare) error

// newViewmentCounter counts newViewments on prepared Request.
//
// The supplied Prepare message is assumed to be valid and should have
// a UI assigned. The return value done indicates if enough
// newViewments from different replicas are counted for the supplied
// Prepare, such that the threshold to execute the prepared operation
// has been reached. An error is returned if any inconsistency is
// detected.
type newViewmentCounter func(replicaID uint32, prepare messages.Prepare) (done bool, err error)

// makeNewViewValidator constructs an instance of newViewValidator using
// the supplied abstractions.
func makeNewViewValidator(verifyUI uiVerifier, validatePrepare prepareValidator) newViewValidator {
	return func(newView messages.NewView) error {

		_ = newView

		return nil
	}
}

// makeNewViewApplier constructs an instance of newViewApplier using the
// supplied abstractions.
func makeNewViewApplier(collectNewViewment newViewmentCollector) newViewApplier {
	return func(newView messages.NewView, active bool) error {

		_ = newView

		return nil
	}
}

// makeNewViewmentCollector constructs an instance of
// newViewmentCollector using the supplied abstractions.
func makeNewViewmentCollector(countNewViewment newViewmentCounter, retireSeq requestSeqRetirer, pendingReq requestlist.List, stopReqTimer requestTimerStopper, executeRequest requestExecutor) newViewmentCollector {
	var lock sync.Mutex

	return func(replicaID uint32, prepare messages.Prepare) error {
		lock.Lock()
		defer lock.Unlock()

		if done, err := countNewViewment(replicaID, prepare); err != nil {
			return err
		} else if !done {
			return nil
		}

		request := prepare.Request()

		if new := retireSeq(request); !new {
			return nil // request already accepted for execution
		}

		pendingReq.Remove(request.ClientID())
		stopReqTimer(request)
		executeRequest(request)

		return nil
	}
}

// makeNewViewmentCounter constructs an instance of newViewmentCounter
// given the number of tolerated faulty nodes.
func makeNewViewmentCounter(f uint32) newViewmentCounter {
	// Replica ID -> newViewted
	type replicasNewViewtedMap map[uint32]bool

	var (
		// Current view number
		view uint64

		// UI counter of the first Prepare in the view
		firstCV uint64 = 1

		// Primary UI counter of the last quorum
		lastDoneCV uint64

		// Primary UI counter -> replicasNewViewtedMap
		prepareStates = make(map[uint64]replicasNewViewtedMap)
	)

	return func(replicaID uint32, prepare messages.Prepare) (done bool, err error) {
		primaryID := prepare.ReplicaID()
		prepareView := prepare.View()
		prepareUI, err := parseMessageUI(prepare)
		if err != nil {
			panic(err)
		}
		prepareCV := prepareUI.Counter

		if prepareView < view {
			return false, nil
		} else if prepareView > view {
			view = prepareView
			firstCV = prepareCV
			lastDoneCV = 0
			prepareStates = make(map[uint64]replicasNewViewtedMap)
		}

		if prepareCV <= lastDoneCV {
			return true, nil
		}

		replicasNewViewted := prepareStates[prepareCV]
		if replicasNewViewted == nil {
			replicasNewViewted = replicasNewViewtedMap{
				primaryID: true,
			}
			prepareStates[prepareCV] = replicasNewViewted
		}

		if replicaID != primaryID {
			for cv := prepareCV - 1; cv >= firstCV; cv-- {
				s := prepareStates[cv]
				if s[replicaID] {
					break
				} else if s != nil {
					return false, fmt.Errorf("Skipped newViewment")
				}
			}

			if replicasNewViewted[replicaID] {
				return false, fmt.Errorf("Duplicated newViewment")
			}

			replicasNewViewted[replicaID] = true
		}

		if len(replicasNewViewted) <= int(f) {
			return false, nil
		}

		lastDoneCV = prepareCV
		delete(prepareStates, prepareCV)

		return true, nil
	}
}
