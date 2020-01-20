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
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hyperledger-labs/minbft/messages"
)

func TestReqViewChange(t *testing.T) {
	impl := NewImpl()

	t.Run("Fields", func(t *testing.T) {
		r := rand.Uint32()
		nv := rand.Uint64()
		rvc := impl.NewReqViewChange(r, nv)
		require.Equal(t, r, rvc.ReplicaID())
		require.Equal(t, nv, rvc.NewView())
	})
	t.Run("SignedPayload", func(t *testing.T) {
		rvc := randReqViewChange(impl)
		r := rvc.ReplicaID()
		sp := rvc.SignedPayload()
		require.NotEqual(t, sp, newTestReqViewChange(impl, r, rand.Uint64()).SignedPayload())
	})
	t.Run("SetSignature", func(t *testing.T) {
		rvc := randReqViewChange(impl)
		sig := testSig(rvc.SignedPayload())
		rvc.SetSignature(sig)
		require.Equal(t, sig, rvc.Signature())
	})
	t.Run("Marshaling", func(t *testing.T) {
		rvc := randReqViewChange(impl)
		requireReqViewChangeEqual(t, rvc, remarshalMsg(impl, rvc).(messages.ReqViewChange))
	})
}

func randReqViewChange(impl messages.MessageImpl) messages.ReqViewChange {
	return newTestReqViewChange(impl, rand.Uint32(), rand.Uint64())
}

func newTestReqViewChange(impl messages.MessageImpl, r uint32, nv uint64) messages.ReqViewChange {
	rvc := impl.NewReqViewChange(r, nv)
	rvc.SetSignature(testSig(rvc.SignedPayload()))
	return rvc
}

func requireReqViewChangeEqual(t *testing.T, rvc1, rvc2 messages.ReqViewChange) {
	require.Equal(t, rvc1.ReplicaID(), rvc2.ReplicaID())
	require.Equal(t, rvc1.NewView(), rvc2.NewView())
	require.Equal(t, rvc1.SignedPayload(), rvc2.SignedPayload())
	require.Equal(t, rvc1.Signature(), rvc2.Signature())
}
