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

	"github.com/hyperledger-labs/minbft/messages"
	"github.com/hyperledger-labs/minbft/core/internal/viewstate"
)

// reqViewChangeValidator validates a ReqViewChange message.
//
// It authenticates and checks the supplied message for internal
// consistency. It does not use replica's current state and has no
// side-effect. It is safe to invoke concurrently.
type reqViewChangeValidator func(reqViewChange messages.ReqViewChange) error

// reqViewChangeApplier applies ReqViewChange message to current replica state.
//
// The supplied message is applied to the current replica state by
// changing the state accordingly and producing any required messages
// or side effects. The supplied message is assumed to be authentic
// and internally consistent. Parameter active indicates if the
// message refers to the active view. It is safe to invoke
// concurrently.
type reqViewChangeApplier func(reqViewChange messages.ReqViewChange) error

// reqViewChangeValidator constructs an instance of reqViewChangeValidator
// using n as the total number of nodes, and the supplied abstract
// interfaces.
func makeReqViewChangeValidator(n uint32, viewState viewstate.State) reqViewChangeValidator {
		fmt.Println("makeReqViewChangeValidator was called!")
	return func(reqViewChange messages.ReqViewChange) error {
		replicaID := reqViewChange.ReplicaID()
		requestedView:= reqViewChange.RequestedView()
		currentView := reqViewChange.View()

		if isPrimary(currentView, replicaID, n) {
			return fmt.Errorf("ReqViewChange from Primary %d", replicaID)
		}
		
		currentView, _, release := viewState.HoldView()
		defer release()
		
		if currentView <= requestedView {
			return fmt.Errorf("Number of requested View was invalid, currentView: %d, requestedView: %d", currentView, requestedView)
		}

		return nil
	}
}

// makeReqViewChangeApplier constructs an instance of reqViewChangeApplier using
// id as the current replica ID, and the supplied abstract interfaces.
func makeReqViewChangeApplier(id uint32, collectReqViewChange reqViewChangeCollector, handleGeneratedMessage generatedMessageHandler) reqViewChangeApplier {
	return func(reqViewChange messages.ReqViewChange) error {

		requestedView := reqViewChange.RequestedView()

		// Increase expectedView in viewState
		ok, release := viewState.AdvanceExpectedView(requestedView)
		if !ok {
			return fmt.Errorf("ExpectedView could not be increased to %d", requestedView)
		}
		defer release()

		// TODO: Stop reqViewChangeTimer here!

		if err := collectReqViewChange(newPrimaryID); err != nil {
			return fmt.Errorf("ReqViewChange cannot be taken into account: %s", err)
		}

		handleGeneratedMessage(messageImpl.NewViewChange(id, reqViewChange))

		return nil
	}
}
