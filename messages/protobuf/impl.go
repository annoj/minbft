// Copyright (c) 2019 NEC Laboratories Europe GmbH.
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
	"golang.org/x/xerrors"

	"github.com/golang/protobuf/proto"

	"github.com/hyperledger-labs/minbft/messages"
	"github.com/hyperledger-labs/minbft/messages/protobuf/pb"
)

type impl struct{}

func NewImpl() messages.MessageImpl {
	return &impl{}
}

func (*impl) NewFromBinary(data []byte) (messages.Message, error) {
	msg := &pb.Message{}
	if err := proto.Unmarshal(data, msg); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal message wrapper: %w", err)
	}

	switch t := msg.Type.(type) {
	case *pb.Message_Request:
		req := newRequest()
		req.set(t.Request)
		return req, nil
	case *pb.Message_Prepare:
		prep := newPrepare()
		prep.set(t.Prepare)
		return prep, nil
	case *pb.Message_Commit:
		comm := newCommit()
		comm.set(t.Commit)
		return comm, nil
	case *pb.Message_Reply:
		reply := newReply()
		reply.set(t.Reply)
		return reply, nil
	default:
		return nil, xerrors.New("unknown message type")
	}
}

func (*impl) NewRequest(cl uint32, seq uint64, op []byte) messages.Request {
	m := newRequest()
	m.init(cl, seq, op)
	return m
}

func (*impl) NewPrepare(r uint32, v uint64, req messages.Request) messages.Prepare {
	m := newPrepare()
	m.init(r, v, req)
	return m
}

func (*impl) NewCommit(r uint32, prep messages.Prepare) messages.Commit {
	m := newCommit()
	m.init(r, prep)
	return m
}

func (*impl) NewReply(r, cl uint32, seq uint64, res []byte) messages.Reply {
	m := newReply()
	m.init(r, cl, seq, res)
	return m
}
