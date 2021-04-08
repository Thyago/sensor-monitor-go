package models

import "time"

type SensorDataNumeric struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	SensorID  uint64 `gorm:"required;index:idx_sensor_id_timestamp;not null" json:"sensor_id"`
	Sensor    *Sensor
	Data      float64   `gorm:"size:255" json:"data"`
	Timestamp time.Time `gorm:"required;index:idx_sensor_id_timestamp;not null" json:"timestamp"`
}

func NewSensorDataNumeric(s *Sensor, data float64, timestamp time.Time) *SensorDataNumeric {
	return &SensorDataNumeric{
		Sensor:    s,
		SensorID:  s.ID,
		Data:      data,
		Timestamp: timestamp,
	}
}
