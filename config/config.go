package config

import (
	"os"
	"strconv"
)

var Config appConfig

type appConfig struct {
	DBUser                    string
	DBPassword                string
	DBPort                    string
	DBHost                    string
	DBName                    string
	ServerPort                string
	ParinAPIKEY               string
	ParinSensorCheckFrequency int // seconds
}

func LoadConfig() {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	parinSensorCheckFrequency, err := strconv.ParseInt(os.Getenv("PARIN_SENSOR_CHECK_FREQUENCY"), 110, 16)
	if err != nil {
		parinSensorCheckFrequency = 10
	}

	Config = appConfig{
		DBUser:                    os.Getenv("DB_USER"),
		DBPassword:                os.Getenv("DB_PASSWORD"),
		DBPort:                    dbPort,
		DBHost:                    os.Getenv("DB_HOST"),
		DBName:                    os.Getenv("DB_NAME"),
		ServerPort:                serverPort,
		ParinAPIKEY:               os.Getenv("PARIN_API_KEY"),
		ParinSensorCheckFrequency: int(parinSensorCheckFrequency),
	}
}
