package interfaces

import (
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
)

type VehicleRepo interface {
	FindAllVehicles() (*models.Vehicles, error)
	FindVehiclesByName(r *http.Request) (*models.Vehicles, error)
	SaveVehicle(data *models.Vehicle) (*models.Vehicle, error)
	ChangeVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
	RemoveVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
}

type VehicleService interface {
	GetAllVehicles() (*models.Vehicles, error)
	SearchVehicles(r *http.Request) (*models.Vehicles, error)
	AddVehicle(data *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
	DeleteVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
}
