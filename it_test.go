//+build test_all integration

package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"github.com/thyago/sensor-monitor-go/config"
	"github.com/thyago/sensor-monitor-go/controllers"
	"gorm.io/gorm"
)

type smTestSuite struct {
	suite.Suite
	dbConnectionStr string
	port            int
	dbConn          *gorm.DB
}

func TestSMTestSuite(t *testing.T) {
	suite.Run(t, &smTestSuite{})
}

func (s *smTestSuite) SetupSuite() {
	server := controllers.NewServer()

	serverReady := make(chan bool)
	go server.Run()
	<-serverReady
}

func (s *smTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

// SENSOR

func (s *smTestSuite) TestIntegration_CreateValidSensor() {
	//TODO

	/*reqStr := `{"name":"Sensor 1", "model": "parin", "handle": "1234"}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%v/v1/sensors", config.Config.ServerPort), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"data":{"name":"Sensor 1", "id": 1, "model": "parin", "handle": "1234"}}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()*/
}

func TestIntegration_CreateInvalidSensor(t *testing.T) {
	//TODO
}

func TestIntegration_ListSensor(t *testing.T) {
	//TODO
}

func TestIntegration_UpdateValidSensor(t *testing.T) {
	//TODO
}

func TestIntegration_UpdateInvalidSensor(t *testing.T) {
	//TODO
}

func TestIntegration_GetExistingSensor(t *testing.T) {
	//TODO
}

func TestIntegration_GetNonExistingSensor(t *testing.T) {
	//TODO
}

func TestIntegration_DeleteSensor(t *testing.T) {
	//TODO
}

// SENSOR DATA NUMERIC

func TestIntegration_ListSensorDataNumericValidQuery(t *testing.T) {
	//TODO
}

func TestIntegration_ListSensorDataNumericInvalidQuery(t *testing.T) {
	//TODO
}

func TestIntegration_DeleteAllSensorDataNumeric(t *testing.T) {
	//TODO
}
