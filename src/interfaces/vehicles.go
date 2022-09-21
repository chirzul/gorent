package interfaces

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
)

type VehicleRepo interface {
	FindAllVehicles() (*models.Vehicles, error)
	FindPopularVehicles() (*models.Vehicles, error)
	FindVehiclesByName(name string) (*models.Vehicles, error)
	SaveVehicle(data *models.Vehicle) (*models.Vehicle, error)
	ChangeVehicle(data *models.Vehicle, id string) (*models.Vehicle, error)
	RemoveVehicle(data *models.Vehicle, id string) (*models.Vehicle, error)
}

type VehicleService interface {
	GetAllVehicles() (*models.Vehicles, error)
	GetPopularVehicles() (*models.Vehicles, error)
	SearchVehicles(name string) (*models.Vehicles, error)
	AddVehicle(data *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(data *models.Vehicle, id string) (*models.Vehicle, error)
	DeleteVehicle(data *models.Vehicle, id string) (*models.Vehicle, error)
}
