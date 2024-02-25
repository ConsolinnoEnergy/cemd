package util

import (
	eebusutil "github.com/enbility/eebus-go/util"
	"github.com/enbility/spine-go/model"
	"github.com/stretchr/testify/assert"
)

func (s *UtilSuite) Test_MeasurementValueForScope() {
	value, err := MeasurementValueForScope(s.service, s.mockRemoteEntity, model.ScopeTypeTypeACPower)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0.0, value)

	value, err = MeasurementValueForScope(s.service, s.monitoredEntity, model.ScopeTypeTypeACPower)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0.0, value)

	descData := &model.MeasurementDescriptionListDataType{
		MeasurementDescriptionData: []model.MeasurementDescriptionDataType{
			{
				MeasurementId: eebusutil.Ptr(model.MeasurementIdType(0)),
				ScopeType:     eebusutil.Ptr(model.ScopeTypeTypeACPower),
			},
		},
	}

	rFeature := s.remoteDevice.FeatureByEntityTypeAndRole(s.monitoredEntity, model.FeatureTypeTypeMeasurement, model.RoleTypeServer)
	fErr := rFeature.UpdateData(model.FunctionTypeMeasurementDescriptionListData, descData, nil, nil)
	assert.Nil(s.T(), fErr)

	value, err = MeasurementValueForScope(s.service, s.monitoredEntity, model.ScopeTypeTypeACPower)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0.0, value)

	data := &model.MeasurementListDataType{
		MeasurementData: []model.MeasurementDataType{
			{
				MeasurementId: eebusutil.Ptr(model.MeasurementIdType(0)),
				Value:         model.NewScaledNumberType(80),
			},
		},
	}

	fErr = rFeature.UpdateData(model.FunctionTypeMeasurementListData, data, nil, nil)
	assert.Nil(s.T(), fErr)

	value, err = MeasurementValueForScope(s.service, s.monitoredEntity, model.ScopeTypeTypeACPower)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 80.0, value)
}

func (s *UtilSuite) Test_MeasurementValuesForTypeCommodityScope() {
	measurementType := model.MeasurementTypeTypePower
	commodityType := model.CommodityTypeTypeElectricity
	scopeType := model.ScopeTypeTypeACPower
	energyDirection := model.EnergyDirectionTypeConsume

	data, err := MeasurementValuesForTypeCommodityScope(
		s.service,
		s.mockRemoteEntity,
		measurementType,
		commodityType,
		scopeType,
		energyDirection,
		PhaseNameMapping,
	)
	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), data)

	data, err = MeasurementValuesForTypeCommodityScope(
		s.service,
		s.monitoredEntity,
		measurementType,
		commodityType,
		scopeType,
		energyDirection,
		PhaseNameMapping,
	)
	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), data)

	descData := &model.MeasurementDescriptionListDataType{
		MeasurementDescriptionData: []model.MeasurementDescriptionDataType{
			{
				MeasurementId:   eebusutil.Ptr(model.MeasurementIdType(0)),
				MeasurementType: eebusutil.Ptr(model.MeasurementTypeTypePower),
				CommodityType:   eebusutil.Ptr(model.CommodityTypeTypeElectricity),
				ScopeType:       eebusutil.Ptr(model.ScopeTypeTypeACPower),
			},
			{
				MeasurementId:   eebusutil.Ptr(model.MeasurementIdType(1)),
				MeasurementType: eebusutil.Ptr(model.MeasurementTypeTypePower),
				CommodityType:   eebusutil.Ptr(model.CommodityTypeTypeElectricity),
				ScopeType:       eebusutil.Ptr(model.ScopeTypeTypeACPower),
			},
			{
				MeasurementId:   eebusutil.Ptr(model.MeasurementIdType(2)),
				MeasurementType: eebusutil.Ptr(model.MeasurementTypeTypePower),
				CommodityType:   eebusutil.Ptr(model.CommodityTypeTypeElectricity),
				ScopeType:       eebusutil.Ptr(model.ScopeTypeTypeACPower),
			},
		},
	}

	rFeature := s.remoteDevice.FeatureByEntityTypeAndRole(s.monitoredEntity, model.FeatureTypeTypeMeasurement, model.RoleTypeServer)
	fErr := rFeature.UpdateData(model.FunctionTypeMeasurementDescriptionListData, descData, nil, nil)
	assert.Nil(s.T(), fErr)

	data, err = MeasurementValuesForTypeCommodityScope(
		s.service,
		s.monitoredEntity,
		measurementType,
		commodityType,
		scopeType,
		energyDirection,
		PhaseNameMapping,
	)
	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), data)

	measData := &model.MeasurementListDataType{
		MeasurementData: []model.MeasurementDataType{
			{
				MeasurementId: eebusutil.Ptr(model.MeasurementIdType(0)),
				Value:         model.NewScaledNumberType(10),
			},
			{
				MeasurementId: eebusutil.Ptr(model.MeasurementIdType(1)),
				Value:         model.NewScaledNumberType(10),
			},
			{
				MeasurementId: eebusutil.Ptr(model.MeasurementIdType(2)),
				Value:         model.NewScaledNumberType(10),
			},
		},
	}

	fErr = rFeature.UpdateData(model.FunctionTypeMeasurementListData, measData, nil, nil)
	assert.Nil(s.T(), fErr)

	data, err = MeasurementValuesForTypeCommodityScope(
		s.service,
		s.monitoredEntity,
		measurementType,
		commodityType,
		scopeType,
		energyDirection,
		PhaseNameMapping,
	)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 0, len(data))

	elParamData := &model.ElectricalConnectionParameterDescriptionListDataType{
		ElectricalConnectionParameterDescriptionData: []model.ElectricalConnectionParameterDescriptionDataType{
			{
				ElectricalConnectionId: eebusutil.Ptr(model.ElectricalConnectionIdType(0)),
				MeasurementId:          eebusutil.Ptr(model.MeasurementIdType(0)),
				AcMeasuredPhases:       eebusutil.Ptr(model.ElectricalConnectionPhaseNameTypeA),
			},
			{
				ElectricalConnectionId: eebusutil.Ptr(model.ElectricalConnectionIdType(0)),
				MeasurementId:          eebusutil.Ptr(model.MeasurementIdType(1)),
				AcMeasuredPhases:       eebusutil.Ptr(model.ElectricalConnectionPhaseNameTypeB),
			},
			{
				ElectricalConnectionId: eebusutil.Ptr(model.ElectricalConnectionIdType(0)),
				MeasurementId:          eebusutil.Ptr(model.MeasurementIdType(2)),
				AcMeasuredPhases:       eebusutil.Ptr(model.ElectricalConnectionPhaseNameTypeC),
			},
		},
	}

	rElFeature := s.remoteDevice.FeatureByEntityTypeAndRole(s.monitoredEntity, model.FeatureTypeTypeElectricalConnection, model.RoleTypeServer)
	fErr = rElFeature.UpdateData(model.FunctionTypeElectricalConnectionParameterDescriptionListData, elParamData, nil, nil)
	assert.Nil(s.T(), fErr)

	elDescData := &model.ElectricalConnectionDescriptionListDataType{
		ElectricalConnectionDescriptionData: []model.ElectricalConnectionDescriptionDataType{
			{
				ElectricalConnectionId:  eebusutil.Ptr(model.ElectricalConnectionIdType(0)),
				PositiveEnergyDirection: eebusutil.Ptr(model.EnergyDirectionTypeConsume),
			},
		},
	}

	fErr = rElFeature.UpdateData(model.FunctionTypeElectricalConnectionDescriptionListData, elDescData, nil, nil)
	assert.Nil(s.T(), fErr)

	data, err = MeasurementValuesForTypeCommodityScope(
		s.service,
		s.monitoredEntity,
		measurementType,
		commodityType,
		scopeType,
		energyDirection,
		PhaseNameMapping,
	)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), []float64{10, 10, 10}, data)

}
