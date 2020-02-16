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

// Package messages defines interface for the protocol messages.
package messages

import (
	"encoding"
)

// MessageImpl provides an implementation of the message representation.
type MessageImpl interface {
	NewFromBinary(data []byte) (Message, error)
	NewRequest(clientID uint32, sequence uint64, operation []byte) Request
	NewPrepare(replicaID uint32, view uint64, request Request) Prepare
	NewCommit(replicaID uint32, prepare Prepare) Commit
	NewReqViewChange(replicaID uint32, newView uint64) ReqViewChange
	NewViewChange(replicaID uint32, reqViewChange ReqViewChange) ViewChange
	NewReply(replicaID, clientID uint32, sequence uint64, result []byte) Reply
}

// Testing only!!
type CheckpointCertificate interface {
	
}

type Message interface {
	encoding.BinaryMarshaler
}

// ClientMessage represents a message generated by a client.
type ClientMessage interface {
	Message
	ClientID() uint32
	ImplementsClientMessage()
}

// ReplicaMessage represents a message generated by a replica.
//
// EmbeddedMessages method returns a sequence of messages embedded
// into this one.
type ReplicaMessage interface {
	Message
	ReplicaID() uint32
	ImplementsReplicaMessage()
}

// PeerMessage represents a message exchanged between replicas.
type PeerMessage interface {
	ReplicaMessage
	ImplementsPeerMessage()
}

// CertifiedMessage represents a message certified with a UI.
type CertifiedMessage interface {
	ReplicaMessage
	UIBytes() []byte
	SetUIBytes(ui []byte)
}

// SignedMessage represents a message signed with a normal signature.
type SignedMessage interface {
	Signature() []byte
	SetSignature(signature []byte)
}

type Request interface {
	ClientMessage
	SignedMessage
	Sequence() uint64
	Operation() []byte
	ImplementsRequest()
}

type Prepare interface {
	CertifiedMessage
	View() uint64
	Request() Request
	ImplementsPeerMessage()
	ImplementsPrepare()
}

type Commit interface {
	CertifiedMessage
	Prepare() Prepare
	ImplementsPeerMessage()
	ImplementsCommit()
}

type Reply interface {
	ReplicaMessage
	SignedMessage
	ClientID() uint32
	Sequence() uint64
	Result() []byte
	ImplementsReply()
}

type ReqViewChange interface {
	ReplicaMessage
	SignedMessage
	View() uint64
	RequestedView() uint64
	ImplementsPeerMessage()
	ImplementsReqViewChange()
}

type ViewChange interface {
	CertifiedMessage
	NewView() uint64
	CheckpointCertificate() CheckpointCertificate
	MessagesSinceCheckpoint() []ReplicaMessage 
	ReqViewChange() ReqViewChange
	ImplementsPeerMessage()
	ImplementsViewChange()
}
