// Copyright (c) 2020 NEC Laboratories Europe GmbH.
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

package protobuf

import (
	"github.com/hyperledger-labs/minbft/messages"
	"github.com/hyperledger-labs/minbft/messages/protobuf/pb"
)

type viewChange struct {
	pbMsg *pb.ViewChange
}

func newViewChange(r uint32, reqViewChange messages.ReqViewChange) *viewChange {
	return &viewChange{pbMsg: &pb.ViewChange{
		ReplicaId:		r,
		ReqViewChange:	pbReqViewChangeFromAPI(reqViewChange),
	}}
}

func newViewChangeFromPb(pbMsg *pb.ViewChange) *viewChange {
	return &viewChange{pbMsg: pbMsg}
}

func (m *viewChange) MarshalBinary() ([]byte, error) {
	return marshalMessage(m.pbMsg)
}

func (m *viewChange) ReplicaID() uint32 {
	return m.pbMsg.GetReplicaId()
}

func (m *viewChange) ReqViewChange() messages.ReqViewChange {
	return nil // m.pbMsg.GetReqViewChange()
}

func (m *viewChange) UIBytes() []byte {
	return nil // m.pbMsg.Ui
}

func (m *viewChange) SetUIBytes(uiBytes []byte) {
	// m.pbMsg.Ui = uiBytes
}

func (viewChange) ImplementsReplicaMessage() {}
func (viewChange) ImplementsPeerMessage()    {}
func (viewChange) ImplementsViewChange() 	 {}
