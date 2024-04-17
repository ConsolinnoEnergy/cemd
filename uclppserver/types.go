package uclppserver

import "github.com/enbility/cemd/api"

const (
	// Load control obligation limit data updated
	//
	// The callback with this message provides:
	//   - the device of the e.g. SMGW
	//   - the entity of the e.g. SMGW
	//
	// Use Case LPC, Scenario 1
	DataUpdateLimit api.EventType = "DataUpdateLimit"

	// Failsafe limit for the produced active (real) power of the
	// Controllable System data updated
	//
	// The callback with this message provides:
	//   - the device of the e.g. SMGW
	//   - the entity of the e.g. SMGW
	//
	// Use Case LPC, Scenario 2
	//
	// Note: the referred data may be updated together with all other configuration items of this use case
	DataUpdateFailsafeProductionActivePowerLimit api.EventType = "DataUpdateFailsafeProductionActivePowerLimit"

	// Minimum time the Controllable System remains in "failsafe state" unless conditions
	// specified in this Use Case permit leaving the "failsafe state" data updated
	//
	// The callback with this message provides:
	//   - the device of the e.g. SMGW
	//   - the entity of the e.g. SMGW
	//
	// Use Case LPC, Scenario 2
	//
	// Note: the referred data may be updated together with all other configuration items of this use case
	DataUpdateFailsafeDurationMinimum api.EventType = "DataUpdateFailsafeDurationMinimum"
)
