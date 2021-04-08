package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thyago/sensor-monitor-go/daos"
	"github.com/thyago/sensor-monitor-go/services"
	"github.com/thyago/sensor-monitor-go/util"
)

var (
	sensorDataNumericService *services.SensorDataNumericService = nil
)

type SensorDataNumericResponseData struct {
	ID        uint64    `json:"id"`
	SensorID  uint64    `json:"sensor_id"`
	Data      float64   `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func (s *Server) getSensorDataNumericService() *services.SensorDataNumericService {
	if sensorDataNumericService == nil {
		sensorDataNumericService = services.NewSensorDataNumericService(daos.NewSensorDataNumericDAO(s.db))
	}
	return sensorDataNumericService
}

func (s *Server) ListSensorDataNumeric(c *gin.Context) {
	sID, err := getUrlParamUINT64(c, "sensor-id")
	if err != nil {
		return
	}
	timeLayout := "2006-01-02T15:04:05.999Z"

	// TODO: Improve parameter validation
	startTime, err := time.Parse(timeLayout, c.Query("start"))
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Invalid value for \"start\"", err)
		return
	}
	endTime, err := time.Parse(timeLayout, c.Query("end"))
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Invalid value for \"end\"", err)
		return
	}
	if startTime.Equal(endTime) || startTime.After(endTime) {
		setResponseError(c, http.StatusBadRequest, "\"start\" must be before \"end\"", err)
	}
	dimension, err := util.GetTimeDimension(c.Query("dimension"))
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Invalid value for \"dimension\". Use \"hour\", \"minute\" or \"second\"", err)
		return
	}

	sensorDataNumerics, err := s.getSensorDataNumericService().List(sID, startTime, endTime, dimension)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	items := make([]interface{}, len(sensorDataNumerics))
	for i := 0; i < len(sensorDataNumerics); i++ {
		items[i] = SensorDataNumericResponseData{
			ID:        sensorDataNumerics[i].ID,
			SensorID:  sensorDataNumerics[i].SensorID,
			Data:      sensorDataNumerics[i].Data,
			Timestamp: sensorDataNumerics[i].Timestamp,
		}
	}
	r := &ResponseDataArray{Data: items}
	c.JSON(http.StatusOK, r)
}

func (s *Server) RemoveAllSensorDataNumeric(c *gin.Context) {
	sID, err := getUrlParamUINT64(c, "sensor-id")
	if err != nil {
		return
	}
	err = s.getSensorDataNumericService().RemoveAll(sID)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
