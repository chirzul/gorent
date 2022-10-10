package interfaces

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/libs"
)

type VehicleRepo interface {
	GetAllVehicles() (*models.Vehicles, error)
	GetPopularVehicles() (*models.Vehicles, error)
	GetVehiclesByCategory(category string) (*models.Vehicles, error)
	SearchVehicles(name string) (*models.Vehicles, error)
	AddVehicle(data *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(data *models.Vehicle, id string) (*models.Vehicle, error)
	DeleteVehicle(data *models.Vehicle, id string) (*models.Vehicle, error)
	CheckVehicle(id string) bool
}

type VehicleService interface {
	GetAllVehicles() *libs.Response
	GetPopularVehicles() *libs.Response
	GetVehiclesByCategory(category string) *libs.Response
	SearchVehicles(name string) *libs.Response
	AddVehicle(data *models.Vehicle) *libs.Response
	UpdateVehicle(data *models.Vehicle, id string) *libs.Response
	DeleteVehicle(data *models.Vehicle, id string) *libs.Response
}
