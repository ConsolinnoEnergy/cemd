package util

import (
	"fmt"
	"time"

	"github.com/DerAndereAndi/eebus-go/service"
	"github.com/DerAndereAndi/eebus-go/spine"
	"github.com/DerAndereAndi/eebus-go/spine/model"
)

type TimeSeriesSlotType struct {
	TimeSeriesSlotId uint
	PeriodStartTime  time.Duration
	PeriodEndTime    time.Duration
	Duration         time.Duration
	Value            float64
	ValueMin         float64
	ValueMax         float64
}

type TimeSeriesType struct {
	TimeSeriesId    uint
	PeriodStartTime time.Duration
	PeriodEndTime   time.Duration
	Slots           []TimeSeriesSlotType
}

type TimeSeriesDescriptionType struct {
	TimeSeriesId     uint
	Type             model.TimeSeriesTypeType    // Description
	MeasurementId    uint                        // Description
	IsWriteable      bool                        // Description
	IsUpdateRequired bool                        // Description
	Unit             model.UnitOfMeasurementType // Description
}

type TimeSeriesConstraintsType struct {
	TimeSeriesId      uint
	SlotCountMin      uint          // Constraints
	SlotCountMax      uint          // Constraints
	SlotDurationMin   time.Duration // Constraints
	SlotDurationMax   time.Duration // Constraints
	SlotDurationStep  time.Duration // Constraints
	EarliestStartTime time.Time     // Constraints
	LatestEndTime     time.Time     // Constraints
	SlotValueMin      float64       // Constraints
	SlotValueMax      float64       // Constraints
	SlotValueStep     float64       // Constraints
}

// request FunctionTypeTimeSeriesDescriptionListData from a remote entity
func RequestTimeSeriesDescription(service *service.EEBUSService, entity *spine.EntityRemoteImpl) error {
	featureLocal, featureRemote, err := service.GetLocalClientAndRemoteServerFeatures(model.FeatureTypeTypeTimeSeries, entity)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = requestData(featureLocal, featureRemote, model.FunctionTypeTimeSeriesDescriptionListData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// request FunctionTypeTimeSeriesConstraintsListData from a remote entity
func RequestTimeSeriesConstraints(service *service.EEBUSService, entity *spine.EntityRemoteImpl) error {
	featureLocal, featureRemote, err := service.GetLocalClientAndRemoteServerFeatures(model.FeatureTypeTypeTimeSeries, entity)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = requestData(featureLocal, featureRemote, model.FunctionTypeTimeSeriesConstraintsListData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// request FunctionTypeTimeSeriesListData from a remote device
func RequestTimeSeriesList(service *service.EEBUSService, entity *spine.EntityRemoteImpl) (*model.MsgCounterType, error) {
	featureLocal, featureRemote, err := service.GetLocalClientAndRemoteServerFeatures(model.FeatureTypeTypeTimeSeries, entity)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	msgCounter, err := requestData(featureLocal, featureRemote, model.FunctionTypeTimeSeriesListData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return msgCounter, nil
}

// return current values for Time Series
func GetTimeSeriesValues(service *service.EEBUSService, entity *spine.EntityRemoteImpl) ([]TimeSeriesType, error) {
	_, featureRemote, err := service.GetLocalClientAndRemoteServerFeatures(model.FeatureTypeTypeTimeSeries, entity)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rData := featureRemote.Data(model.FunctionTypeTimeSeriesListData)
	if rData == nil {
		return nil, ErrDataNotAvailable
	}

	data := rData.(*model.TimeSeriesListDataType)
	var resultSet []TimeSeriesType

	for _, item := range data.TimeSeriesData {
		if item.TimeSeriesId == nil {
			continue
		}

		result := TimeSeriesType{
			TimeSeriesId: uint(*item.TimeSeriesId),
		}

		if item.TimePeriod != nil {
			if item.TimePeriod.StartTime != nil {
				if value, err := model.GetDurationFromString(*item.TimePeriod.StartTime); err == nil {
					result.PeriodStartTime = value
				}
			}
			if item.TimePeriod.EndTime != nil {
				if value, err := model.GetDurationFromString(*item.TimePeriod.EndTime); err == nil {
					result.PeriodEndTime = value
				}
			}
		}

		var slots []TimeSeriesSlotType
		for _, slot := range item.TimeSeriesSlot {
			element := TimeSeriesSlotType{
				TimeSeriesSlotId: uint(*slot.TimeSeriesSlotId),
			}
			if slot.Value != nil {
				element.Value = slot.Value.GetValue()
			}
			if slot.MinValue != nil {
				element.ValueMin = slot.MinValue.GetValue()
			}
			if slot.MaxValue != nil {
				element.ValueMax = slot.MaxValue.GetValue()
			}
			if slot.TimePeriod != nil {
				if slot.TimePeriod.StartTime != nil {
					if value, err := model.GetDurationFromString(*slot.TimePeriod.StartTime); err == nil {
						element.PeriodStartTime = value
					}
				}
				if slot.TimePeriod.EndTime != nil {
					if value, err := model.GetDurationFromString(*slot.TimePeriod.EndTime); err == nil {
						element.PeriodEndTime = value
					}
				}
			}
			if slot.Duration != nil {
				if value, err := model.GetDurationFromString(*slot.Duration); err == nil {
					element.Duration = value
				}
			}

			slots = append(slots, element)
		}
		result.Slots = slots

		resultSet = append(resultSet, result)
	}

	return resultSet, nil
}

// return current description values for Time Series
func GetTimeSeriesDescriptionValues(service *service.EEBUSService, entity *spine.EntityRemoteImpl) ([]TimeSeriesDescriptionType, error) {
	_, featureRemote, err := service.GetLocalClientAndRemoteServerFeatures(model.FeatureTypeTypeTimeSeries, entity)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rData := featureRemote.Data(model.FunctionTypeTimeSeriesDescriptionListData)
	if rData == nil {
		return nil, ErrDataNotAvailable
	}

	data := rData.(*model.TimeSeriesDescriptionListDataType)
	var resultSet []TimeSeriesDescriptionType

	for _, item := range data.TimeSeriesDescriptionData {
		if item.TimeSeriesId == nil {
			continue
		}

		result := TimeSeriesDescriptionType{
			TimeSeriesId: uint(*item.TimeSeriesId),
		}

		if item.TimeSeriesType != nil {
			result.Type = *item.TimeSeriesType
		}
		if item.MeasurementId != nil {
			result.MeasurementId = uint(*item.MeasurementId)
		}
		if item.TimeSeriesWriteable != nil {
			result.IsWriteable = *item.TimeSeriesWriteable
		}
		if item.UpdateRequired != nil {
			result.IsUpdateRequired = *item.UpdateRequired
		}
		if item.Unit != nil {
			result.Unit = *item.Unit
		}

		resultSet = append(resultSet, result)
	}

	return resultSet, nil
}

// return current constraint values for Time Series
func GetTimeSeriesConstraintValues(service *service.EEBUSService, entity *spine.EntityRemoteImpl) ([]TimeSeriesConstraintsType, error) {
	_, featureRemote, err := service.GetLocalClientAndRemoteServerFeatures(model.FeatureTypeTypeTimeSeries, entity)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rData := featureRemote.Data(model.FunctionTypeTimeSeriesConstraintsListData)
	switch constraintsData := rData.(type) {
	case *model.TimeSeriesConstraintsListDataType:
		if constraintsData == nil {
			return nil, ErrDataNotAvailable
		}
	}

	data := rData.(*model.TimeSeriesConstraintsListDataType)
	var resultSet []TimeSeriesConstraintsType

	for _, item := range data.TimeSeriesConstraintsData {
		if item.TimeSeriesId == nil {
			continue
		}

		result := TimeSeriesConstraintsType{
			TimeSeriesId: uint(*item.TimeSeriesId),
		}

		if item.SlotCountMin != nil {
			result.SlotCountMin = uint(*item.SlotCountMin)
		}
		if item.SlotCountMax != nil {
			result.SlotCountMax = uint(*item.SlotCountMax)
		}
		if item.SlotDurationMin != nil {
			if value, err := model.GetDurationFromString(*item.SlotDurationMin); err == nil {
				result.SlotDurationMin = value
			}
		}
		if item.SlotDurationMax != nil {
			if value, err := model.GetDurationFromString(*item.SlotDurationMax); err == nil {
				result.SlotDurationMax = value
			}
		}
		if item.SlotDurationStepSize != nil {
			if value, err := model.GetDurationFromString(*item.SlotDurationStepSize); err == nil {
				result.SlotDurationStep = value
			}
		}
		if item.EarliestTimeSeriesStartTime != nil {
			if value, err := model.GetTimeFromString(string(*item.EarliestTimeSeriesStartTime)); err == nil {
				result.EarliestStartTime = value
			}
		}
		if item.LatestTimeSeriesEndTime != nil {
			if value, err := model.GetTimeFromString(string(*item.LatestTimeSeriesEndTime)); err == nil {
				result.LatestEndTime = value
			}
		}
		if item.SlotValueMin != nil {
			result.SlotValueMin = item.SlotValueMin.GetValue()
		}
		if item.SlotValueMax != nil {
			result.SlotValueMax = item.SlotValueMax.GetValue()
		}
		if item.SlotValueStepSize != nil {
			result.SlotValueStep = item.SlotValueStepSize.GetValue()
		}

		resultSet = append(resultSet, result)
	}

	return resultSet, nil
}
