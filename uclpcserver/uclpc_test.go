package uclpcserver

import (
	eebusutil "github.com/enbility/eebus-go/util"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/stretchr/testify/assert"
)

func (s *UCLPCServerSuite) Test_loadControlWriteCB() {
	msg := &spineapi.Message{}

	s.sut.loadControlWriteCB(msg)

	msg = &spineapi.Message{
		RequestHeader: &model.HeaderType{
			MsgCounter: eebusutil.Ptr(model.MsgCounterType(500)),
		},
	}

	s.sut.loadControlWriteCB(msg)
}

func (s *UCLPCServerSuite) Test_UpdateUseCaseAvailability() {
	s.sut.UpdateUseCaseAvailability(true)
}

func (s *UCLPCServerSuite) Test_IsUseCaseSupported() {
	data, err := s.sut.IsUseCaseSupported(s.mockRemoteEntity)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), false, data)

	data, err = s.sut.IsUseCaseSupported(s.monitoredEntity)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), false, data)

	ucData := &model.NodeManagementUseCaseDataType{
		UseCaseInformation: []model.UseCaseInformationDataType{
			{
				Actor: eebusutil.Ptr(model.UseCaseActorTypeEnergyGuard),
				UseCaseSupport: []model.UseCaseSupportType{
					{
						UseCaseName:      eebusutil.Ptr(model.UseCaseNameTypeLimitationOfPowerConsumption),
						UseCaseAvailable: eebusutil.Ptr(true),
						ScenarioSupport:  []model.UseCaseScenarioSupportType{1, 2, 3, 4},
					},
				},
			},
		},
	}

	nodemgmtEntity := s.remoteDevice.Entity([]model.AddressEntityType{0})
	nodeFeature := s.remoteDevice.FeatureByEntityTypeAndRole(nodemgmtEntity, model.FeatureTypeTypeNodeManagement, model.RoleTypeSpecial)
	fErr := nodeFeature.UpdateData(model.FunctionTypeNodeManagementUseCaseData, ucData, nil, nil)
	assert.Nil(s.T(), fErr)

	data, err = s.sut.IsUseCaseSupported(s.monitoredEntity)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), true, data)
}
