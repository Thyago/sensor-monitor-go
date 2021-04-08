package daos

import (
	"time"

	"github.com/thyago/sensor-monitor-go/db"
	"github.com/thyago/sensor-monitor-go/models"
	"github.com/thyago/sensor-monitor-go/util"
)

type SensorDataNumericDAO struct {
	database *db.Database
}

func NewSensorDataNumericDAO(db *db.Database) *SensorDataNumericDAO {
	return &SensorDataNumericDAO{db}
}

// FindByPeriod retrieves a list of SensorDataNumeric for the specified "sensorID" where the recorded
// time is between "startTime" and "endTime".
// The result is grouped by specified "dimension" (hour, minute, second) and ordered by timestamp.
func (dao *SensorDataNumericDAO) FindByPeriod(sensorID uint64, startTime, endTime time.Time, dimension *util.TimeDimension) ([]models.SensorDataNumeric, error) {
	//TODO: Add pagination
	dORM := dao.database.GetORM()
	sdns := []models.SensorDataNumeric{}
	err := dORM.
		Debug().
		Model(&models.SensorDataNumeric{}).
		Select("*, MAX(data) as data").
		Where("sensor_id = ?", sensorID).
		Where("timestamp >= ? AND timestamp <= ?", startTime, endTime).
		Group(dimension.GroupBy("timestamp")).
		Order("timestamp ASC, id ASC").
		Find(&sdns).
		Error
	if err != nil {
		return nil, err
	}
	return sdns, nil
}

func (dao *SensorDataNumericDAO) FindLast(sensorID uint64) (*models.SensorDataNumeric, error) {
	dORM := dao.database.GetORM()
	sdns := &models.SensorDataNumeric{}
	err := dORM.
		Debug().
		Where("sensor_id = ?", sensorID).
		Order("timestamp DESC").
		Last(sdns).
		Error
	if err != nil {
		return nil, err
	}
	return sdns, nil
}

func (dao *SensorDataNumericDAO) CreateMany(sdn []models.SensorDataNumeric) error {
	dORM := dao.database.GetORM()
	err := dORM.Debug().Create(&sdn).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *SensorDataNumericDAO) DeleteAll(sensorID uint64) error {
	dORM := dao.database.GetORM()
	err := dORM.
		Debug().
		Where("sensor_id = ?", sensorID).
		Delete(&models.SensorDataNumeric{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
