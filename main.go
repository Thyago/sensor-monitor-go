package main

import (
	"fmt"

	"github.com/thyago/sensor-monitor-go/config"
	"github.com/thyago/sensor-monitor-go/controllers"
	"github.com/thyago/sensor-monitor-go/db"
	"github.com/thyago/sensor-monitor-go/models"
	"github.com/thyago/sensor-monitor-go/routines"
)

func main() {
	// Load App Config
	config.LoadConfig()

	// Create DB connection
	db := db.NewDatabase(config.Config.DBUser, config.Config.DBPassword, config.Config.DBHost, config.Config.DBPort, config.Config.DBName)
	err := db.Open()
	if err != nil {
		fmt.Printf("Failed to run: %v", err)
		return
	}

	// Migrate DB if needed
	db.Migrate(
		&models.Sensor{},
		&models.SensorDataNumeric{},
	)

	// Create background scheduler to process active sensors
	sched := routines.NewScheduler(db)
	go sched.Run()

	// Create server
	serv := controllers.NewServer(db)
	serv.Run()
}
