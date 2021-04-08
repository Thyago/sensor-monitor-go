package models

import (
	"encoding/json"
	"errors"
)

type Sensor struct {
	ID     uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Model  SensorModel `gorm:"uniqueIndex:idx_model_handle;size:255;not null" validate:"required" json:"model"`
	Handle string      `gorm:"uniqueIndex:idx_model_handle;size:255;not null" validate:"required" json:"handle"`
	Name   string      `gorm:"size:255" json:"name"`
}

func NewSensor(model SensorModel, handle, name string) *Sensor {
	return &Sensor{
		Model:  model,
		Handle: handle,
		Name:   name,
	}
}

func (s *Sensor) UnmarshalJSON(data []byte) error {
	// Define a secondary type to prevent recursive call to json.Unmarshal
	type Aux Sensor
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch s.Model {
	case SensorModelParin:
		return nil
	default:
		s.Model = ""
		return errors.New("invalid value for Key")
	}
}
