package histories

import (
	"errors"
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *history_repo {
	return &history_repo{db}
}

func (repo *history_repo) FindAllHistories() (*models.Histories, error) {
	var data models.Histories
	result := repo.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, name, created_at, updated_at")
		}).
		Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, created_at, updated_at")
		}).
		Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed to get data history")
	}
	if len(data) == 0 {
		return nil, errors.New("data history is empty")
	}

	return &data, nil
}

func (repo *history_repo) FindHistory(r *http.Request) (*models.Histories, error) {
	var data models.Histories
	var result *gorm.DB

	varsVehicle := r.URL.Query().Get("vehicle_id")
	varsUser := r.URL.Query().Get("user_id")

	if varsVehicle != "" {
		result = repo.db.Order("created_at desc").Where("vehicle_id = ?", varsVehicle).Find(&data)
	}
	if varsUser != "" {
		result = repo.db.Order("created_at desc").Where("user_id = ?", varsUser).Find(&data)
	}
	if result.Error != nil {
		return nil, errors.New("failed to get data vehicle")
	}
	if len(data) == 0 {
		return nil, errors.New("history not found")
	}

	return &data, nil
}

func (repo *history_repo) SaveHistory(data *models.History) (*models.History, error) {
	var checkUser models.Users
	repo.db.Where("user_id = ?", data.UserID).Find(&checkUser)
	if len(checkUser) == 0 {
		return nil, errors.New("user does not exist")
	}

	var checkVehicle models.Vehicles
	repo.db.Where("vehicle_id = ?", data.VehicleID).Find(&checkVehicle)
	if len(checkVehicle) == 0 {
		return nil, errors.New("vehicle does not exist")
	}

	result := repo.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failed to add data history")
	}

	var dataVehicle models.Vehicle
	repo.db.First(&dataVehicle, "vehicle_id = ?", data.VehicleID)
	repo.db.Model(&dataVehicle).Where("vehicle_id = ?", data.VehicleID).Update("total_rented", dataVehicle.TotalRented+1)

	return data, nil
}

func (repo *history_repo) ChangeHistory(r *http.Request, data *models.History) (*models.History, error) {
	vars := mux.Vars(r)
	result := repo.db.Model(&data).Where("history_id = ?", vars["history_id"]).Updates(data)

	if result.Error != nil {
		return nil, errors.New("failed to update data")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("data history is not exist")
	}

	return data, nil
}

func (repo *history_repo) RemoveHistory(r *http.Request, data *models.History) (*models.History, error) {
	vars := mux.Vars(r)
	result := repo.db.Delete(data, vars["history_id"])

	if result.Error != nil {
		return nil, errors.New("failed to delete data history")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("data history is not exist")
	}

	return data, nil
}
