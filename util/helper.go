package util

import (
	"github.com/enbility/eebus-go/features"
	"github.com/enbility/eebus-go/service"
	"github.com/enbility/eebus-go/spine"
	"github.com/enbility/eebus-go/spine/model"
)

var PhaseNameMapping = []model.ElectricalConnectionPhaseNameType{model.ElectricalConnectionPhaseNameTypeA, model.ElectricalConnectionPhaseNameTypeB, model.ElectricalConnectionPhaseNameTypeC}

// check if the given usecase, actor is supported by the remote device
func IsUsecaseSupported(usecase model.UseCaseNameType, actor model.UseCaseActorType, remoteDevice *spine.DeviceRemoteImpl) bool {
	uci := remoteDevice.UseCaseManager().UseCaseInformation()
	for _, element := range uci {
		if *element.Actor != actor {
			continue
		}
		for _, uc := range element.UseCaseSupport {
			if uc.UseCaseName != nil && *uc.UseCaseName == usecase {
				return true
			}
		}
	}

	return false
}

// return the remote entity of a given type and device ski
func EntityOfTypeForSki(service *service.EEBUSService, entityType model.EntityTypeType, ski string) (*spine.EntityRemoteImpl, error) {
	rDevice := service.LocalDevice().RemoteDeviceForSki(ski)

	if rDevice == nil {
		return nil, features.ErrEntityNotFound
	}

	entities := rDevice.Entities()
	for _, entity := range entities {
		if entity.EntityType() == entityType {
			return entity, nil
		}
	}

	return nil, features.ErrEntityNotFound
}
