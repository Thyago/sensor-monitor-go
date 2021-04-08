package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thyago/sensor-monitor-go/daos"
	"github.com/thyago/sensor-monitor-go/models"
	"github.com/thyago/sensor-monitor-go/services"
	"github.com/thyago/sensor-monitor-go/util"
)

var (
	sensorService *services.SensorService = nil
)

type SensorResponseData struct {
	ID     uint64 `json:"id"`
	Model  string `json:"model"`
	Handle string `json:"handle"`
	Name   string `json:"name"`
}

func (s *Server) getSensorService() *services.SensorService {
	if sensorService == nil {
		sensorService = services.NewSensorService(daos.NewSensorDAO(s.db))
	}
	return sensorService
}

func (s *Server) CreateSensor(c *gin.Context) {
	req := &struct {
		Model  models.SensorModel
		Handle string
		Name   string
	}{}
	err := unmarshalBody(c, req)
	if err != nil {
		return
	}

	sensor, err := s.getSensorService().Create(req.Model, req.Handle, req.Name)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	r := ResponseData{SensorResponseData{
		ID:     sensor.ID,
		Model:  sensor.Model.String(),
		Handle: sensor.Handle,
		Name:   sensor.Name,
	}}
	c.JSON(http.StatusCreated, r)
}

func (s *Server) ListSensor(c *gin.Context) {
	sensors, err := s.getSensorService().List()
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	items := make([]interface{}, len(sensors))
	for i := 0; i < len(sensors); i++ {
		items[i] = SensorResponseData{
			ID:     sensors[i].ID,
			Model:  sensors[i].Model.String(),
			Handle: sensors[i].Handle,
			Name:   sensors[i].Name,
		}
	}
	r := &ResponseDataArray{Data: items}
	c.JSON(http.StatusOK, r)
}

func (s *Server) UpdateSensor(c *gin.Context) {
	sID, err := getUrlParamUINT64(c, "sensor-id")
	if err != nil {
		return
	}

	req := &struct{ Name string }{}
	err = unmarshalBody(c, req)
	if err != nil {
		return
	}

	err = s.getSensorService().Update(sID, req.Name)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Failed", err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (s *Server) GetSensor(c *gin.Context) {
	sID, err := getUrlParamUINT64(c, "sensor-id")
	if err != nil {
		return
	}

	sensor, err := s.getSensorService().Get(sID)
	if err != nil {
		if errors.Is(err, util.ErrNotFound) {
			setResponseError(c, http.StatusNotFound, "Not found", err)
		} else {
			setResponseError(c, http.StatusBadRequest, "Failed", err)
		}
		return
	}

	r := &ResponseData{&SensorResponseData{
		ID:     sensor.ID,
		Model:  sensor.Model.String(),
		Handle: sensor.Handle,
		Name:   sensor.Name,
	}}
	c.JSON(http.StatusOK, r)
}

func (s *Server) RemoveSensor(c *gin.Context) {
	sID, err := getUrlParamUINT64(c, "sensor-id")
	if err != nil {
		return
	}

	err = s.getSensorService().Remove(sID)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Failed", err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
