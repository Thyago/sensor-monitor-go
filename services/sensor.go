package services

import (
	"github.com/thyago/sensor-monitor-go/models"
)

type SensorDAO interface {
	FindAll() ([]models.Sensor, error)
	FindByID(id uint64) (*models.Sensor, error)
	Save(dg *models.Sensor) error
	Delete(id uint64) error
}

type SensorService struct {
	dao SensorDAO
}

func NewSensorService(dao SensorDAO) *SensorService {
	return &SensorService{dao: dao}
}

func (s *SensorService) Create(model models.SensorModel, handle, name string) (*models.Sensor, error) {
	sensor := models.NewSensor(model, handle, name)
	err := s.dao.Save(sensor)
	if err != nil {
		return nil, err
	}
	return sensor, nil
}

func (s *SensorService) List() ([]models.Sensor, error) {
	return s.dao.FindAll()
}

func (s *SensorService) Update(sensorID uint64, name string) error {
	sensor, err := s.dao.FindByID(sensorID)
	if err != nil {
		return err
	}
	sensor.Name = name
	return s.dao.Save(sensor)
}

func (s *SensorService) Get(sensorID uint64) (*models.Sensor, error) {
	m, err := s.dao.FindByID(sensorID)
	return m, err
}

func (s *SensorService) Remove(sensorID uint64) error {
	return s.dao.Delete(sensorID)
}
