package ucevcc

import (
	"github.com/enbility/cemd/api"
	"github.com/enbility/cemd/util"
	"github.com/enbility/eebus-go/features"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

// return if an EV is connected
//
// this includes all required features and
// minimal data being available
func (e *UCEvCC) EVConnected(entity spineapi.EntityRemoteInterface) bool {
	if entity == nil || entity.Device() == nil {
		return false
	}

	remoteDevice := e.service.LocalDevice().RemoteDeviceForSki(entity.Device().Ski())
	if remoteDevice == nil {
		return false
	}

	// check if the device still has an entity assigned with the provided entities address
	return remoteDevice.Entity(entity.Address().Entity) == entity
}

func (e *UCEvCC) deviceConfigurationValueForKeyName(
	entity spineapi.EntityRemoteInterface,
	keyname model.DeviceConfigurationKeyNameType,
	valueType model.DeviceConfigurationKeyValueTypeType) (any, error) {
	if entity == nil || entity.EntityType() != model.EntityTypeTypeEV {
		return nil, api.ErrNoEvEntity
	}

	evDeviceConfiguration, err := util.DeviceConfiguration(e.service, entity)
	if err != nil {
		return nil, features.ErrDataNotAvailable
	}

	// check if device configuration descriptions has an communication standard key name
	_, err = evDeviceConfiguration.GetDescriptionForKeyName(keyname)
	if err != nil {
		return nil, err
	}

	data, err := evDeviceConfiguration.GetKeyValueForKeyName(keyname, valueType)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, features.ErrDataNotAvailable
	}

	return data, nil
}

// return the current communication standard type used to communicate between EVSE and EV
//
// if an EV is connected via IEC61851, no ISO15118 specific data can be provided!
// sometimes the connection starts with IEC61851 before it switches
// to ISO15118, and sometimes it falls back again. so the error return is
// never absolut for the whole connection time, except if the use case
// is not supported
//
// the values are not constant and can change due to communication problems, bugs, and
// sometimes communication starts with IEC61851 before it switches to ISO
//
// possible errors:
//   - ErrDataNotAvailable if that information is not (yet) available
//   - ErrNotSupported if getting the communication standard is not supported
//   - and others
func (e *UCEvCC) EVCommunicationStandard(entity spineapi.EntityRemoteInterface) (string, error) {
	unknown := UcEVCCUnknownCommunicationStandard

	data, err := e.deviceConfigurationValueForKeyName(entity, model.DeviceConfigurationKeyNameTypeCommunicationsStandard, model.DeviceConfigurationKeyValueTypeTypeString)
	if err != nil {
		return unknown, err
	}

	if data == nil {
		return unknown, features.ErrDataNotAvailable
	}

	value := data.(*model.DeviceConfigurationKeyValueStringType)

	return string(*value), nil
}

// return if the EV supports asymmetric charging
//
// possible errors:
//   - ErrDataNotAvailable if that information is not (yet) available
func (e *UCEvCC) EVAsymmetricChargingSupported(entity spineapi.EntityRemoteInterface) (bool, error) {
	data, err := e.deviceConfigurationValueForKeyName(entity, model.DeviceConfigurationKeyNameTypeAsymmetricChargingSupported, model.DeviceConfigurationKeyValueTypeTypeBoolean)
	if err != nil {
		return false, err
	}

	if data == nil {
		return false, features.ErrDataNotAvailable
	}

	value := data.(*bool)

	return bool(*value), nil
}

// return the identifications of the currently connected EV or nil if not available
//
// possible errors:
//   - ErrDataNotAvailable if that information is not (yet) available
//   - and others
func (e *UCEvCC) EVIdentifications(entity spineapi.EntityRemoteInterface) ([]IdentificationItem, error) {
	if entity == nil || entity.EntityType() != model.EntityTypeTypeEV {
		return nil, api.ErrNoEvEntity
	}

	evIdentification, err := util.Identification(e.service, entity)
	if err != nil {
		return nil, features.ErrDataNotAvailable
	}

	identifications, err := evIdentification.GetValues()
	if err != nil {
		return nil, err
	}

	var ids []IdentificationItem
	for _, identification := range identifications {
		item := IdentificationItem{}

		typ := identification.IdentificationType
		if typ != nil {
			item.ValueType = *typ
		}

		value := identification.IdentificationValue
		if value != nil {
			item.Value = string(*value)
		}

		ids = append(ids, item)
	}

	return ids, nil
}

// the manufacturer data of an EVSE
// returns deviceName, serialNumber, error
func (e *UCEvCC) EVManufacturerData(
	entity spineapi.EntityRemoteInterface,
) (
	string,
	string,
	error,
) {
	deviceName := ""
	serialNumber := ""

	if entity == nil || entity.EntityType() != model.EntityTypeTypeEV {
		return deviceName, serialNumber, api.ErrNoEvEntity
	}

	evDeviceClassification, err := util.DeviceClassification(e.service, entity)
	if err != nil {
		return deviceName, serialNumber, features.ErrDataNotAvailable
	}

	data, err := evDeviceClassification.GetManufacturerDetails()
	if err != nil {
		return deviceName, serialNumber, err
	}

	if data.DeviceName != nil {
		deviceName = string(*data.DeviceName)
	}

	if data.SerialNumber != nil {
		serialNumber = string(*data.SerialNumber)
	}

	return deviceName, serialNumber, nil
}

// return the number of ac connected phases of the EV or 0 if it is unknown
func (e *UCEvCC) EVConnectedPhases(entity spineapi.EntityRemoteInterface) (uint, error) {
	if entity == nil || entity.EntityType() != model.EntityTypeTypeEV {
		return 0, api.ErrNoEvEntity
	}

	evElectricalConnection, err := util.ElectricalConnection(e.service, entity)
	if err != nil {
		return 0, features.ErrDataNotAvailable
	}

	data, err := evElectricalConnection.GetDescriptions()
	if err != nil {
		return 0, features.ErrDataNotAvailable
	}

	for _, item := range data {
		if item.ElectricalConnectionId != nil && item.AcConnectedPhases != nil {
			return *item.AcConnectedPhases, nil
		}
	}

	// default to 0 if the value is not available
	return 0, nil
}

// return the min, max, default limits for each phase of the connected EV
//
// possible errors:
//   - ErrDataNotAvailable if no such measurement is (yet) available
//   - and others
func (e *UCEvCC) EVCurrentLimits(entity spineapi.EntityRemoteInterface) ([]float64, []float64, []float64, error) {
	if entity == nil || entity.EntityType() != model.EntityTypeTypeEV {
		return nil, nil, nil, api.ErrNoEvEntity
	}

	evElectricalConnection, err := util.ElectricalConnection(e.service, entity)
	if err != nil {
		return nil, nil, nil, features.ErrDataNotAvailable
	}

	var resultMin, resultMax, resultDefault []float64

	for _, phaseName := range util.PhaseNameMapping {
		// electricalParameterDescription contains the measured phase for each measurementId
		elParamDesc, err := evElectricalConnection.GetParameterDescriptionForMeasuredPhase(phaseName)
		if err != nil || elParamDesc.ParameterId == nil {
			continue
		}

		dataMin, dataMax, dataDefault, err := evElectricalConnection.GetLimitsForParameterId(*elParamDesc.ParameterId)
		if err != nil {
			continue
		}

		// Min current data should be derived from min power data
		// but as this value is only properly provided via VAS the
		// currrent min values can not be trusted.

		resultMin = append(resultMin, dataMin)
		resultMax = append(resultMax, dataMax)
		resultDefault = append(resultDefault, dataDefault)
	}

	if len(resultMin) == 0 {
		return nil, nil, nil, features.ErrDataNotAvailable
	}

	return resultMin, resultMax, resultDefault, nil
}

// is the EV in sleep mode
// returns operatingState, lastErrorCode, error
func (e *UCEvCC) EVInSleepMode(
	entity spineapi.EntityRemoteInterface,
) (bool, error) {
	if entity == nil || entity.EntityType() != model.EntityTypeTypeEV {
		return false, api.ErrNoEvseEntity
	}

	evseDeviceDiagnosis, err := util.DeviceDiagnosis(e.service, entity)
	if err != nil {
		return false, err
	}

	data, err := evseDeviceDiagnosis.GetState()
	if err != nil {
		return false, err
	}

	if data.OperatingState != nil &&
		*data.OperatingState == model.DeviceDiagnosisOperatingStateTypeStandby {
		return true, nil
	}

	return false, nil
}
