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

type reqViewChange struct {
	pbMsg *pb.ReqViewChange
}

func newReqViewChange(r uint32, rv uint64) *reqViewChange {
	return &reqViewChange{pbMsg: &pb.ReqViewChange{
		ReplicaId: 		r,
		RequestedView:  rv,
	}}
}

func newReqViewChangeFromPb(pbMsg *pb.ReqViewChange) *reqViewChange {
	return &reqViewChange{pbMsg: pbMsg}
}

func (m *reqViewChange) MarshalBinary() ([]byte, error) {
	return marshalMessage(m.pbMsg)
}

func (m *reqViewChange) ReplicaID() uint32 {
	return m.pbMsg.GetReplicaId()
}

func (m *reqViewChange) View() uint64 {
	return m.pbMsg.GetView()
}

func (m *reqViewChange) RequestedView() uint64 {
	return m.pbMsg.GetRequestedView()
}

func (m *reqViewChange) Signature() []byte {
	return m.pbMsg.Signature
}

func (m *reqViewChange) SetSignature(signature []byte) {
	m.pbMsg.Signature = signature
}

func (reqViewChange) ImplementsReplicaMessage() {}
func (reqViewChange) ImplementsPeerMessage()    {}
func (reqViewChange) ImplementsReqViewChange()  {}

func pbReqViewChangeFromAPI(m messages.ReqViewChange) *pb.ReqViewChange{
	if m, ok := m.(*reqViewChange); ok {
		return m.pbMsg
	}

	return pb.ReqViewChangeFromAPI(m)
}
