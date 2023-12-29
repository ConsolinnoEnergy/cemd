package inverterbatteryvis

import (
	"github.com/enbility/eebus-go/features"
	"github.com/enbility/eebus-go/service"
	"github.com/enbility/eebus-go/spine"
	"github.com/enbility/eebus-go/spine/model"
	"github.com/enbility/eebus-go/util"
)

type InverterBatteryVisI interface {
	CurrentDisChargePower() (float64, error)
	TotalChargeEnergy() (float64, error)
	TotalDischargeEnergy() (float64, error)
	CurrentStateOfCharge() (float64, error)
}

type InverterBatteryVisImpl struct {
	entity *spine.EntityLocalImpl

	service *service.EEBUSService

	inverterEntity               *spine.EntityRemoteImpl
	inverterElectricalConnection *features.ElectricalConnection
	inverterMeasurement          *features.Measurement

	ski string
}

var _ InverterBatteryVisI = (*InverterBatteryVisImpl)(nil)

// Add InverterBatteryVis support
func NewInverterBatteryVis(service *service.EEBUSService, details *service.ServiceDetails) *InverterBatteryVisImpl {
	ski := util.NormalizeSKI(details.SKI)

	localEntity := service.LocalDevice().EntityForType(model.EntityTypeTypeCEM)

	inverter := &InverterBatteryVisImpl{
		service: service,
		entity:  localEntity,
		ski:     ski,
	}
	_ = spine.Events.Subscribe(inverter)

	service.RegisterRemoteSKI(ski, true)

	return inverter
}
