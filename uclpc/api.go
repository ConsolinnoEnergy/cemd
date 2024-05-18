package uclpc

import (
	"time"

	"github.com/enbility/cemd/api"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

//go:generate mockery

// interface for the Limitation of Power Consumption UseCase
type UCLPCInterface interface {
	api.UseCaseInterface

	// Scenario 1

	// return the current consumption limit data
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//
	// return values:
	//   - limit: load limit data
	//
	// possible errors:
	//   - ErrDataNotAvailable if no such limit is (yet) available
	//   - and others
	ConsumptionLimit(entity spineapi.EntityRemoteInterface) (limit api.LoadLimit, resultErr error)

	// send new LoadControlLimits
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//   - limit: load limit data
	WriteConsumptionLimit(entity spineapi.EntityRemoteInterface, limit api.LoadLimit) (*model.MsgCounterType, error)

	// Scenario 2

	// return Failsafe limit for the consumed active (real) power of the
	// Controllable System. This limit becomes activated in "init" state or "failsafe state".
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//
	// return values:
	//   - positive values are used for consumption
	FailsafeConsumptionActivePowerLimit(entity spineapi.EntityRemoteInterface) (float64, error)

	// send new Failsafe Consumption Active Power Limit
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//   - value: the new limit in W
	WriteFailsafeConsumptionActivePowerLimit(entity spineapi.EntityRemoteInterface, value float64) (*model.MsgCounterType, error)

	// return minimum time the Controllable System remains in "failsafe state" unless conditions
	// specified in this Use Case permit leaving the "failsafe state"
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//
	// return values:
	//   - negative values are used for production
	FailsafeDurationMinimum(entity spineapi.EntityRemoteInterface) (time.Duration, error)

	// send new Failsafe Duration Minimum
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//   - duration: the duration, between 2h and 24h
	WriteFailsafeDurationMinimum(entity spineapi.EntityRemoteInterface, duration time.Duration) (*model.MsgCounterType, error)

	// Scenario 3

	// this is automatically covered by the SPINE implementation

	// Scenario 4

	// return nominal maximum active (real) power the Controllable System is
	// able to consume according to the device label or data sheet.
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	PowerConsumptionNominalMax(entity spineapi.EntityRemoteInterface) (float64, error)
}
