package routines

import (
	"fmt"
	"time"

	"github.com/thyago/sensor-monitor-go/config"
	"github.com/thyago/sensor-monitor-go/daos"
	"github.com/thyago/sensor-monitor-go/db"
	"github.com/thyago/sensor-monitor-go/models"
	"github.com/thyago/sensor-monitor-go/services"
)

type Scheduler struct {
	sensorService            *services.SensorService
	sensorDataNumericService *services.SensorDataNumericService
}

func NewScheduler(db *db.Database) *Scheduler {
	return &Scheduler{
		sensorService:            services.NewSensorService(daos.NewSensorDAO(db)),
		sensorDataNumericService: services.NewSensorDataNumericService(daos.NewSensorDataNumericDAO(db)),
	}
}

func (s *Scheduler) Run() {
	fmt.Printf("Running background process sensors")
	s.processSensors()
}

func (s *Scheduler) processSensors() error {
	// Retrieve active sensors
	sensors, err := s.sensorService.List()
	if err != nil {
		return err
	}

	// Process sensors every "ParinSensorCheckFrequency" seconds
	for {
		processSensorsJob(sensors, s.sensorDataNumericService)
		time.Sleep(time.Duration(config.Config.ParinSensorCheckFrequency) * time.Second)
	}
}

func processSensorsJob(sensors []models.Sensor, service *services.SensorDataNumericService) {
	for _, sensor := range sensors {
		fmt.Printf("Processing sensor: %v", sensor.ID)
		err := service.Process(&sensor)
		if err != nil {
			fmt.Printf("Failed to process sensor %v: %v", sensor.ID, err)
		}
	}
}
