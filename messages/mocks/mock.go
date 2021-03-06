// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger-labs/minbft/messages (interfaces: Message,ClientMessage,ReplicaMessage,PeerMessage,CertifiedMessage,SignedMessage)

// Package mock_messages is a generated GoMock package.
package mock_messages

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMessage is a mock of Message interface
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// MarshalBinary mocks base method
func (m *MockMessage) MarshalBinary() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalBinary")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalBinary indicates an expected call of MarshalBinary
func (mr *MockMessageMockRecorder) MarshalBinary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalBinary", reflect.TypeOf((*MockMessage)(nil).MarshalBinary))
}

// MockClientMessage is a mock of ClientMessage interface
type MockClientMessage struct {
	ctrl     *gomock.Controller
	recorder *MockClientMessageMockRecorder
}

// MockClientMessageMockRecorder is the mock recorder for MockClientMessage
type MockClientMessageMockRecorder struct {
	mock *MockClientMessage
}

// NewMockClientMessage creates a new mock instance
func NewMockClientMessage(ctrl *gomock.Controller) *MockClientMessage {
	mock := &MockClientMessage{ctrl: ctrl}
	mock.recorder = &MockClientMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientMessage) EXPECT() *MockClientMessageMockRecorder {
	return m.recorder
}

// ClientID mocks base method
func (m *MockClientMessage) ClientID() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// ClientID indicates an expected call of ClientID
func (mr *MockClientMessageMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockClientMessage)(nil).ClientID))
}

// ImplementsClientMessage mocks base method
func (m *MockClientMessage) ImplementsClientMessage() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ImplementsClientMessage")
}

// ImplementsClientMessage indicates an expected call of ImplementsClientMessage
func (mr *MockClientMessageMockRecorder) ImplementsClientMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImplementsClientMessage", reflect.TypeOf((*MockClientMessage)(nil).ImplementsClientMessage))
}

// MarshalBinary mocks base method
func (m *MockClientMessage) MarshalBinary() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalBinary")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalBinary indicates an expected call of MarshalBinary
func (mr *MockClientMessageMockRecorder) MarshalBinary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalBinary", reflect.TypeOf((*MockClientMessage)(nil).MarshalBinary))
}

// MockReplicaMessage is a mock of ReplicaMessage interface
type MockReplicaMessage struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaMessageMockRecorder
}

// MockReplicaMessageMockRecorder is the mock recorder for MockReplicaMessage
type MockReplicaMessageMockRecorder struct {
	mock *MockReplicaMessage
}

// NewMockReplicaMessage creates a new mock instance
func NewMockReplicaMessage(ctrl *gomock.Controller) *MockReplicaMessage {
	mock := &MockReplicaMessage{ctrl: ctrl}
	mock.recorder = &MockReplicaMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReplicaMessage) EXPECT() *MockReplicaMessageMockRecorder {
	return m.recorder
}

// ImplementsReplicaMessage mocks base method
func (m *MockReplicaMessage) ImplementsReplicaMessage() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ImplementsReplicaMessage")
}

// ImplementsReplicaMessage indicates an expected call of ImplementsReplicaMessage
func (mr *MockReplicaMessageMockRecorder) ImplementsReplicaMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImplementsReplicaMessage", reflect.TypeOf((*MockReplicaMessage)(nil).ImplementsReplicaMessage))
}

// MarshalBinary mocks base method
func (m *MockReplicaMessage) MarshalBinary() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalBinary")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalBinary indicates an expected call of MarshalBinary
func (mr *MockReplicaMessageMockRecorder) MarshalBinary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalBinary", reflect.TypeOf((*MockReplicaMessage)(nil).MarshalBinary))
}

// ReplicaID mocks base method
func (m *MockReplicaMessage) ReplicaID() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicaID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// ReplicaID indicates an expected call of ReplicaID
func (mr *MockReplicaMessageMockRecorder) ReplicaID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaID", reflect.TypeOf((*MockReplicaMessage)(nil).ReplicaID))
}

// MockPeerMessage is a mock of PeerMessage interface
type MockPeerMessage struct {
	ctrl     *gomock.Controller
	recorder *MockPeerMessageMockRecorder
}

// MockPeerMessageMockRecorder is the mock recorder for MockPeerMessage
type MockPeerMessageMockRecorder struct {
	mock *MockPeerMessage
}

// NewMockPeerMessage creates a new mock instance
func NewMockPeerMessage(ctrl *gomock.Controller) *MockPeerMessage {
	mock := &MockPeerMessage{ctrl: ctrl}
	mock.recorder = &MockPeerMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeerMessage) EXPECT() *MockPeerMessageMockRecorder {
	return m.recorder
}

// ImplementsPeerMessage mocks base method
func (m *MockPeerMessage) ImplementsPeerMessage() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ImplementsPeerMessage")
}

// ImplementsPeerMessage indicates an expected call of ImplementsPeerMessage
func (mr *MockPeerMessageMockRecorder) ImplementsPeerMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImplementsPeerMessage", reflect.TypeOf((*MockPeerMessage)(nil).ImplementsPeerMessage))
}

// ImplementsReplicaMessage mocks base method
func (m *MockPeerMessage) ImplementsReplicaMessage() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ImplementsReplicaMessage")
}

// ImplementsReplicaMessage indicates an expected call of ImplementsReplicaMessage
func (mr *MockPeerMessageMockRecorder) ImplementsReplicaMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImplementsReplicaMessage", reflect.TypeOf((*MockPeerMessage)(nil).ImplementsReplicaMessage))
}

// MarshalBinary mocks base method
func (m *MockPeerMessage) MarshalBinary() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalBinary")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalBinary indicates an expected call of MarshalBinary
func (mr *MockPeerMessageMockRecorder) MarshalBinary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalBinary", reflect.TypeOf((*MockPeerMessage)(nil).MarshalBinary))
}

// ReplicaID mocks base method
func (m *MockPeerMessage) ReplicaID() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicaID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// ReplicaID indicates an expected call of ReplicaID
func (mr *MockPeerMessageMockRecorder) ReplicaID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaID", reflect.TypeOf((*MockPeerMessage)(nil).ReplicaID))
}

// MockCertifiedMessage is a mock of CertifiedMessage interface
type MockCertifiedMessage struct {
	ctrl     *gomock.Controller
	recorder *MockCertifiedMessageMockRecorder
}

// MockCertifiedMessageMockRecorder is the mock recorder for MockCertifiedMessage
type MockCertifiedMessageMockRecorder struct {
	mock *MockCertifiedMessage
}

// NewMockCertifiedMessage creates a new mock instance
func NewMockCertifiedMessage(ctrl *gomock.Controller) *MockCertifiedMessage {
	mock := &MockCertifiedMessage{ctrl: ctrl}
	mock.recorder = &MockCertifiedMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertifiedMessage) EXPECT() *MockCertifiedMessageMockRecorder {
	return m.recorder
}

// ImplementsReplicaMessage mocks base method
func (m *MockCertifiedMessage) ImplementsReplicaMessage() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ImplementsReplicaMessage")
}

// ImplementsReplicaMessage indicates an expected call of ImplementsReplicaMessage
func (mr *MockCertifiedMessageMockRecorder) ImplementsReplicaMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImplementsReplicaMessage", reflect.TypeOf((*MockCertifiedMessage)(nil).ImplementsReplicaMessage))
}

// MarshalBinary mocks base method
func (m *MockCertifiedMessage) MarshalBinary() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalBinary")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalBinary indicates an expected call of MarshalBinary
func (mr *MockCertifiedMessageMockRecorder) MarshalBinary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalBinary", reflect.TypeOf((*MockCertifiedMessage)(nil).MarshalBinary))
}

// ReplicaID mocks base method
func (m *MockCertifiedMessage) ReplicaID() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicaID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// ReplicaID indicates an expected call of ReplicaID
func (mr *MockCertifiedMessageMockRecorder) ReplicaID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaID", reflect.TypeOf((*MockCertifiedMessage)(nil).ReplicaID))
}

// SetUIBytes mocks base method
func (m *MockCertifiedMessage) SetUIBytes(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUIBytes", arg0)
}

// SetUIBytes indicates an expected call of SetUIBytes
func (mr *MockCertifiedMessageMockRecorder) SetUIBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUIBytes", reflect.TypeOf((*MockCertifiedMessage)(nil).SetUIBytes), arg0)
}

// UIBytes mocks base method
func (m *MockCertifiedMessage) UIBytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UIBytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// UIBytes indicates an expected call of UIBytes
func (mr *MockCertifiedMessageMockRecorder) UIBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UIBytes", reflect.TypeOf((*MockCertifiedMessage)(nil).UIBytes))
}

// MockSignedMessage is a mock of SignedMessage interface
type MockSignedMessage struct {
	ctrl     *gomock.Controller
	recorder *MockSignedMessageMockRecorder
}

// MockSignedMessageMockRecorder is the mock recorder for MockSignedMessage
type MockSignedMessageMockRecorder struct {
	mock *MockSignedMessage
}

// NewMockSignedMessage creates a new mock instance
func NewMockSignedMessage(ctrl *gomock.Controller) *MockSignedMessage {
	mock := &MockSignedMessage{ctrl: ctrl}
	mock.recorder = &MockSignedMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSignedMessage) EXPECT() *MockSignedMessageMockRecorder {
	return m.recorder
}

// SetSignature mocks base method
func (m *MockSignedMessage) SetSignature(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSignature", arg0)
}

// SetSignature indicates an expected call of SetSignature
func (mr *MockSignedMessageMockRecorder) SetSignature(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSignature", reflect.TypeOf((*MockSignedMessage)(nil).SetSignature), arg0)
}

// Signature mocks base method
func (m *MockSignedMessage) Signature() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signature")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Signature indicates an expected call of Signature
func (mr *MockSignedMessageMockRecorder) Signature() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signature", reflect.TypeOf((*MockSignedMessage)(nil).Signature))
}
