package inverterpvvis

import (
	"github.com/enbility/eebus-go/features"
	"github.com/enbility/eebus-go/service"
	"github.com/enbility/eebus-go/spine"
	"github.com/enbility/eebus-go/spine/model"
	"github.com/enbility/eebus-go/util"
)

type InverterPVVisI interface {
	CurrentProductionPower() (float64, error)
	NominalPeakPower() (float64, error)
	TotalPVYield() (float64, error)
}

type InverterPVVisImpl struct {
	entity *spine.EntityLocalImpl

	service *service.EEBUSService

	inverterEntity               *spine.EntityRemoteImpl
	inverterDeviceConfiguration  *features.DeviceConfiguration
	inverterElectricalConnection *features.ElectricalConnection
	inverterMeasurement          *features.Measurement

	ski string
}

var _ InverterPVVisI = (*InverterPVVisImpl)(nil)

// Add InverterPVVis support
func NewInverterPVVis(service *service.EEBUSService, details *service.ServiceDetails) *InverterPVVisImpl {
	ski := util.NormalizeSKI(details.SKI)

	localEntity := service.LocalDevice().EntityForType(model.EntityTypeTypeCEM)

	inverter := &InverterPVVisImpl{
		service: service,
		entity:  localEntity,
		ski:     ski,
	}
	_ = spine.Events.Subscribe(inverter)

	service.RegisterRemoteSKI(ski, true)

	return inverter
}
