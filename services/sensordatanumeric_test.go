//+build test_all unit

package services_test

import (
	"testing"
	"time"

	"github.com/thyago/sensor-monitor-go/models"
	"github.com/thyago/sensor-monitor-go/services"
	"github.com/thyago/sensor-monitor-go/util"
)

func setupSensorDataNumeric() (*MockSensorDataNumericDAO, *services.SensorDataNumericService) {
	sdnDAO := &MockSensorDataNumericDAO{nextID: 1}
	s := services.NewSensorDataNumericService(sdnDAO)
	return sdnDAO, s
}

func Test_ListSensorDataNumeric(t *testing.T) {
	sdnDAO, s := setupSensorDataNumeric()

	// Empty
	sdns, err := s.List(1, time.Now(), time.Now(), util.TimeDimensionHour)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := len(sdns), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	dt1 := time.Now()
	dt2 := dt1.AddDate(0, 0, -1)
	sdnDAO.sdns = append(sdnDAO.sdns, models.SensorDataNumeric{ID: 1, SensorID: 1, Data: 84.3, Timestamp: dt1})
	sdnDAO.sdns = append(sdnDAO.sdns, models.SensorDataNumeric{ID: 2, SensorID: 1, Data: 82.3, Timestamp: dt2})
	sdns, err = s.List(1, dt2, dt1, util.TimeDimensionHour)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := len(sdns), 2; got != want {
		t.Errorf("got %v, wanted %v", got, want)
		return
	}
	if got, want := sdnDAO.sdns[0].ID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sdnDAO.sdns[0].SensorID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sdnDAO.sdns[0].Data, 84.3; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sdnDAO.sdns[0].Timestamp, dt1; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sdnDAO.sdns[1].ID, uint64(2); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sdnDAO.sdns[1].SensorID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sdnDAO.sdns[1].Data, 82.3; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := sdnDAO.sdns[1].Timestamp, dt2; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_RemoveAllSensorDataNumeric(t *testing.T) {
	sdnDAO, s := setupSensorDataNumeric()
	sdnDAO.sdns = append(sdnDAO.sdns, models.SensorDataNumeric{ID: 1, SensorID: 1, Data: 84.3, Timestamp: time.Now()})
	sdnDAO.sdns = append(sdnDAO.sdns, models.SensorDataNumeric{ID: 2, SensorID: 1, Data: 82.3, Timestamp: time.Now()})

	// Remove
	err := s.RemoveAll(1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}

	// Get
	if len(sdnDAO.sdns) > 0 {
		t.Error("Not removed")
	}
}

func Test_ProcessSensor(t *testing.T) {
	//TODO
}

// MOCK

type MockSensorDataNumericDAO struct {
	nextID uint64
	sdns   []models.SensorDataNumeric
}

func (m *MockSensorDataNumericDAO) FindByPeriod(sensorID uint64, startTime, endTime time.Time, dimension *util.TimeDimension) ([]models.SensorDataNumeric, error) {
	ret := []models.SensorDataNumeric{}
	for _, sensorDataNumeric := range m.sdns {
		if sensorDataNumeric.SensorID != sensorID || sensorDataNumeric.Timestamp.Before(startTime) || sensorDataNumeric.Timestamp.After(endTime) {
			continue
		}
		ret = append(ret, sensorDataNumeric)
	}
	return ret, nil
}

func (m *MockSensorDataNumericDAO) FindLast(sensorID uint64) (*models.SensorDataNumeric, error) {
	for i := len(m.sdns) - 1; i >= 0; i-- {
		if m.sdns[i].SensorID == sensorID {
			return &m.sdns[i], nil
		}
	}
	return nil, util.ErrNotFound
}

func (m *MockSensorDataNumericDAO) CreateMany(sdn []models.SensorDataNumeric) error {
	for _, sensorDataNumeric := range sdn {
		sensorDataNumeric.ID = m.nextID
		m.nextID++
		m.sdns = append(m.sdns, sensorDataNumeric)
	}
	return nil
}

func (m *MockSensorDataNumericDAO) DeleteAll(sensorID uint64) error {
	for i := 0; i < len(m.sdns); i++ {
		if m.sdns[i].SensorID == sensorID {
			m.sdns = append(m.sdns[:i], m.sdns[i+1:]...)
			i--
		}
	}
	return nil
}

func (m *MockSensorDataNumericDAO) GetNextID() uint64 {
	return m.nextID
}
