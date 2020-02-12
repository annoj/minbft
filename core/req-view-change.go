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
// TODO: View numbers have also to be checked for consistency.
func makeReqViewChangeValidator(n uint32, verifyUI uiVerifier) reqViewChangeValidator {
		fmt.Println("makeReqViewChangeValidator was called!")
	return func(reqViewChange messages.ReqViewChange) error {
		replicaID := reqViewChange.ReplicaID()
		newView := reqViewChange.NewView()

		_ = newView
		_ = replicaID
		_ = n
		_ = verifyUI

		// Check REQ VIEW CHANGE message is not coming from the current primary
		// if isPrimary(newView, replicaID, n) {
		// 	return fmt.Errorf("ReqViewChange from primary %d for new view %d", replicaID, newView)
		// }

		// TODO: Should be validateRequest RequestValidator??
		// if err := validateRequest(reqViewChange.Request()); err != nil {
		// 	return fmt.Errorf("Request invalid: %s", err)
		// }

		// TODO: Is this actually a message supposed to be signed?
		// if _, err := verifyUI(reqViewChange); err != nil {
		// 	return fmt.Errorf("UI not valid: %s", err)
		// }

		return nil
	}
}

// makeReqViewChangeApplier constructs an instance of reqViewChangeApplier using
// id as the current replica ID, and the supplied abstract interfaces.
func makeReqViewChangeApplier(id uint32, handleGeneratedMessage generatedMessageHandler) reqViewChangeApplier {
	return func(reqViewChange messages.ReqViewChange) error {
		newPrimaryID := reqViewChange.ReplicaID()
		_ = newPrimaryID

		// TODO: Stop reqViewChangeTimer here!

		handleGeneratedMessage(messageImpl.NewViewChange(id, reqViewChange))

		return nil
	}
}
