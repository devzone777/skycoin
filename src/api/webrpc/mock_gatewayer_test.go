// Code generated by mockery v1.0.0. DO NOT EDIT.

package webrpc

import cipher "github.com/skycoin/skycoin/src/cipher"
import coin "github.com/skycoin/skycoin/src/coin"
import historydb "github.com/skycoin/skycoin/src/visor/historydb"
import mock "github.com/stretchr/testify/mock"
import visor "github.com/skycoin/skycoin/src/visor"

// MockGatewayer is an autogenerated mock type for the Gatewayer type
type MockGatewayer struct {
	mock.Mock
}

// GetAddrUxOuts provides a mock function with given fields: addr
func (_m *MockGatewayer) GetAddrUxOuts(addr []cipher.Address) ([]historydb.UxOut, error) {
	ret := _m.Called(addr)

	var r0 []historydb.UxOut
	if rf, ok := ret.Get(0).(func([]cipher.Address) []historydb.UxOut); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]historydb.UxOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]cipher.Address) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlocks provides a mock function with given fields: vs
func (_m *MockGatewayer) GetBlocks(vs []uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(vs)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func([]uint64) []coin.SignedBlock); ok {
		r0 = rf(vs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint64) error); ok {
		r1 = rf(vs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlocksInRange provides a mock function with given fields: start, end
func (_m *MockGatewayer) GetBlocksInRange(start uint64, end uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(start, end)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(uint64, uint64) []coin.SignedBlock); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBlocks provides a mock function with given fields: num
func (_m *MockGatewayer) GetLastBlocks(num uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(num)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(uint64) []coin.SignedBlock); ok {
		r0 = rf(num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTimeNow provides a mock function with given fields:
func (_m *MockGatewayer) GetTimeNow() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetTransaction provides a mock function with given fields: txid
func (_m *MockGatewayer) GetTransaction(txid cipher.SHA256) (*visor.Transaction, error) {
	ret := _m.Called(txid)

	var r0 *visor.Transaction
	if rf, ok := ret.Get(0).(func(cipher.SHA256) *visor.Transaction); ok {
		r0 = rf(txid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*visor.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cipher.SHA256) error); ok {
		r1 = rf(txid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnspentOutputsSummary provides a mock function with given fields: filters
func (_m *MockGatewayer) GetUnspentOutputsSummary(filters []visor.OutputsFilter) (*visor.UnspentOutputsSummary, error) {
	ret := _m.Called(filters)

	var r0 *visor.UnspentOutputsSummary
	if rf, ok := ret.Get(0).(func([]visor.OutputsFilter) *visor.UnspentOutputsSummary); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*visor.UnspentOutputsSummary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]visor.OutputsFilter) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InjectBroadcastTransaction provides a mock function with given fields: tx
func (_m *MockGatewayer) InjectBroadcastTransaction(tx coin.Transaction) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(coin.Transaction) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
