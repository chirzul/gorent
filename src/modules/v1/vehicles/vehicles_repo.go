package vehicles

import (
	"errors"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"gorm.io/gorm"
)

type vehicle_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicle_repo {
	return &vehicle_repo{db}
}

func (repo *vehicle_repo) FindAllVehicles() (*models.Vehicles, error) {
	var data models.Vehicles

	result := repo.db.Order("created_at DESC").Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data vehicle")
	}
	if len(data) == 0 {
		return nil, errors.New("data vehicle is empty")
	}

	return &data, nil
}

func (repo *vehicle_repo) FindPopularVehicles() (*models.Vehicles, error) {
	var data models.Vehicles
	result := repo.db.Order("total_rented DESC").Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data vehicle")
	}
	if len(data) == 0 {
		return nil, errors.New("data vehicle is empty")
	}
	return &data, nil
}

func (repo *vehicle_repo) FindVehiclesByName(name string) (*models.Vehicles, error) {
	var data models.Vehicles
	result := repo.db.Order("created_at DESC").Where("LOWER(name) LIKE ?", "%"+name+"%").Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data vehicle")
	}
	if len(data) == 0 {
		return nil, errors.New("vehicle not found")
	}
	return &data, nil
}

func (repo *vehicle_repo) SaveVehicle(data *models.Vehicle) (*models.Vehicle, error) {
	result := repo.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failed to add data vehicle")
	}
	return data, nil
}

func (repo *vehicle_repo) ChangeVehicle(data *models.Vehicle, id string) (*models.Vehicle, error) {
	result := repo.db.Model(&data).Where("vehicle_id = ?", id).Updates(data)
	if result.Error != nil {
		return nil, errors.New("failed to update data vehicle")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("data vehicle is not exist")
	}
	return data, nil
}

func (repo *vehicle_repo) RemoveVehicle(data *models.Vehicle, id string) (*models.Vehicle, error) {
	result := repo.db.Delete(&data, id)
	if result.Error != nil {
		return nil, errors.New("failed to delete data vehicle")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("data vehicle is not exist")
	}
	return data, nil
}
