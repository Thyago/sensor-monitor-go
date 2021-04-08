package models

type SensorModel string

const (
	SensorModelParin SensorModel = "parin"
)

func (sm *SensorModel) String() string {
	return string(*sm)
}
