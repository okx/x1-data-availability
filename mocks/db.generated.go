// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygon/cdk-data-availability/types"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

type DB_Expecter struct {
	mock *mock.Mock
}

func (_m *DB) EXPECT() *DB_Expecter {
	return &DB_Expecter{mock: &_m.Mock}
}

// CountOffchainData provides a mock function with given fields: ctx
func (_m *DB) CountOffchainData(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CountOffchainData")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_CountOffchainData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountOffchainData'
type DB_CountOffchainData_Call struct {
	*mock.Call
}

// CountOffchainData is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DB_Expecter) CountOffchainData(ctx interface{}) *DB_CountOffchainData_Call {
	return &DB_CountOffchainData_Call{Call: _e.mock.On("CountOffchainData", ctx)}
}

func (_c *DB_CountOffchainData_Call) Run(run func(ctx context.Context)) *DB_CountOffchainData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_CountOffchainData_Call) Return(_a0 uint64, _a1 error) *DB_CountOffchainData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_CountOffchainData_Call) RunAndReturn(run func(context.Context) (uint64, error)) *DB_CountOffchainData_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUnresolvedBatchKeys provides a mock function with given fields: ctx, bks
func (_m *DB) DeleteUnresolvedBatchKeys(ctx context.Context, bks []types.BatchKey) error {
	ret := _m.Called(ctx, bks)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUnresolvedBatchKeys")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.BatchKey) error); ok {
		r0 = rf(ctx, bks)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_DeleteUnresolvedBatchKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUnresolvedBatchKeys'
type DB_DeleteUnresolvedBatchKeys_Call struct {
	*mock.Call
}

// DeleteUnresolvedBatchKeys is a helper method to define mock.On call
//   - ctx context.Context
//   - bks []types.BatchKey
func (_e *DB_Expecter) DeleteUnresolvedBatchKeys(ctx interface{}, bks interface{}) *DB_DeleteUnresolvedBatchKeys_Call {
	return &DB_DeleteUnresolvedBatchKeys_Call{Call: _e.mock.On("DeleteUnresolvedBatchKeys", ctx, bks)}
}

func (_c *DB_DeleteUnresolvedBatchKeys_Call) Run(run func(ctx context.Context, bks []types.BatchKey)) *DB_DeleteUnresolvedBatchKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.BatchKey))
	})
	return _c
}

func (_c *DB_DeleteUnresolvedBatchKeys_Call) Return(_a0 error) *DB_DeleteUnresolvedBatchKeys_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_DeleteUnresolvedBatchKeys_Call) RunAndReturn(run func(context.Context, []types.BatchKey) error) *DB_DeleteUnresolvedBatchKeys_Call {
	_c.Call.Return(run)
	return _c
}

// DetectOffchainDataGaps provides a mock function with given fields: ctx
func (_m *DB) DetectOffchainDataGaps(ctx context.Context) (map[uint64]uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DetectOffchainDataGaps")
	}

	var r0 map[uint64]uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[uint64]uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[uint64]uint64); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint64]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_DetectOffchainDataGaps_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DetectOffchainDataGaps'
type DB_DetectOffchainDataGaps_Call struct {
	*mock.Call
}

// DetectOffchainDataGaps is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DB_Expecter) DetectOffchainDataGaps(ctx interface{}) *DB_DetectOffchainDataGaps_Call {
	return &DB_DetectOffchainDataGaps_Call{Call: _e.mock.On("DetectOffchainDataGaps", ctx)}
}

func (_c *DB_DetectOffchainDataGaps_Call) Run(run func(ctx context.Context)) *DB_DetectOffchainDataGaps_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_DetectOffchainDataGaps_Call) Return(_a0 map[uint64]uint64, _a1 error) *DB_DetectOffchainDataGaps_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_DetectOffchainDataGaps_Call) RunAndReturn(run func(context.Context) (map[uint64]uint64, error)) *DB_DetectOffchainDataGaps_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastProcessedBlock provides a mock function with given fields: ctx, task
func (_m *DB) GetLastProcessedBlock(ctx context.Context, task string) (uint64, error) {
	ret := _m.Called(ctx, task)

	if len(ret) == 0 {
		panic("no return value specified for GetLastProcessedBlock")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uint64, error)); ok {
		return rf(ctx, task)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetLastProcessedBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastProcessedBlock'
type DB_GetLastProcessedBlock_Call struct {
	*mock.Call
}

// GetLastProcessedBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - task string
func (_e *DB_Expecter) GetLastProcessedBlock(ctx interface{}, task interface{}) *DB_GetLastProcessedBlock_Call {
	return &DB_GetLastProcessedBlock_Call{Call: _e.mock.On("GetLastProcessedBlock", ctx, task)}
}

func (_c *DB_GetLastProcessedBlock_Call) Run(run func(ctx context.Context, task string)) *DB_GetLastProcessedBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_GetLastProcessedBlock_Call) Return(_a0 uint64, _a1 error) *DB_GetLastProcessedBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetLastProcessedBlock_Call) RunAndReturn(run func(context.Context, string) (uint64, error)) *DB_GetLastProcessedBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetOffChainData provides a mock function with given fields: ctx, key
func (_m *DB) GetOffChainData(ctx context.Context, key common.Hash) (*types.OffChainData, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for GetOffChainData")
	}

	var r0 *types.OffChainData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.OffChainData, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.OffChainData); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.OffChainData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetOffChainData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOffChainData'
type DB_GetOffChainData_Call struct {
	*mock.Call
}

// GetOffChainData is a helper method to define mock.On call
//   - ctx context.Context
//   - key common.Hash
func (_e *DB_Expecter) GetOffChainData(ctx interface{}, key interface{}) *DB_GetOffChainData_Call {
	return &DB_GetOffChainData_Call{Call: _e.mock.On("GetOffChainData", ctx, key)}
}

func (_c *DB_GetOffChainData_Call) Run(run func(ctx context.Context, key common.Hash)) *DB_GetOffChainData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *DB_GetOffChainData_Call) Return(_a0 *types.OffChainData, _a1 error) *DB_GetOffChainData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetOffChainData_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.OffChainData, error)) *DB_GetOffChainData_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnresolvedBatchKeys provides a mock function with given fields: ctx, limit
func (_m *DB) GetUnresolvedBatchKeys(ctx context.Context, limit uint) ([]types.BatchKey, error) {
	ret := _m.Called(ctx, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetUnresolvedBatchKeys")
	}

	var r0 []types.BatchKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) ([]types.BatchKey, error)); ok {
		return rf(ctx, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) []types.BatchKey); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.BatchKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetUnresolvedBatchKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnresolvedBatchKeys'
type DB_GetUnresolvedBatchKeys_Call struct {
	*mock.Call
}

// GetUnresolvedBatchKeys is a helper method to define mock.On call
//   - ctx context.Context
//   - limit uint
func (_e *DB_Expecter) GetUnresolvedBatchKeys(ctx interface{}, limit interface{}) *DB_GetUnresolvedBatchKeys_Call {
	return &DB_GetUnresolvedBatchKeys_Call{Call: _e.mock.On("GetUnresolvedBatchKeys", ctx, limit)}
}

func (_c *DB_GetUnresolvedBatchKeys_Call) Run(run func(ctx context.Context, limit uint)) *DB_GetUnresolvedBatchKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *DB_GetUnresolvedBatchKeys_Call) Return(_a0 []types.BatchKey, _a1 error) *DB_GetUnresolvedBatchKeys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetUnresolvedBatchKeys_Call) RunAndReturn(run func(context.Context, uint) ([]types.BatchKey, error)) *DB_GetUnresolvedBatchKeys_Call {
	_c.Call.Return(run)
	return _c
}

// ListOffChainData provides a mock function with given fields: ctx, keys
func (_m *DB) ListOffChainData(ctx context.Context, keys []common.Hash) ([]types.OffChainData, error) {
	ret := _m.Called(ctx, keys)

	if len(ret) == 0 {
		panic("no return value specified for ListOffChainData")
	}

	var r0 []types.OffChainData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []common.Hash) ([]types.OffChainData, error)); ok {
		return rf(ctx, keys)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []common.Hash) []types.OffChainData); ok {
		r0 = rf(ctx, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.OffChainData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []common.Hash) error); ok {
		r1 = rf(ctx, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_ListOffChainData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOffChainData'
type DB_ListOffChainData_Call struct {
	*mock.Call
}

// ListOffChainData is a helper method to define mock.On call
//   - ctx context.Context
//   - keys []common.Hash
func (_e *DB_Expecter) ListOffChainData(ctx interface{}, keys interface{}) *DB_ListOffChainData_Call {
	return &DB_ListOffChainData_Call{Call: _e.mock.On("ListOffChainData", ctx, keys)}
}

func (_c *DB_ListOffChainData_Call) Run(run func(ctx context.Context, keys []common.Hash)) *DB_ListOffChainData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]common.Hash))
	})
	return _c
}

func (_c *DB_ListOffChainData_Call) Return(_a0 []types.OffChainData, _a1 error) *DB_ListOffChainData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_ListOffChainData_Call) RunAndReturn(run func(context.Context, []common.Hash) ([]types.OffChainData, error)) *DB_ListOffChainData_Call {
	_c.Call.Return(run)
	return _c
}

// StoreLastProcessedBlock provides a mock function with given fields: ctx, block, task
func (_m *DB) StoreLastProcessedBlock(ctx context.Context, block uint64, task string) error {
	ret := _m.Called(ctx, block, task)

	if len(ret) == 0 {
		panic("no return value specified for StoreLastProcessedBlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) error); ok {
		r0 = rf(ctx, block, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_StoreLastProcessedBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreLastProcessedBlock'
type DB_StoreLastProcessedBlock_Call struct {
	*mock.Call
}

// StoreLastProcessedBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - block uint64
//   - task string
func (_e *DB_Expecter) StoreLastProcessedBlock(ctx interface{}, block interface{}, task interface{}) *DB_StoreLastProcessedBlock_Call {
	return &DB_StoreLastProcessedBlock_Call{Call: _e.mock.On("StoreLastProcessedBlock", ctx, block, task)}
}

func (_c *DB_StoreLastProcessedBlock_Call) Run(run func(ctx context.Context, block uint64, task string)) *DB_StoreLastProcessedBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(string))
	})
	return _c
}

func (_c *DB_StoreLastProcessedBlock_Call) Return(_a0 error) *DB_StoreLastProcessedBlock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_StoreLastProcessedBlock_Call) RunAndReturn(run func(context.Context, uint64, string) error) *DB_StoreLastProcessedBlock_Call {
	_c.Call.Return(run)
	return _c
}

// StoreOffChainData provides a mock function with given fields: ctx, od
func (_m *DB) StoreOffChainData(ctx context.Context, od []types.OffChainData) error {
	ret := _m.Called(ctx, od)

	if len(ret) == 0 {
		panic("no return value specified for StoreOffChainData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.OffChainData) error); ok {
		r0 = rf(ctx, od)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_StoreOffChainData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreOffChainData'
type DB_StoreOffChainData_Call struct {
	*mock.Call
}

// StoreOffChainData is a helper method to define mock.On call
//   - ctx context.Context
//   - od []types.OffChainData
func (_e *DB_Expecter) StoreOffChainData(ctx interface{}, od interface{}) *DB_StoreOffChainData_Call {
	return &DB_StoreOffChainData_Call{Call: _e.mock.On("StoreOffChainData", ctx, od)}
}

func (_c *DB_StoreOffChainData_Call) Run(run func(ctx context.Context, od []types.OffChainData)) *DB_StoreOffChainData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.OffChainData))
	})
	return _c
}

func (_c *DB_StoreOffChainData_Call) Return(_a0 error) *DB_StoreOffChainData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_StoreOffChainData_Call) RunAndReturn(run func(context.Context, []types.OffChainData) error) *DB_StoreOffChainData_Call {
	_c.Call.Return(run)
	return _c
}

// StoreUnresolvedBatchKeys provides a mock function with given fields: ctx, bks
func (_m *DB) StoreUnresolvedBatchKeys(ctx context.Context, bks []types.BatchKey) error {
	ret := _m.Called(ctx, bks)

	if len(ret) == 0 {
		panic("no return value specified for StoreUnresolvedBatchKeys")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.BatchKey) error); ok {
		r0 = rf(ctx, bks)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_StoreUnresolvedBatchKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreUnresolvedBatchKeys'
type DB_StoreUnresolvedBatchKeys_Call struct {
	*mock.Call
}

// StoreUnresolvedBatchKeys is a helper method to define mock.On call
//   - ctx context.Context
//   - bks []types.BatchKey
func (_e *DB_Expecter) StoreUnresolvedBatchKeys(ctx interface{}, bks interface{}) *DB_StoreUnresolvedBatchKeys_Call {
	return &DB_StoreUnresolvedBatchKeys_Call{Call: _e.mock.On("StoreUnresolvedBatchKeys", ctx, bks)}
}

func (_c *DB_StoreUnresolvedBatchKeys_Call) Run(run func(ctx context.Context, bks []types.BatchKey)) *DB_StoreUnresolvedBatchKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.BatchKey))
	})
	return _c
}

func (_c *DB_StoreUnresolvedBatchKeys_Call) Return(_a0 error) *DB_StoreUnresolvedBatchKeys_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_StoreUnresolvedBatchKeys_Call) RunAndReturn(run func(context.Context, []types.BatchKey) error) *DB_StoreUnresolvedBatchKeys_Call {
	_c.Call.Return(run)
	return _c
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
