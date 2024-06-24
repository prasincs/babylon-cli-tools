// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/expected_interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	services "github.com/babylonchain/cli-tools/internal/services"
	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"
	wire "github.com/btcsuite/btcd/wire"
	gomock "github.com/golang/mock/gomock"
)

// MockCovenantSigner is a mock of CovenantSigner interface.
type MockCovenantSigner struct {
	ctrl     *gomock.Controller
	recorder *MockCovenantSignerMockRecorder
}

// MockCovenantSignerMockRecorder is the mock recorder for MockCovenantSigner.
type MockCovenantSignerMockRecorder struct {
	mock *MockCovenantSigner
}

// NewMockCovenantSigner creates a new mock instance.
func NewMockCovenantSigner(ctrl *gomock.Controller) *MockCovenantSigner {
	mock := &MockCovenantSigner{ctrl: ctrl}
	mock.recorder = &MockCovenantSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCovenantSigner) EXPECT() *MockCovenantSignerMockRecorder {
	return m.recorder
}

// SignUnbondingTransaction mocks base method.
func (m *MockCovenantSigner) SignUnbondingTransaction(req *services.SignRequest) (*services.PubKeySigPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUnbondingTransaction", req)
	ret0, _ := ret[0].(*services.PubKeySigPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUnbondingTransaction indicates an expected call of SignUnbondingTransaction.
func (mr *MockCovenantSignerMockRecorder) SignUnbondingTransaction(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUnbondingTransaction", reflect.TypeOf((*MockCovenantSigner)(nil).SignUnbondingTransaction), req)
}

// MockBtcSender is a mock of BtcSender interface.
type MockBtcSender struct {
	ctrl     *gomock.Controller
	recorder *MockBtcSenderMockRecorder
}

// MockBtcSenderMockRecorder is the mock recorder for MockBtcSender.
type MockBtcSenderMockRecorder struct {
	mock *MockBtcSender
}

// NewMockBtcSender creates a new mock instance.
func NewMockBtcSender(ctrl *gomock.Controller) *MockBtcSender {
	mock := &MockBtcSender{ctrl: ctrl}
	mock.recorder = &MockBtcSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBtcSender) EXPECT() *MockBtcSenderMockRecorder {
	return m.recorder
}

// CheckTxOutSpendable mocks base method.
func (m *MockBtcSender) CheckTxOutSpendable(txHash *chainhash.Hash, index uint32, mempool bool) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTxOutSpendable", txHash, index, mempool)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTxOutSpendable indicates an expected call of CheckTxOutSpendable.
func (mr *MockBtcSenderMockRecorder) CheckTxOutSpendable(txHash, index, mempool interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTxOutSpendable", reflect.TypeOf((*MockBtcSender)(nil).CheckTxOutSpendable), txHash, index, mempool)
}

// SendTx mocks base method.
func (m *MockBtcSender) SendTx(tx *wire.MsgTx) (*chainhash.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTx", tx)
	ret0, _ := ret[0].(*chainhash.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTx indicates an expected call of SendTx.
func (mr *MockBtcSenderMockRecorder) SendTx(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTx", reflect.TypeOf((*MockBtcSender)(nil).SendTx), tx)
}

// TxByHash mocks base method.
func (m *MockBtcSender) TxByHash(txHash *chainhash.Hash, pkScript []byte) (*services.TxInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxByHash", txHash, pkScript)
	ret0, _ := ret[0].(*services.TxInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxByHash indicates an expected call of TxByHash.
func (mr *MockBtcSenderMockRecorder) TxByHash(txHash, pkScript interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxByHash", reflect.TypeOf((*MockBtcSender)(nil).TxByHash), txHash, pkScript)
}

// MockParamsRetriever is a mock of ParamsRetriever interface.
type MockParamsRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockParamsRetrieverMockRecorder
}

// MockParamsRetrieverMockRecorder is the mock recorder for MockParamsRetriever.
type MockParamsRetrieverMockRecorder struct {
	mock *MockParamsRetriever
}

// NewMockParamsRetriever creates a new mock instance.
func NewMockParamsRetriever(ctrl *gomock.Controller) *MockParamsRetriever {
	mock := &MockParamsRetriever{ctrl: ctrl}
	mock.recorder = &MockParamsRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParamsRetriever) EXPECT() *MockParamsRetrieverMockRecorder {
	return m.recorder
}

// ParamsByHeight mocks base method.
func (m *MockParamsRetriever) ParamsByHeight(ctx context.Context, height uint64) (*services.SystemParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamsByHeight", ctx, height)
	ret0, _ := ret[0].(*services.SystemParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParamsByHeight indicates an expected call of ParamsByHeight.
func (mr *MockParamsRetrieverMockRecorder) ParamsByHeight(ctx, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamsByHeight", reflect.TypeOf((*MockParamsRetriever)(nil).ParamsByHeight), ctx, height)
}

// MockUnbondingStore is a mock of UnbondingStore interface.
type MockUnbondingStore struct {
	ctrl     *gomock.Controller
	recorder *MockUnbondingStoreMockRecorder
}

// MockUnbondingStoreMockRecorder is the mock recorder for MockUnbondingStore.
type MockUnbondingStoreMockRecorder struct {
	mock *MockUnbondingStore
}

// NewMockUnbondingStore creates a new mock instance.
func NewMockUnbondingStore(ctrl *gomock.Controller) *MockUnbondingStore {
	mock := &MockUnbondingStore{ctrl: ctrl}
	mock.recorder = &MockUnbondingStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnbondingStore) EXPECT() *MockUnbondingStoreMockRecorder {
	return m.recorder
}

// GetFailedUnbondingTransactions mocks base method.
func (m *MockUnbondingStore) GetFailedUnbondingTransactions(ctx context.Context) ([]*services.UnbondingTxData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFailedUnbondingTransactions", ctx)
	ret0, _ := ret[0].([]*services.UnbondingTxData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFailedUnbondingTransactions indicates an expected call of GetFailedUnbondingTransactions.
func (mr *MockUnbondingStoreMockRecorder) GetFailedUnbondingTransactions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFailedUnbondingTransactions", reflect.TypeOf((*MockUnbondingStore)(nil).GetFailedUnbondingTransactions), ctx)
}

// GetNotProcessedUnbondingTransactions mocks base method.
func (m *MockUnbondingStore) GetNotProcessedUnbondingTransactions(ctx context.Context) ([]*services.UnbondingTxData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotProcessedUnbondingTransactions", ctx)
	ret0, _ := ret[0].([]*services.UnbondingTxData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotProcessedUnbondingTransactions indicates an expected call of GetNotProcessedUnbondingTransactions.
func (mr *MockUnbondingStoreMockRecorder) GetNotProcessedUnbondingTransactions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotProcessedUnbondingTransactions", reflect.TypeOf((*MockUnbondingStore)(nil).GetNotProcessedUnbondingTransactions), ctx)
}

// GetUnbondingTransactionsWithNoQuorum mocks base method.
func (m *MockUnbondingStore) GetUnbondingTransactionsWithNoQuorum(ctx context.Context) ([]*services.UnbondingTxData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnbondingTransactionsWithNoQuorum", ctx)
	ret0, _ := ret[0].([]*services.UnbondingTxData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnbondingTransactionsWithNoQuorum indicates an expected call of GetUnbondingTransactionsWithNoQuorum.
func (mr *MockUnbondingStoreMockRecorder) GetUnbondingTransactionsWithNoQuorum(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnbondingTransactionsWithNoQuorum", reflect.TypeOf((*MockUnbondingStore)(nil).GetUnbondingTransactionsWithNoQuorum), ctx)
}

// SetUnbondingTransactionFailedToGetCovenantSignatures mocks base method.
func (m *MockUnbondingStore) SetUnbondingTransactionFailedToGetCovenantSignatures(ctx context.Context, utx *services.UnbondingTxData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnbondingTransactionFailedToGetCovenantSignatures", ctx, utx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnbondingTransactionFailedToGetCovenantSignatures indicates an expected call of SetUnbondingTransactionFailedToGetCovenantSignatures.
func (mr *MockUnbondingStoreMockRecorder) SetUnbondingTransactionFailedToGetCovenantSignatures(ctx, utx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnbondingTransactionFailedToGetCovenantSignatures", reflect.TypeOf((*MockUnbondingStore)(nil).SetUnbondingTransactionFailedToGetCovenantSignatures), ctx, utx)
}

// SetUnbondingTransactionInputAlreadySpent mocks base method.
func (m *MockUnbondingStore) SetUnbondingTransactionInputAlreadySpent(ctx context.Context, utx *services.UnbondingTxData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnbondingTransactionInputAlreadySpent", ctx, utx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnbondingTransactionInputAlreadySpent indicates an expected call of SetUnbondingTransactionInputAlreadySpent.
func (mr *MockUnbondingStoreMockRecorder) SetUnbondingTransactionInputAlreadySpent(ctx, utx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnbondingTransactionInputAlreadySpent", reflect.TypeOf((*MockUnbondingStore)(nil).SetUnbondingTransactionInputAlreadySpent), ctx, utx)
}

// SetUnbondingTransactionProcessed mocks base method.
func (m *MockUnbondingStore) SetUnbondingTransactionProcessed(ctx context.Context, utx *services.UnbondingTxData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnbondingTransactionProcessed", ctx, utx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnbondingTransactionProcessed indicates an expected call of SetUnbondingTransactionProcessed.
func (mr *MockUnbondingStoreMockRecorder) SetUnbondingTransactionProcessed(ctx, utx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnbondingTransactionProcessed", reflect.TypeOf((*MockUnbondingStore)(nil).SetUnbondingTransactionProcessed), ctx, utx)
}

// SetUnbondingTransactionProcessingFailed mocks base method.
func (m *MockUnbondingStore) SetUnbondingTransactionProcessingFailed(ctx context.Context, utx *services.UnbondingTxData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnbondingTransactionProcessingFailed", ctx, utx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnbondingTransactionProcessingFailed indicates an expected call of SetUnbondingTransactionProcessingFailed.
func (mr *MockUnbondingStoreMockRecorder) SetUnbondingTransactionProcessingFailed(ctx, utx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnbondingTransactionProcessingFailed", reflect.TypeOf((*MockUnbondingStore)(nil).SetUnbondingTransactionProcessingFailed), ctx, utx)
}