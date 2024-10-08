// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/CoreumFoundation/iso20022-client/iso20022/processes (interfaces: ContractClient,AddressBook,Cryptography,SupplementaryDataParser,Parser,MessageQueue,Dtif)
//
// Generated by this command:
//
//	mockgen -destination=model_mocks_test.go -package=processes_test . ContractClient,AddressBook,Cryptography,SupplementaryDataParser,Parser,MessageQueue,Dtif
//

// Package processes_test is a generated GoMock package.
package processes_test

import (
	context "context"
	big "math/big"
	reflect "reflect"
	time "time"

	client "github.com/CoreumFoundation/coreum/v4/pkg/client"
	messages "github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/messages"
	addressbook "github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	coreum "github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	processes "github.com/CoreumFoundation/iso20022-client/iso20022/processes"
	queue "github.com/CoreumFoundation/iso20022-client/iso20022/queue"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	types0 "github.com/cosmos/cosmos-sdk/crypto/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	gomock "go.uber.org/mock/gomock"
)

// MockContractClient is a mock of ContractClient interface.
type MockContractClient struct {
	ctrl     *gomock.Controller
	recorder *MockContractClientMockRecorder
}

// MockContractClientMockRecorder is the mock recorder for MockContractClient.
type MockContractClientMockRecorder struct {
	mock *MockContractClient
}

// NewMockContractClient creates a new mock instance.
func NewMockContractClient(ctrl *gomock.Controller) *MockContractClient {
	mock := &MockContractClient{ctrl: ctrl}
	mock.recorder = &MockContractClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractClient) EXPECT() *MockContractClientMockRecorder {
	return m.recorder
}

// BroadcastMessages mocks base method.
func (m *MockContractClient) BroadcastMessages(arg0 context.Context, arg1 types1.AccAddress, arg2 ...types1.Msg) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BroadcastMessages", varargs...)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastMessages indicates an expected call of BroadcastMessages.
func (mr *MockContractClientMockRecorder) BroadcastMessages(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastMessages", reflect.TypeOf((*MockContractClient)(nil).BroadcastMessages), varargs...)
}

// CancelSession mocks base method.
func (m *MockContractClient) CancelSession(arg0 context.Context, arg1 string, arg2, arg3, arg4 types1.AccAddress) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelSession", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelSession indicates an expected call of CancelSession.
func (mr *MockContractClientMockRecorder) CancelSession(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelSession", reflect.TypeOf((*MockContractClient)(nil).CancelSession), arg0, arg1, arg2, arg3, arg4)
}

// CancelSessions mocks base method.
func (m *MockContractClient) CancelSessions(arg0 context.Context, arg1 types1.AccAddress, arg2 ...coreum.CancelSession) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelSessions", varargs...)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelSessions indicates an expected call of CancelSessions.
func (mr *MockContractClientMockRecorder) CancelSessions(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelSessions", reflect.TypeOf((*MockContractClient)(nil).CancelSessions), varargs...)
}

// ConfirmSession mocks base method.
func (m *MockContractClient) ConfirmSession(arg0 context.Context, arg1 string, arg2, arg3, arg4 types1.AccAddress) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmSession", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmSession indicates an expected call of ConfirmSession.
func (mr *MockContractClientMockRecorder) ConfirmSession(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmSession", reflect.TypeOf((*MockContractClient)(nil).ConfirmSession), arg0, arg1, arg2, arg3, arg4)
}

// ConfirmSessions mocks base method.
func (m *MockContractClient) ConfirmSessions(arg0 context.Context, arg1 types1.AccAddress, arg2 ...coreum.ConfirmSession) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfirmSessions", varargs...)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmSessions indicates an expected call of ConfirmSessions.
func (mr *MockContractClientMockRecorder) ConfirmSessions(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmSessions", reflect.TypeOf((*MockContractClient)(nil).ConfirmSessions), varargs...)
}

// DeployAndInstantiate mocks base method.
func (m *MockContractClient) DeployAndInstantiate(arg0 context.Context, arg1 types1.AccAddress, arg2 string) (types1.AccAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployAndInstantiate", arg0, arg1, arg2)
	ret0, _ := ret[0].(types1.AccAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployAndInstantiate indicates an expected call of DeployAndInstantiate.
func (mr *MockContractClientMockRecorder) DeployAndInstantiate(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployAndInstantiate", reflect.TypeOf((*MockContractClient)(nil).DeployAndInstantiate), arg0, arg1, arg2)
}

// DeployContract mocks base method.
func (m *MockContractClient) DeployContract(arg0 context.Context, arg1 types1.AccAddress, arg2 string) (*types1.TxResponse, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployContract", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeployContract indicates an expected call of DeployContract.
func (mr *MockContractClientMockRecorder) DeployContract(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployContract", reflect.TypeOf((*MockContractClient)(nil).DeployContract), arg0, arg1, arg2)
}

// GetActiveSessions mocks base method.
func (m *MockContractClient) GetActiveSessions(arg0 context.Context, arg1 types1.AccAddress, arg2 coreum.UserType, arg3 *string, arg4 *uint32) ([]coreum.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveSessions", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]coreum.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveSessions indicates an expected call of GetActiveSessions.
func (mr *MockContractClientMockRecorder) GetActiveSessions(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveSessions", reflect.TypeOf((*MockContractClient)(nil).GetActiveSessions), arg0, arg1, arg2, arg3, arg4)
}

// GetClosedSessions mocks base method.
func (m *MockContractClient) GetClosedSessions(arg0 context.Context, arg1 types1.AccAddress, arg2 coreum.UserType, arg3 *string, arg4 *uint32) ([]coreum.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClosedSessions", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]coreum.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClosedSessions indicates an expected call of GetClosedSessions.
func (mr *MockContractClientMockRecorder) GetClosedSessions(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClosedSessions", reflect.TypeOf((*MockContractClient)(nil).GetClosedSessions), arg0, arg1, arg2, arg3, arg4)
}

// GetContractAddress mocks base method.
func (m *MockContractClient) GetContractAddress() types1.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractAddress")
	ret0, _ := ret[0].(types1.AccAddress)
	return ret0
}

// GetContractAddress indicates an expected call of GetContractAddress.
func (mr *MockContractClientMockRecorder) GetContractAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractAddress", reflect.TypeOf((*MockContractClient)(nil).GetContractAddress))
}

// GetMessages mocks base method.
func (m *MockContractClient) GetMessages(arg0 context.Context, arg1 types1.AccAddress, arg2 string, arg3 *uint32) ([]coreum.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessages", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]coreum.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessages indicates an expected call of GetMessages.
func (mr *MockContractClientMockRecorder) GetMessages(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockContractClient)(nil).GetMessages), arg0, arg1, arg2, arg3)
}

// GetNewMessages mocks base method.
func (m *MockContractClient) GetNewMessages(arg0 context.Context, arg1 types1.AccAddress, arg2 *uint32) ([]coreum.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewMessages", arg0, arg1, arg2)
	ret0, _ := ret[0].([]coreum.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNewMessages indicates an expected call of GetNewMessages.
func (mr *MockContractClientMockRecorder) GetNewMessages(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewMessages", reflect.TypeOf((*MockContractClient)(nil).GetNewMessages), arg0, arg1, arg2)
}

// IsInitialized mocks base method.
func (m *MockContractClient) IsInitialized() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInitialized")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInitialized indicates an expected call of IsInitialized.
func (mr *MockContractClientMockRecorder) IsInitialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInitialized", reflect.TypeOf((*MockContractClient)(nil).IsInitialized))
}

// IssueNFTClass mocks base method.
func (m *MockContractClient) IssueNFTClass(arg0 context.Context, arg1 types1.AccAddress, arg2, arg3, arg4 string) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueNFTClass", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueNFTClass indicates an expected call of IssueNFTClass.
func (mr *MockContractClientMockRecorder) IssueNFTClass(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueNFTClass", reflect.TypeOf((*MockContractClient)(nil).IssueNFTClass), arg0, arg1, arg2, arg3, arg4)
}

// MarkAsRead mocks base method.
func (m *MockContractClient) MarkAsRead(arg0 context.Context, arg1 types1.AccAddress, arg2 uint64) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkAsRead", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkAsRead indicates an expected call of MarkAsRead.
func (mr *MockContractClientMockRecorder) MarkAsRead(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAsRead", reflect.TypeOf((*MockContractClient)(nil).MarkAsRead), arg0, arg1, arg2)
}

// MigrateContract mocks base method.
func (m *MockContractClient) MigrateContract(arg0 context.Context, arg1 types1.AccAddress, arg2 uint64) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateContract", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrateContract indicates an expected call of MigrateContract.
func (mr *MockContractClientMockRecorder) MigrateContract(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateContract", reflect.TypeOf((*MockContractClient)(nil).MigrateContract), arg0, arg1, arg2)
}

// MintNFT mocks base method.
func (m *MockContractClient) MintNFT(arg0 context.Context, arg1 types1.AccAddress, arg2, arg3 string, arg4 *types.Any) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintNFT", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MintNFT indicates an expected call of MintNFT.
func (mr *MockContractClientMockRecorder) MintNFT(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintNFT", reflect.TypeOf((*MockContractClient)(nil).MintNFT), arg0, arg1, arg2, arg3, arg4)
}

// QueryNFT mocks base method.
func (m *MockContractClient) QueryNFT(arg0 context.Context, arg1, arg2 string) (*types.Any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryNFT", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryNFT indicates an expected call of QueryNFT.
func (mr *MockContractClientMockRecorder) QueryNFT(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryNFT", reflect.TypeOf((*MockContractClient)(nil).QueryNFT), arg0, arg1, arg2)
}

// SendMessage mocks base method.
func (m *MockContractClient) SendMessage(arg0 context.Context, arg1, arg2 types1.AccAddress, arg3, arg4 string, arg5 coreum.NFTInfo) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockContractClientMockRecorder) SendMessage(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockContractClient)(nil).SendMessage), arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendMessages mocks base method.
func (m *MockContractClient) SendMessages(arg0 context.Context, arg1 types1.AccAddress, arg2 ...coreum.SendMessage) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessages", varargs...)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessages indicates an expected call of SendMessages.
func (mr *MockContractClientMockRecorder) SendMessages(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessages", reflect.TypeOf((*MockContractClient)(nil).SendMessages), varargs...)
}

// SetContractAddress mocks base method.
func (m *MockContractClient) SetContractAddress(arg0 types1.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContractAddress", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetContractAddress indicates an expected call of SetContractAddress.
func (mr *MockContractClientMockRecorder) SetContractAddress(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContractAddress", reflect.TypeOf((*MockContractClient)(nil).SetContractAddress), arg0)
}

// StartSession mocks base method.
func (m *MockContractClient) StartSession(arg0 context.Context, arg1 string, arg2 types1.AccAddress, arg3 coreum.NFTInfo, arg4 types1.AccAddress, arg5 types1.Coins) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSession", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartSession indicates an expected call of StartSession.
func (mr *MockContractClientMockRecorder) StartSession(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSession", reflect.TypeOf((*MockContractClient)(nil).StartSession), arg0, arg1, arg2, arg3, arg4, arg5)
}

// StartSessions mocks base method.
func (m *MockContractClient) StartSessions(arg0 context.Context, arg1 types1.AccAddress, arg2 ...coreum.StartSession) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartSessions", varargs...)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartSessions indicates an expected call of StartSessions.
func (mr *MockContractClientMockRecorder) StartSessions(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSessions", reflect.TypeOf((*MockContractClient)(nil).StartSessions), varargs...)
}

// MockAddressBook is a mock of AddressBook interface.
type MockAddressBook struct {
	ctrl     *gomock.Controller
	recorder *MockAddressBookMockRecorder
}

// MockAddressBookMockRecorder is the mock recorder for MockAddressBook.
type MockAddressBookMockRecorder struct {
	mock *MockAddressBook
}

// NewMockAddressBook creates a new mock instance.
func NewMockAddressBook(ctrl *gomock.Controller) *MockAddressBook {
	mock := &MockAddressBook{ctrl: ctrl}
	mock.recorder = &MockAddressBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddressBook) EXPECT() *MockAddressBookMockRecorder {
	return m.recorder
}

// Lookup mocks base method.
func (m *MockAddressBook) Lookup(arg0 addressbook.Party) (*addressbook.Address, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0)
	ret0, _ := ret[0].(*addressbook.Address)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup.
func (mr *MockAddressBookMockRecorder) Lookup(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockAddressBook)(nil).Lookup), arg0)
}

// LookupByAccountAddress mocks base method.
func (m *MockAddressBook) LookupByAccountAddress(arg0 string) (*addressbook.Address, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupByAccountAddress", arg0)
	ret0, _ := ret[0].(*addressbook.Address)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// LookupByAccountAddress indicates an expected call of LookupByAccountAddress.
func (mr *MockAddressBookMockRecorder) LookupByAccountAddress(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupByAccountAddress", reflect.TypeOf((*MockAddressBook)(nil).LookupByAccountAddress), arg0)
}

// Update mocks base method.
func (m *MockAddressBook) Update(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAddressBookMockRecorder) Update(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAddressBook)(nil).Update), arg0)
}

// Validate mocks base method.
func (m *MockAddressBook) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockAddressBookMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAddressBook)(nil).Validate))
}

// MockCryptography is a mock of Cryptography interface.
type MockCryptography struct {
	ctrl     *gomock.Controller
	recorder *MockCryptographyMockRecorder
}

// MockCryptographyMockRecorder is the mock recorder for MockCryptography.
type MockCryptographyMockRecorder struct {
	mock *MockCryptography
}

// NewMockCryptography creates a new mock instance.
func NewMockCryptography(ctrl *gomock.Controller) *MockCryptography {
	mock := &MockCryptography{ctrl: ctrl}
	mock.recorder = &MockCryptographyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptography) EXPECT() *MockCryptographyMockRecorder {
	return m.recorder
}

// DecryptSymmetric mocks base method.
func (m *MockCryptography) DecryptSymmetric(arg0, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecryptSymmetric", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecryptSymmetric indicates an expected call of DecryptSymmetric.
func (mr *MockCryptographyMockRecorder) DecryptSymmetric(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptSymmetric", reflect.TypeOf((*MockCryptography)(nil).DecryptSymmetric), arg0, arg1)
}

// EncryptSymmetric mocks base method.
func (m *MockCryptography) EncryptSymmetric(arg0, arg1 []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptSymmetric", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// EncryptSymmetric indicates an expected call of EncryptSymmetric.
func (mr *MockCryptographyMockRecorder) EncryptSymmetric(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptSymmetric", reflect.TypeOf((*MockCryptography)(nil).EncryptSymmetric), arg0, arg1)
}

// GenerateSharedKey mocks base method.
func (m *MockCryptography) GenerateSharedKey(arg0 types0.PrivKey, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSharedKey", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateSharedKey indicates an expected call of GenerateSharedKey.
func (mr *MockCryptographyMockRecorder) GenerateSharedKey(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSharedKey", reflect.TypeOf((*MockCryptography)(nil).GenerateSharedKey), arg0, arg1)
}

// GenerateSharedKeyByPrivateKeyName mocks base method.
func (m *MockCryptography) GenerateSharedKeyByPrivateKeyName(arg0 client.Context, arg1 string, arg2 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSharedKeyByPrivateKeyName", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateSharedKeyByPrivateKeyName indicates an expected call of GenerateSharedKeyByPrivateKeyName.
func (mr *MockCryptographyMockRecorder) GenerateSharedKeyByPrivateKeyName(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSharedKeyByPrivateKeyName", reflect.TypeOf((*MockCryptography)(nil).GenerateSharedKeyByPrivateKeyName), arg0, arg1, arg2)
}

// MockSupplementaryDataParser is a mock of SupplementaryDataParser interface.
type MockSupplementaryDataParser struct {
	ctrl     *gomock.Controller
	recorder *MockSupplementaryDataParserMockRecorder
}

// MockSupplementaryDataParserMockRecorder is the mock recorder for MockSupplementaryDataParser.
type MockSupplementaryDataParserMockRecorder struct {
	mock *MockSupplementaryDataParser
}

// NewMockSupplementaryDataParser creates a new mock instance.
func NewMockSupplementaryDataParser(ctrl *gomock.Controller) *MockSupplementaryDataParser {
	mock := &MockSupplementaryDataParser{ctrl: ctrl}
	mock.recorder = &MockSupplementaryDataParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSupplementaryDataParser) EXPECT() *MockSupplementaryDataParserMockRecorder {
	return m.recorder
}

// Parse mocks base method.
func (m *MockSupplementaryDataParser) Parse(arg0 []byte) (messages.Iso20022Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", arg0)
	ret0, _ := ret[0].(messages.Iso20022Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockSupplementaryDataParserMockRecorder) Parse(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockSupplementaryDataParser)(nil).Parse), arg0)
}

// MockParser is a mock of Parser interface.
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser.
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance.
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// ExtractMessageAndMetadataFromIsoMessage mocks base method.
func (m *MockParser) ExtractMessageAndMetadataFromIsoMessage(arg0 []byte) (messages.Iso20022Message, processes.Metadata, *processes.Metadata, processes.SupplementaryDataParser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractMessageAndMetadataFromIsoMessage", arg0)
	ret0, _ := ret[0].(messages.Iso20022Message)
	ret1, _ := ret[1].(processes.Metadata)
	ret2, _ := ret[2].(*processes.Metadata)
	ret3, _ := ret[3].(processes.SupplementaryDataParser)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// ExtractMessageAndMetadataFromIsoMessage indicates an expected call of ExtractMessageAndMetadataFromIsoMessage.
func (mr *MockParserMockRecorder) ExtractMessageAndMetadataFromIsoMessage(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractMessageAndMetadataFromIsoMessage", reflect.TypeOf((*MockParser)(nil).ExtractMessageAndMetadataFromIsoMessage), arg0)
}

// GetSupplementaryDataWithCorrectClearingSystem mocks base method.
func (m *MockParser) GetSupplementaryDataWithCorrectClearingSystem(arg0 messages.Iso20022Message, arg1 string) ([]byte, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupplementaryDataWithCorrectClearingSystem", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetSupplementaryDataWithCorrectClearingSystem indicates an expected call of GetSupplementaryDataWithCorrectClearingSystem.
func (mr *MockParserMockRecorder) GetSupplementaryDataWithCorrectClearingSystem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupplementaryDataWithCorrectClearingSystem", reflect.TypeOf((*MockParser)(nil).GetSupplementaryDataWithCorrectClearingSystem), arg0, arg1)
}

// GetTransactionStatus mocks base method.
func (m *MockParser) GetTransactionStatus(arg0 messages.Iso20022Message) processes.TransactionStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionStatus", arg0)
	ret0, _ := ret[0].(processes.TransactionStatus)
	return ret0
}

// GetTransactionStatus indicates an expected call of GetTransactionStatus.
func (mr *MockParserMockRecorder) GetTransactionStatus(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionStatus", reflect.TypeOf((*MockParser)(nil).GetTransactionStatus), arg0)
}

// MockMessageQueue is a mock of MessageQueue interface.
type MockMessageQueue struct {
	ctrl     *gomock.Controller
	recorder *MockMessageQueueMockRecorder
}

// MockMessageQueueMockRecorder is the mock recorder for MockMessageQueue.
type MockMessageQueueMockRecorder struct {
	mock *MockMessageQueue
}

// NewMockMessageQueue creates a new mock instance.
func NewMockMessageQueue(ctrl *gomock.Controller) *MockMessageQueue {
	mock := &MockMessageQueue{ctrl: ctrl}
	mock.recorder = &MockMessageQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageQueue) EXPECT() *MockMessageQueueMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMessageQueue) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockMessageQueueMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMessageQueue)(nil).Close))
}

// GetStatus mocks base method.
func (m *MockMessageQueue) GetStatus(arg0 string) *queue.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", arg0)
	ret0, _ := ret[0].(*queue.Status)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockMessageQueueMockRecorder) GetStatus(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockMessageQueue)(nil).GetStatus), arg0)
}

// PopFromReceive mocks base method.
func (m *MockMessageQueue) PopFromReceive() ([]byte, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopFromReceive")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// PopFromReceive indicates an expected call of PopFromReceive.
func (mr *MockMessageQueueMockRecorder) PopFromReceive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopFromReceive", reflect.TypeOf((*MockMessageQueue)(nil).PopFromReceive))
}

// PopFromSend mocks base method.
func (m *MockMessageQueue) PopFromSend(arg0 context.Context, arg1 int, arg2 time.Duration) [][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopFromSend", arg0, arg1, arg2)
	ret0, _ := ret[0].([][]byte)
	return ret0
}

// PopFromSend indicates an expected call of PopFromSend.
func (mr *MockMessageQueueMockRecorder) PopFromSend(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopFromSend", reflect.TypeOf((*MockMessageQueue)(nil).PopFromSend), arg0, arg1, arg2)
}

// PushToReceive mocks base method.
func (m *MockMessageQueue) PushToReceive(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PushToReceive", arg0)
}

// PushToReceive indicates an expected call of PushToReceive.
func (mr *MockMessageQueueMockRecorder) PushToReceive(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushToReceive", reflect.TypeOf((*MockMessageQueue)(nil).PushToReceive), arg0)
}

// PushToSend mocks base method.
func (m *MockMessageQueue) PushToSend(arg0 string, arg1 []byte) queue.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushToSend", arg0, arg1)
	ret0, _ := ret[0].(queue.Status)
	return ret0
}

// PushToSend indicates an expected call of PushToSend.
func (mr *MockMessageQueueMockRecorder) PushToSend(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushToSend", reflect.TypeOf((*MockMessageQueue)(nil).PushToSend), arg0, arg1)
}

// SetStatus mocks base method.
func (m *MockMessageQueue) SetStatus(arg0 string, arg1 queue.Status) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStatus", arg0, arg1)
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockMessageQueueMockRecorder) SetStatus(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockMessageQueue)(nil).SetStatus), arg0, arg1)
}

// Start mocks base method.
func (m *MockMessageQueue) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockMessageQueueMockRecorder) Start(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMessageQueue)(nil).Start), arg0)
}

// MockDtif is a mock of Dtif interface.
type MockDtif struct {
	ctrl     *gomock.Controller
	recorder *MockDtifMockRecorder
}

// MockDtifMockRecorder is the mock recorder for MockDtif.
type MockDtifMockRecorder struct {
	mock *MockDtif
}

// NewMockDtif creates a new mock instance.
func NewMockDtif(ctrl *gomock.Controller) *MockDtif {
	mock := &MockDtif{ctrl: ctrl}
	mock.recorder = &MockDtifMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDtif) EXPECT() *MockDtifMockRecorder {
	return m.recorder
}

// LookupByDTI mocks base method.
func (m *MockDtif) LookupByDTI(arg0 string) (string, *big.Int, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupByDTI", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*big.Int)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// LookupByDTI indicates an expected call of LookupByDTI.
func (mr *MockDtifMockRecorder) LookupByDTI(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupByDTI", reflect.TypeOf((*MockDtif)(nil).LookupByDTI), arg0)
}

// LookupByDenom mocks base method.
func (m *MockDtif) LookupByDenom(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupByDenom", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// LookupByDenom indicates an expected call of LookupByDenom.
func (mr *MockDtifMockRecorder) LookupByDenom(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupByDenom", reflect.TypeOf((*MockDtif)(nil).LookupByDenom), arg0)
}

// Update mocks base method.
func (m *MockDtif) Update(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDtifMockRecorder) Update(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDtif)(nil).Update), arg0)
}
