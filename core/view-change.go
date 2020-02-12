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

	"github.com/hyperledger-labs/minbft/messages"
)

// viewChangeValidator validates a ViewChange message.
//
// It authenticates and checks the supplied message for internal
// consistency. It does not use replica's current state and has no
// side-effect. It is safe to invoke concurrently.
type viewChangeValidator func(viewChange messages.ViewChange) error

// viewChangeApplier applies ViewChange message to current replica state.
//
// The supplied message is applied to the current replica state by
// changing the state accordingly and producing any required side
// effects. The supplied message is assumed to be authentic and
// internally consistent. Parameter active indicates if the message
// refers to the active view. It is safe to invoke concurrently.
type viewChangeApplier func(viewChange messages.ViewChange, active bool) error

// reqViewChangeCollector collects viewChange on reqViewChange.
//
// The supplied ReqViewChange message is assumed to be valid and should have
// a UI assigned. Once the threshold of matching viewChanges from
// distinct replicas has been reached, it triggers further required
// actions to complete the viewChange. It is safe to invoke
// concurrently.
type reqViewChangeCollector func(replicaID uint32) error

// viewChangeCounter counts viewChanges on reqViewChange.
//
// The supplied ReqViewChange message is assumed to be valid and should have
// a UI assigned. The return value done indicates if enough
// viewChanges from different replicas are counted for the supplied
// ReqViewChange, such that the threshold to execute the viewChange operation
// has been reached. An error is returned if any inconsistency is
// detected.
type reqViewChangeCounter func(replicaID uint32) (done bool, err error)

// makeViewChangeValidator constructs an instance of viewChangeValidator using
// the supplied abstractions.
func makeViewChangeValidator(verifyUI uiVerifier, validateReqViewChange reqViewChangeValidator) viewChangeValidator {
	return func(viewChange messages.ViewChange) error {

		fmt.Println("viewChangeValidator was invoked.")
		_ =verifyUI 
		_ = validateReqViewChange

		/*
		prepare := viewChange.Prepare()

		if viewChange.ReplicaID() == prepare.ReplicaID() {
			return fmt.Errorf("ViewChange from primary")
		}

		if err := validatePrepare(prepare); err != nil {
			return fmt.Errorf("Invalid Prepare: %s", err)
		}

		if _, err := verifyUI(viewChange); err != nil {
			return fmt.Errorf("UI is not valid: %s", err)
		}
		*/

		return nil
	}
}

// makeViewChangeApplier constructs an instance of viewChangeApplier using the
// supplied abstractions.
func makeViewChangeApplier(collectViewChange reqViewChangeCollector) viewChangeApplier {
	return func(viewChange messages.ViewChange, active bool) error {

		fmt.Println("viewChangeApplier was invoked.")
		replicaID := viewChange.ReplicaID()
		_ = active

		/*
		replicaID := viewChange.ReplicaID()
		prepare := viewChange.Prepare()
		*/

		if err := collectViewChange(replicaID); err != nil {
			return fmt.Errorf("ViewChange cannot be taken into account: %s", err)
		}

		return nil
	}
}

// makeViewChangeCollector constructs an instance of
// reqViewChangeCollector using the supplied abstractions.
func makeReqViewChangeCollector(countReqViewChange reqViewChangeCounter) reqViewChangeCollector {
	var lock sync.Mutex

	return func(replicaID uint32) error {

		fmt.Println("reqViewChangeCollector was invoked.")

		lock.Lock()
		defer lock.Unlock()

		if done, err := countReqViewChange(replicaID); err != nil {
			return err
		} else if !done {
			return nil
		}

		return nil
	}
}

// makeViewChangeCounter constructs an instance of viewChangeCounter
// given the number of tolerated faulty nodes.
func makeReqViewChangeCounter(f uint32) reqViewChangeCounter {
	return func(replicaID uint32) (done bool, err error) {
		return true, nil
	}
}
