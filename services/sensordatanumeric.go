package services

import (
	"errors"
	"time"

	"github.com/thyago/sensor-monitor-go/clients"
	"github.com/thyago/sensor-monitor-go/config"
	"github.com/thyago/sensor-monitor-go/models"
	"github.com/thyago/sensor-monitor-go/util"
)

type SensorDataNumericDAO interface {
	FindByPeriod(sensorID uint64, startTime, endTime time.Time, dimension *util.TimeDimension) ([]models.SensorDataNumeric, error)
	FindLast(sensorID uint64) (*models.SensorDataNumeric, error)
	CreateMany(sdn []models.SensorDataNumeric) error
	DeleteAll(sensorID uint64) error
}

type SensorDataNumericService struct {
	dao SensorDataNumericDAO
}

func NewSensorDataNumericService(dao SensorDataNumericDAO) *SensorDataNumericService {
	return &SensorDataNumericService{dao: dao}
}

func (s *SensorDataNumericService) List(sensorID uint64, startTime, endTime time.Time, dimension *util.TimeDimension) ([]models.SensorDataNumeric, error) {
	return s.dao.FindByPeriod(sensorID, startTime, endTime, dimension)
}

func (s *SensorDataNumericService) RemoveAll(sensorID uint64) error {
	return s.dao.DeleteAll(sensorID)
}

func (s *SensorDataNumericService) Process(sensor *models.Sensor) error {
	if sensor.Model == models.SensorModelParin {
		return s.processParin(sensor)
	}
	return errors.New("Invalid sensor model")
}

func (s *SensorDataNumericService) processParin(sensor *models.Sensor) error {
	// Get last processed data
	lastSensorDataNumeric, err := s.dao.FindLast(sensor.ID)
	if err != nil && err != util.ErrNotFound {
		return err
	}

	// Retrieve sensor data
	c := clients.NewParinClient(config.Config.ParinAPIKEY, nil)
	sensorData, err := c.ListSensorData(sensor.Handle)
	if err != nil {
		return err
	}

	// Prepare data for insert
	nextIndex := 0
	sensorDataNumerics := make([]models.SensorDataNumeric, len(sensorData))
	for _, data := range sensorData {
		if lastSensorDataNumeric == nil || data.Timestamp.After(lastSensorDataNumeric.Timestamp) {
			sensorDataNumerics[nextIndex] = *models.NewSensorDataNumeric(sensor, data.Temperature, data.Timestamp)
			nextIndex++
		}
	}
	sensorDataNumerics = sensorDataNumerics[:nextIndex]

	// If nothing to insert, return
	if len(sensorDataNumerics) == 0 {
		return nil
	}

	// Insert data on DB
	err = s.dao.CreateMany(sensorDataNumerics)
	if err != nil {
		return err
	}

	return nil
}
