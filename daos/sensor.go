package daos

import (
	"github.com/thyago/sensor-monitor-go/db"
	"github.com/thyago/sensor-monitor-go/models"
)

type SensorDAO struct {
	database *db.Database
}

func NewSensorDAO(db *db.Database) *SensorDAO {
	return &SensorDAO{db}
}

func (dao *SensorDAO) FindAll() ([]models.Sensor, error) {
	// TODO: Add pagination

	dORM := dao.database.GetORM()
	sensors := []models.Sensor{}
	err := dORM.Debug().Find(&sensors).Limit(100).Error
	if err != nil {
		return nil, err
	}
	return sensors, nil
}

func (dao *SensorDAO) FindByID(id uint64) (*models.Sensor, error) {
	dORM := dao.database.GetORM()
	sensor := &models.Sensor{}
	err := dORM.Debug().Where("id = ?", id).Take(sensor).Error
	if err != nil {
		return nil, err
	}
	return sensor, nil
}

func (dao *SensorDAO) Save(dg *models.Sensor) error {
	dORM := dao.database.GetORM()
	err := dORM.Debug().Save(dg).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *SensorDAO) Delete(id uint64) error {
	dORM := dao.database.GetORM()
	err := dORM.Debug().Where("id = ?", id).Delete(&models.Sensor{}).Error
	if err != nil {
		return err
	}
	return nil
}
