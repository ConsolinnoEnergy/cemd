package uccevc

import (
	"time"

	"github.com/enbility/cemd/api"
	eebusutil "github.com/enbility/eebus-go/util"
	"github.com/enbility/spine-go/model"
	"github.com/stretchr/testify/assert"
)

func (s *UCCEVCSuite) Test_EVChargeStrategy() {
	data := s.sut.ChargeStrategy(s.mockRemoteEntity)
	assert.Equal(s.T(), api.EVChargeStrategyTypeUnknown, data)

	data = s.sut.ChargeStrategy(s.evEntity)
	assert.Equal(s.T(), api.EVChargeStrategyTypeUnknown, data)

	descData := &model.DeviceConfigurationKeyValueDescriptionListDataType{
		DeviceConfigurationKeyValueDescriptionData: []model.DeviceConfigurationKeyValueDescriptionDataType{
			{
				KeyId:   eebusutil.Ptr(model.DeviceConfigurationKeyIdType(0)),
				KeyName: eebusutil.Ptr(model.DeviceConfigurationKeyNameTypeCommunicationsStandard),
			},
		},
	}

	rFeature := s.remoteDevice.FeatureByEntityTypeAndRole(s.evEntity, model.FeatureTypeTypeDeviceConfiguration, model.RoleTypeServer)
	fErr := rFeature.UpdateData(model.FunctionTypeDeviceConfigurationKeyValueDescriptionListData, descData, nil, nil)
	assert.Nil(s.T(), fErr)

	data = s.sut.ChargeStrategy(s.evEntity)
	assert.Equal(s.T(), api.EVChargeStrategyTypeUnknown, data)

	keyData := &model.DeviceConfigurationKeyValueListDataType{
		DeviceConfigurationKeyValueData: []model.DeviceConfigurationKeyValueDataType{
			{
				KeyId: eebusutil.Ptr(model.DeviceConfigurationKeyIdType(0)),
				Value: &model.DeviceConfigurationKeyValueValueType{
					String: eebusutil.Ptr(model.DeviceConfigurationKeyValueStringType(model.DeviceConfigurationKeyValueStringTypeISO151182ED2)),
				},
			},
		},
	}

	fErr = rFeature.UpdateData(model.FunctionTypeDeviceConfigurationKeyValueListData, keyData, nil, nil)
	assert.Nil(s.T(), fErr)

	data = s.sut.ChargeStrategy(s.evEntity)
	assert.Equal(s.T(), api.EVChargeStrategyTypeUnknown, data)

	timeDescData := &model.TimeSeriesDescriptionListDataType{
		TimeSeriesDescriptionData: []model.TimeSeriesDescriptionDataType{
			{
				TimeSeriesId:   eebusutil.Ptr(model.TimeSeriesIdType(0)),
				TimeSeriesType: eebusutil.Ptr(model.TimeSeriesTypeTypeSingleDemand),
			},
		},
	}

	rTimeFeature := s.remoteDevice.FeatureByEntityTypeAndRole(s.evEntity, model.FeatureTypeTypeTimeSeries, model.RoleTypeServer)
	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesDescriptionListData, timeDescData, nil, nil)
	assert.Nil(s.T(), fErr)

	timeData := &model.TimeSeriesListDataType{
		TimeSeriesData: []model.TimeSeriesDataType{
			{
				TimeSeriesId: eebusutil.Ptr(model.TimeSeriesIdType(0)),
			},
		},
	}

	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesListData, timeData, nil, nil)
	assert.Nil(s.T(), fErr)

	data = s.sut.ChargeStrategy(s.evEntity)
	assert.Equal(s.T(), api.EVChargeStrategyTypeUnknown, data)

	timeData = &model.TimeSeriesListDataType{
		TimeSeriesData: []model.TimeSeriesDataType{
			{
				TimeSeriesId: eebusutil.Ptr(model.TimeSeriesIdType(0)),
				TimeSeriesSlot: []model.TimeSeriesSlotType{
					{
						TimeSeriesSlotId: eebusutil.Ptr(model.TimeSeriesSlotIdType(0)),
					},
				},
			},
		},
	}

	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesListData, timeData, nil, nil)
	assert.Nil(s.T(), fErr)

	data = s.sut.ChargeStrategy(s.evEntity)
	assert.Equal(s.T(), api.EVChargeStrategyTypeNoDemand, data)

	timeData = &model.TimeSeriesListDataType{
		TimeSeriesData: []model.TimeSeriesDataType{
			{
				TimeSeriesId: eebusutil.Ptr(model.TimeSeriesIdType(0)),
				TimeSeriesSlot: []model.TimeSeriesSlotType{
					{
						TimeSeriesSlotId: eebusutil.Ptr(model.TimeSeriesSlotIdType(0)),
						Duration:         eebusutil.Ptr(model.DurationType("PT0S")),
						Value:            model.NewScaledNumberType(0),
					},
				},
			},
		},
	}

	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesListData, timeData, nil, nil)
	assert.Nil(s.T(), fErr)

	data = s.sut.ChargeStrategy(s.evEntity)
	assert.Equal(s.T(), api.EVChargeStrategyTypeNoDemand, data)

	timeData = &model.TimeSeriesListDataType{
		TimeSeriesData: []model.TimeSeriesDataType{
			{
				TimeSeriesId: eebusutil.Ptr(model.TimeSeriesIdType(0)),
				TimeSeriesSlot: []model.TimeSeriesSlotType{
					{
						TimeSeriesSlotId: eebusutil.Ptr(model.TimeSeriesSlotIdType(0)),
						Value:            model.NewScaledNumberType(10000),
					},
				},
			},
		},
	}

	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesListData, timeData, nil, nil)
	assert.Nil(s.T(), fErr)

	data = s.sut.ChargeStrategy(s.evEntity)
	assert.Equal(s.T(), api.EVChargeStrategyTypeDirectCharging, data)

	timeData = &model.TimeSeriesListDataType{
		TimeSeriesData: []model.TimeSeriesDataType{
			{
				TimeSeriesId: eebusutil.Ptr(model.TimeSeriesIdType(0)),
				TimeSeriesSlot: []model.TimeSeriesSlotType{
					{
						TimeSeriesSlotId: eebusutil.Ptr(model.TimeSeriesSlotIdType(0)),
						Value:            model.NewScaledNumberType(10000),
						Duration:         model.NewDurationType(2 * time.Hour),
					},
				},
			},
		},
	}

	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesListData, timeData, nil, nil)
	assert.Nil(s.T(), fErr)

	data = s.sut.ChargeStrategy(s.evEntity)
	assert.Equal(s.T(), api.EVChargeStrategyTypeTimedCharging, data)
}

func (s *UCCEVCSuite) Test_EVEnergySingleDemand() {
	demand, err := s.sut.EnergyDemand(s.mockRemoteEntity)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0.0, demand.MinDemand)
	assert.Equal(s.T(), 0.0, demand.OptDemand)
	assert.Equal(s.T(), 0.0, demand.MaxDemand)
	assert.Equal(s.T(), 0.0, demand.DurationUntilStart)
	assert.Equal(s.T(), 0.0, demand.DurationUntilEnd)

	demand, err = s.sut.EnergyDemand(s.evEntity)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0.0, demand.MinDemand)
	assert.Equal(s.T(), 0.0, demand.OptDemand)
	assert.Equal(s.T(), 0.0, demand.MaxDemand)
	assert.Equal(s.T(), 0.0, demand.DurationUntilStart)
	assert.Equal(s.T(), 0.0, demand.DurationUntilEnd)

	descData := &model.DeviceConfigurationKeyValueDescriptionListDataType{
		DeviceConfigurationKeyValueDescriptionData: []model.DeviceConfigurationKeyValueDescriptionDataType{
			{
				KeyId:   eebusutil.Ptr(model.DeviceConfigurationKeyIdType(0)),
				KeyName: eebusutil.Ptr(model.DeviceConfigurationKeyNameTypeCommunicationsStandard),
			},
		},
	}

	rFeature := s.remoteDevice.FeatureByEntityTypeAndRole(s.evEntity, model.FeatureTypeTypeDeviceConfiguration, model.RoleTypeServer)
	fErr := rFeature.UpdateData(model.FunctionTypeDeviceConfigurationKeyValueDescriptionListData, descData, nil, nil)
	assert.Nil(s.T(), fErr)

	demand, err = s.sut.EnergyDemand(s.evEntity)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0.0, demand.MinDemand)
	assert.Equal(s.T(), 0.0, demand.OptDemand)
	assert.Equal(s.T(), 0.0, demand.MaxDemand)
	assert.Equal(s.T(), 0.0, demand.DurationUntilStart)
	assert.Equal(s.T(), 0.0, demand.DurationUntilEnd)

	keyData := &model.DeviceConfigurationKeyValueListDataType{
		DeviceConfigurationKeyValueData: []model.DeviceConfigurationKeyValueDataType{
			{
				KeyId: eebusutil.Ptr(model.DeviceConfigurationKeyIdType(0)),
				Value: &model.DeviceConfigurationKeyValueValueType{
					String: eebusutil.Ptr(model.DeviceConfigurationKeyValueStringTypeISO151182ED2),
				},
			},
		},
	}

	fErr = rFeature.UpdateData(model.FunctionTypeDeviceConfigurationKeyValueListData, keyData, nil, nil)
	assert.Nil(s.T(), fErr)

	demand, err = s.sut.EnergyDemand(s.evEntity)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0.0, demand.MinDemand)
	assert.Equal(s.T(), 0.0, demand.OptDemand)
	assert.Equal(s.T(), 0.0, demand.MaxDemand)
	assert.Equal(s.T(), 0.0, demand.DurationUntilStart)
	assert.Equal(s.T(), 0.0, demand.DurationUntilEnd)

	timeDescData := &model.TimeSeriesDescriptionListDataType{
		TimeSeriesDescriptionData: []model.TimeSeriesDescriptionDataType{
			{
				TimeSeriesId:   eebusutil.Ptr(model.TimeSeriesIdType(0)),
				TimeSeriesType: eebusutil.Ptr(model.TimeSeriesTypeTypeSingleDemand),
			},
		},
	}

	rTimeFeature := s.remoteDevice.FeatureByEntityTypeAndRole(s.evEntity, model.FeatureTypeTypeTimeSeries, model.RoleTypeServer)
	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesDescriptionListData, timeDescData, nil, nil)
	assert.Nil(s.T(), fErr)

	timeData := &model.TimeSeriesListDataType{
		TimeSeriesData: []model.TimeSeriesDataType{
			{
				TimeSeriesId: eebusutil.Ptr(model.TimeSeriesIdType(0)),
			},
		},
	}

	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesListData, timeData, nil, nil)
	assert.Nil(s.T(), fErr)

	demand, err = s.sut.EnergyDemand(s.evEntity)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0.0, demand.MinDemand)
	assert.Equal(s.T(), 0.0, demand.OptDemand)
	assert.Equal(s.T(), 0.0, demand.MaxDemand)
	assert.Equal(s.T(), 0.0, demand.DurationUntilStart)
	assert.Equal(s.T(), 0.0, demand.DurationUntilEnd)

	timeData = &model.TimeSeriesListDataType{
		TimeSeriesData: []model.TimeSeriesDataType{
			{
				TimeSeriesId: eebusutil.Ptr(model.TimeSeriesIdType(0)),
				TimeSeriesSlot: []model.TimeSeriesSlotType{
					{
						TimeSeriesSlotId: eebusutil.Ptr(model.TimeSeriesSlotIdType(0)),
						TimePeriod: &model.TimePeriodType{
							StartTime: model.NewAbsoluteOrRelativeTimeType("PT0S"),
						},
					},
				},
			},
		},
	}

	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesListData, timeData, nil, nil)
	assert.Nil(s.T(), fErr)

	demand, err = s.sut.EnergyDemand(s.evEntity)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 0.0, demand.MinDemand)
	assert.Equal(s.T(), 0.0, demand.OptDemand)
	assert.Equal(s.T(), 0.0, demand.MaxDemand)
	assert.Equal(s.T(), 0.0, demand.DurationUntilStart)
	assert.Equal(s.T(), 0.0, demand.DurationUntilEnd)

	timeData = &model.TimeSeriesListDataType{
		TimeSeriesData: []model.TimeSeriesDataType{
			{
				TimeSeriesId: eebusutil.Ptr(model.TimeSeriesIdType(0)),
				TimePeriod: &model.TimePeriodType{
					StartTime: model.NewAbsoluteOrRelativeTimeType("PT0S"),
				},
				TimeSeriesSlot: []model.TimeSeriesSlotType{
					{
						TimeSeriesSlotId: eebusutil.Ptr(model.TimeSeriesSlotIdType(0)),
						MinValue:         model.NewScaledNumberType(1000),
						Value:            model.NewScaledNumberType(10000),
						MaxValue:         model.NewScaledNumberType(100000),
					},
				},
			},
		},
	}

	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesListData, timeData, nil, nil)
	assert.Nil(s.T(), fErr)

	demand, err = s.sut.EnergyDemand(s.evEntity)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1000.0, demand.MinDemand)
	assert.Equal(s.T(), 10000.0, demand.OptDemand)
	assert.Equal(s.T(), 100000.0, demand.MaxDemand)
	assert.Equal(s.T(), 0.0, demand.DurationUntilStart)
	assert.Equal(s.T(), 0.0, demand.DurationUntilEnd)

	timeData = &model.TimeSeriesListDataType{
		TimeSeriesData: []model.TimeSeriesDataType{
			{
				TimeSeriesId: eebusutil.Ptr(model.TimeSeriesIdType(0)),
				TimePeriod: &model.TimePeriodType{
					StartTime: model.NewAbsoluteOrRelativeTimeType("PT0S"),
				},
				TimeSeriesSlot: []model.TimeSeriesSlotType{
					{
						TimeSeriesSlotId: eebusutil.Ptr(model.TimeSeriesSlotIdType(0)),
						Value:            model.NewScaledNumberType(10000),
						Duration:         model.NewDurationType(2 * time.Hour),
					},
				},
			},
		},
	}

	fErr = rTimeFeature.UpdateData(model.FunctionTypeTimeSeriesListData, timeData, nil, nil)
	assert.Nil(s.T(), fErr)

	demand, err = s.sut.EnergyDemand(s.evEntity)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 0.0, demand.MinDemand)
	assert.Equal(s.T(), 10000.0, demand.OptDemand)
	assert.Equal(s.T(), 0.0, demand.MaxDemand)
	assert.Equal(s.T(), 0.0, demand.DurationUntilStart)
	assert.Equal(s.T(), time.Duration(2*time.Hour).Seconds(), demand.DurationUntilEnd)
}
