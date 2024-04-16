// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	cemdapi "github.com/enbility/cemd/api"
	api "github.com/enbility/spine-go/api"

	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"

	time "time"
)

// UCLPCInterface is an autogenerated mock type for the UCLPCInterface type
type UCLPCInterface struct {
	mock.Mock
}

type UCLPCInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *UCLPCInterface) EXPECT() *UCLPCInterface_Expecter {
	return &UCLPCInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *UCLPCInterface) AddFeatures() {
	_m.Called()
}

// UCLPCInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type UCLPCInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *UCLPCInterface_Expecter) AddFeatures() *UCLPCInterface_AddFeatures_Call {
	return &UCLPCInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *UCLPCInterface_AddFeatures_Call) Run(run func()) *UCLPCInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UCLPCInterface_AddFeatures_Call) Return() *UCLPCInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *UCLPCInterface_AddFeatures_Call) RunAndReturn(run func()) *UCLPCInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *UCLPCInterface) AddUseCase() {
	_m.Called()
}

// UCLPCInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type UCLPCInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *UCLPCInterface_Expecter) AddUseCase() *UCLPCInterface_AddUseCase_Call {
	return &UCLPCInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *UCLPCInterface_AddUseCase_Call) Run(run func()) *UCLPCInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UCLPCInterface_AddUseCase_Call) Return() *UCLPCInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *UCLPCInterface_AddUseCase_Call) RunAndReturn(run func()) *UCLPCInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// ConsumptionLimit provides a mock function with given fields: entity
func (_m *UCLPCInterface) ConsumptionLimit(entity api.EntityRemoteInterface) (cemdapi.LoadLimit, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for ConsumptionLimit")
	}

	var r0 cemdapi.LoadLimit
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (cemdapi.LoadLimit, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) cemdapi.LoadLimit); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(cemdapi.LoadLimit)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCLPCInterface_ConsumptionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConsumptionLimit'
type UCLPCInterface_ConsumptionLimit_Call struct {
	*mock.Call
}

// ConsumptionLimit is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCLPCInterface_Expecter) ConsumptionLimit(entity interface{}) *UCLPCInterface_ConsumptionLimit_Call {
	return &UCLPCInterface_ConsumptionLimit_Call{Call: _e.mock.On("ConsumptionLimit", entity)}
}

func (_c *UCLPCInterface_ConsumptionLimit_Call) Run(run func(entity api.EntityRemoteInterface)) *UCLPCInterface_ConsumptionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCLPCInterface_ConsumptionLimit_Call) Return(limit cemdapi.LoadLimit, resultErr error) *UCLPCInterface_ConsumptionLimit_Call {
	_c.Call.Return(limit, resultErr)
	return _c
}

func (_c *UCLPCInterface_ConsumptionLimit_Call) RunAndReturn(run func(api.EntityRemoteInterface) (cemdapi.LoadLimit, error)) *UCLPCInterface_ConsumptionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeConsumptionActivePowerLimit provides a mock function with given fields: entity
func (_m *UCLPCInterface) FailsafeConsumptionActivePowerLimit(entity api.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for FailsafeConsumptionActivePowerLimit")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeConsumptionActivePowerLimit'
type UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call struct {
	*mock.Call
}

// FailsafeConsumptionActivePowerLimit is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCLPCInterface_Expecter) FailsafeConsumptionActivePowerLimit(entity interface{}) *UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	return &UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call{Call: _e.mock.On("FailsafeConsumptionActivePowerLimit", entity)}
}

func (_c *UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call) Run(run func(entity api.EntityRemoteInterface)) *UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call) Return(_a0 float64, _a1 error) *UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call) RunAndReturn(run func(api.EntityRemoteInterface) (float64, error)) *UCLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeDurationMinimum provides a mock function with given fields: entity
func (_m *UCLPCInterface) FailsafeDurationMinimum(entity api.EntityRemoteInterface) (time.Duration, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for FailsafeDurationMinimum")
	}

	var r0 time.Duration
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (time.Duration, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) time.Duration); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCLPCInterface_FailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeDurationMinimum'
type UCLPCInterface_FailsafeDurationMinimum_Call struct {
	*mock.Call
}

// FailsafeDurationMinimum is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCLPCInterface_Expecter) FailsafeDurationMinimum(entity interface{}) *UCLPCInterface_FailsafeDurationMinimum_Call {
	return &UCLPCInterface_FailsafeDurationMinimum_Call{Call: _e.mock.On("FailsafeDurationMinimum", entity)}
}

func (_c *UCLPCInterface_FailsafeDurationMinimum_Call) Run(run func(entity api.EntityRemoteInterface)) *UCLPCInterface_FailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCLPCInterface_FailsafeDurationMinimum_Call) Return(_a0 time.Duration, _a1 error) *UCLPCInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCLPCInterface_FailsafeDurationMinimum_Call) RunAndReturn(run func(api.EntityRemoteInterface) (time.Duration, error)) *UCLPCInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// IsUseCaseSupported provides a mock function with given fields: remoteEntity
func (_m *UCLPCInterface) IsUseCaseSupported(remoteEntity api.EntityRemoteInterface) (bool, error) {
	ret := _m.Called(remoteEntity)

	if len(ret) == 0 {
		panic("no return value specified for IsUseCaseSupported")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (bool, error)); ok {
		return rf(remoteEntity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) bool); ok {
		r0 = rf(remoteEntity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(remoteEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCLPCInterface_IsUseCaseSupported_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsUseCaseSupported'
type UCLPCInterface_IsUseCaseSupported_Call struct {
	*mock.Call
}

// IsUseCaseSupported is a helper method to define mock.On call
//   - remoteEntity api.EntityRemoteInterface
func (_e *UCLPCInterface_Expecter) IsUseCaseSupported(remoteEntity interface{}) *UCLPCInterface_IsUseCaseSupported_Call {
	return &UCLPCInterface_IsUseCaseSupported_Call{Call: _e.mock.On("IsUseCaseSupported", remoteEntity)}
}

func (_c *UCLPCInterface_IsUseCaseSupported_Call) Run(run func(remoteEntity api.EntityRemoteInterface)) *UCLPCInterface_IsUseCaseSupported_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCLPCInterface_IsUseCaseSupported_Call) Return(_a0 bool, _a1 error) *UCLPCInterface_IsUseCaseSupported_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCLPCInterface_IsUseCaseSupported_Call) RunAndReturn(run func(api.EntityRemoteInterface) (bool, error)) *UCLPCInterface_IsUseCaseSupported_Call {
	_c.Call.Return(run)
	return _c
}

// PowerConsumptionNominalMax provides a mock function with given fields: entity
func (_m *UCLPCInterface) PowerConsumptionNominalMax(entity api.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for PowerConsumptionNominalMax")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCLPCInterface_PowerConsumptionNominalMax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PowerConsumptionNominalMax'
type UCLPCInterface_PowerConsumptionNominalMax_Call struct {
	*mock.Call
}

// PowerConsumptionNominalMax is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCLPCInterface_Expecter) PowerConsumptionNominalMax(entity interface{}) *UCLPCInterface_PowerConsumptionNominalMax_Call {
	return &UCLPCInterface_PowerConsumptionNominalMax_Call{Call: _e.mock.On("PowerConsumptionNominalMax", entity)}
}

func (_c *UCLPCInterface_PowerConsumptionNominalMax_Call) Run(run func(entity api.EntityRemoteInterface)) *UCLPCInterface_PowerConsumptionNominalMax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCLPCInterface_PowerConsumptionNominalMax_Call) Return(_a0 float64, _a1 error) *UCLPCInterface_PowerConsumptionNominalMax_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCLPCInterface_PowerConsumptionNominalMax_Call) RunAndReturn(run func(api.EntityRemoteInterface) (float64, error)) *UCLPCInterface_PowerConsumptionNominalMax_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *UCLPCInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// UCLPCInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type UCLPCInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *UCLPCInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *UCLPCInterface_UpdateUseCaseAvailability_Call {
	return &UCLPCInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *UCLPCInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *UCLPCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *UCLPCInterface_UpdateUseCaseAvailability_Call) Return() *UCLPCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *UCLPCInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *UCLPCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// UseCaseName provides a mock function with given fields:
func (_m *UCLPCInterface) UseCaseName() model.UseCaseNameType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UseCaseName")
	}

	var r0 model.UseCaseNameType
	if rf, ok := ret.Get(0).(func() model.UseCaseNameType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.UseCaseNameType)
	}

	return r0
}

// UCLPCInterface_UseCaseName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UseCaseName'
type UCLPCInterface_UseCaseName_Call struct {
	*mock.Call
}

// UseCaseName is a helper method to define mock.On call
func (_e *UCLPCInterface_Expecter) UseCaseName() *UCLPCInterface_UseCaseName_Call {
	return &UCLPCInterface_UseCaseName_Call{Call: _e.mock.On("UseCaseName")}
}

func (_c *UCLPCInterface_UseCaseName_Call) Run(run func()) *UCLPCInterface_UseCaseName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UCLPCInterface_UseCaseName_Call) Return(_a0 model.UseCaseNameType) *UCLPCInterface_UseCaseName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UCLPCInterface_UseCaseName_Call) RunAndReturn(run func() model.UseCaseNameType) *UCLPCInterface_UseCaseName_Call {
	_c.Call.Return(run)
	return _c
}

// WriteConsumptionLimit provides a mock function with given fields: entity, limit
func (_m *UCLPCInterface) WriteConsumptionLimit(entity api.EntityRemoteInterface, limit cemdapi.LoadLimit) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, limit)

	if len(ret) == 0 {
		panic("no return value specified for WriteConsumptionLimit")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface, cemdapi.LoadLimit) (*model.MsgCounterType, error)); ok {
		return rf(entity, limit)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface, cemdapi.LoadLimit) *model.MsgCounterType); ok {
		r0 = rf(entity, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface, cemdapi.LoadLimit) error); ok {
		r1 = rf(entity, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCLPCInterface_WriteConsumptionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteConsumptionLimit'
type UCLPCInterface_WriteConsumptionLimit_Call struct {
	*mock.Call
}

// WriteConsumptionLimit is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
//   - limit cemdapi.LoadLimit
func (_e *UCLPCInterface_Expecter) WriteConsumptionLimit(entity interface{}, limit interface{}) *UCLPCInterface_WriteConsumptionLimit_Call {
	return &UCLPCInterface_WriteConsumptionLimit_Call{Call: _e.mock.On("WriteConsumptionLimit", entity, limit)}
}

func (_c *UCLPCInterface_WriteConsumptionLimit_Call) Run(run func(entity api.EntityRemoteInterface, limit cemdapi.LoadLimit)) *UCLPCInterface_WriteConsumptionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface), args[1].(cemdapi.LoadLimit))
	})
	return _c
}

func (_c *UCLPCInterface_WriteConsumptionLimit_Call) Return(_a0 *model.MsgCounterType, _a1 error) *UCLPCInterface_WriteConsumptionLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCLPCInterface_WriteConsumptionLimit_Call) RunAndReturn(run func(api.EntityRemoteInterface, cemdapi.LoadLimit) (*model.MsgCounterType, error)) *UCLPCInterface_WriteConsumptionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFailsafeConsumptionActivePowerLimit provides a mock function with given fields: entity, value
func (_m *UCLPCInterface) WriteFailsafeConsumptionActivePowerLimit(entity api.EntityRemoteInterface, value float64) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, value)

	if len(ret) == 0 {
		panic("no return value specified for WriteFailsafeConsumptionActivePowerLimit")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface, float64) (*model.MsgCounterType, error)); ok {
		return rf(entity, value)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface, float64) *model.MsgCounterType); ok {
		r0 = rf(entity, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface, float64) error); ok {
		r1 = rf(entity, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFailsafeConsumptionActivePowerLimit'
type UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call struct {
	*mock.Call
}

// WriteFailsafeConsumptionActivePowerLimit is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
//   - value float64
func (_e *UCLPCInterface_Expecter) WriteFailsafeConsumptionActivePowerLimit(entity interface{}, value interface{}) *UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call {
	return &UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call{Call: _e.mock.On("WriteFailsafeConsumptionActivePowerLimit", entity, value)}
}

func (_c *UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call) Run(run func(entity api.EntityRemoteInterface, value float64)) *UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface), args[1].(float64))
	})
	return _c
}

func (_c *UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call) Return(_a0 *model.MsgCounterType, _a1 error) *UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call) RunAndReturn(run func(api.EntityRemoteInterface, float64) (*model.MsgCounterType, error)) *UCLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFailsafeDurationMinimum provides a mock function with given fields: entity, duration
func (_m *UCLPCInterface) WriteFailsafeDurationMinimum(entity api.EntityRemoteInterface, duration time.Duration) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, duration)

	if len(ret) == 0 {
		panic("no return value specified for WriteFailsafeDurationMinimum")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface, time.Duration) (*model.MsgCounterType, error)); ok {
		return rf(entity, duration)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface, time.Duration) *model.MsgCounterType); ok {
		r0 = rf(entity, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface, time.Duration) error); ok {
		r1 = rf(entity, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCLPCInterface_WriteFailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFailsafeDurationMinimum'
type UCLPCInterface_WriteFailsafeDurationMinimum_Call struct {
	*mock.Call
}

// WriteFailsafeDurationMinimum is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
//   - duration time.Duration
func (_e *UCLPCInterface_Expecter) WriteFailsafeDurationMinimum(entity interface{}, duration interface{}) *UCLPCInterface_WriteFailsafeDurationMinimum_Call {
	return &UCLPCInterface_WriteFailsafeDurationMinimum_Call{Call: _e.mock.On("WriteFailsafeDurationMinimum", entity, duration)}
}

func (_c *UCLPCInterface_WriteFailsafeDurationMinimum_Call) Run(run func(entity api.EntityRemoteInterface, duration time.Duration)) *UCLPCInterface_WriteFailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface), args[1].(time.Duration))
	})
	return _c
}

func (_c *UCLPCInterface_WriteFailsafeDurationMinimum_Call) Return(_a0 *model.MsgCounterType, _a1 error) *UCLPCInterface_WriteFailsafeDurationMinimum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCLPCInterface_WriteFailsafeDurationMinimum_Call) RunAndReturn(run func(api.EntityRemoteInterface, time.Duration) (*model.MsgCounterType, error)) *UCLPCInterface_WriteFailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// NewUCLPCInterface creates a new instance of UCLPCInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUCLPCInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UCLPCInterface {
	mock := &UCLPCInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
