package routines

import (
	"fmt"

	"github.com/thyago/sensor-monitor-go/daos"
	"github.com/thyago/sensor-monitor-go/db"
	"github.com/thyago/sensor-monitor-go/models"
	"github.com/thyago/sensor-monitor-go/services"
)

func ProcessSensors(db *db.Database) error {
	fmt.Printf("Running background process sensors")
	// Retrieve active sensors
	sensorService := services.NewSensorService(daos.NewSensorDAO(db))
	sensors, err := sensorService.List()
	if err != nil {
		return err
	}

	// Process sensors
	sensorDataNumericService := services.NewSensorDataNumericService(daos.NewSensorDataNumericDAO(db))
	for _, sensor := range sensors {
		// Create go routine to process sensors in parallel
		go processSensor(sensorDataNumericService, &sensor)
	}

	return nil
}

func processSensor(service *services.SensorDataNumericService, sensor *models.Sensor) {
	fmt.Printf("Processing sensor: %v", sensor.ID)
	err := service.Process(sensor)
	if err != nil {
		fmt.Printf("Failed to process sensor %v: %v", sensor.ID, err)
	}
}
