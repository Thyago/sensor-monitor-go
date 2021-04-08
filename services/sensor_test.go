//+build test_all unit

package services_test

import (
	"testing"

	"github.com/thyago/sensor-monitor-go/models"
	"github.com/thyago/sensor-monitor-go/services"
	"github.com/thyago/sensor-monitor-go/util"
)

func setupSensor() (*MockSensorDAO, *services.SensorService) {
	sensorDAO := &MockSensorDAO{nextID: 1}
	s := services.NewSensorService(sensorDAO)
	return sensorDAO, s
}

func Test_CreateSensor(t *testing.T) {
	sensorDAO, s := setupSensor()
	sensor, err := s.Create(models.SensorModelParin, "anyhandle", "Sensor 1")
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := sensor.ID, sensorDAO.GetNextID()-1; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sensor.Model, models.SensorModelParin; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sensor.Handle, "anyhandle"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sensor.Name, "Sensor 1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_ListSensor(t *testing.T) {
	sensorDAO, s := setupSensor()

	// Empty
	sensors, err := s.List()
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := len(sensors), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	sensorDAO.sensors = append(sensorDAO.sensors, models.Sensor{ID: 1, Name: "Test 1", Handle: "handle1", Model: models.SensorModelParin})
	sensorDAO.sensors = append(sensorDAO.sensors, models.Sensor{ID: 2, Name: "Test 2", Handle: "handle2", Model: models.SensorModelParin})
	sensors, err = s.List()
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := len(sensors), 2; got != want {
		t.Errorf("got %v, wanted %v", got, want)
		return
	}
	if got, want := sensorDAO.sensors[0].ID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sensorDAO.sensors[0].Name, "Test 1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sensorDAO.sensors[0].Handle, "handle1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sensorDAO.sensors[0].Model, models.SensorModelParin; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sensorDAO.sensors[1].ID, uint64(2); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sensorDAO.sensors[1].Name, "Test 2"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sensorDAO.sensors[1].Model, models.SensorModelParin; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_GetSensor(t *testing.T) {
	sensorDAO, s := setupSensor()
	sensorDAO.sensors = append(sensorDAO.sensors, models.Sensor{ID: 1, Name: "Sensor 1", Handle: "anyhandle", Model: models.SensorModelParin})

	// Get non existing
	sensor, err := s.Get(2)
	if err != util.ErrNotFound {
		t.Errorf("Expected %v, got %v", util.ErrSelfLoop, err)
		return
	}
	if sensor != nil {
		t.Errorf("Expected nil, got %v", sensor)
	}

	// Get created
	getsensor, err := s.Get(1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := getsensor.ID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := getsensor.Model, models.SensorModelParin; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := getsensor.Handle, "anyhandle"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := getsensor.Name, "Sensor 1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_UpdateSensor(t *testing.T) {
	sensorDAO, s := setupSensor()
	sensorDAO.sensors = append(sensorDAO.sensors, models.Sensor{ID: 1, Name: "Sensor 1", Handle: "anyhandle", Model: models.SensorModelParin})

	// Update
	err := s.Update(1, "Sensor 2")
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}

	if got, want := sensorDAO.sensors[0].Name, "Sensor 2"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_RemoveSensor(t *testing.T) {
	sensorDAO, s := setupSensor()
	sensorDAO.sensors = append(sensorDAO.sensors, models.Sensor{ID: 1, Name: "Sensor 1", Handle: "anyhandle", Model: models.SensorModelParin})

	// Remove
	err := s.Remove(1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}

	// Get
	if len(sensorDAO.sensors) > 0 {
		t.Error("Not removed")
	}
}

// MOCK

type MockSensorDAO struct {
	nextID  uint64
	sensors []models.Sensor
}

func (m *MockSensorDAO) FindAll() ([]models.Sensor, error) {
	return m.sensors, nil
}

func (m *MockSensorDAO) FindByID(id uint64) (*models.Sensor, error) {
	for _, element := range m.sensors {
		if element.ID == id {
			return &element, nil
		}
	}
	return nil, util.ErrNotFound
}

func (m *MockSensorDAO) Save(sensor *models.Sensor) error {
	if sensor.ID != 0 {
		for index, element := range m.sensors {
			if element.ID == sensor.ID {
				m.sensors[index] = *sensor
				return nil
			}
		}
		return util.ErrNotFound
	}
	sensor.ID = m.nextID
	m.nextID++
	m.sensors = append(m.sensors, *sensor)
	return nil
}

func (m *MockSensorDAO) Delete(id uint64) error {
	for index, element := range m.sensors {
		if element.ID == id {
			m.sensors = append(m.sensors[:index], m.sensors[index+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockSensorDAO) GetNextID() uint64 {
	return m.nextID
}
