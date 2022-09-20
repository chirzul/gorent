package vehicles

import (
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
)

type vehicle_service struct {
	repo interfaces.VehicleRepo
}

func NewService(repo interfaces.VehicleRepo) *vehicle_service {
	return &vehicle_service{repo}
}

func (s *vehicle_service) GetAllVehicles() (*models.Vehicles, error) {
	data, err := s.repo.FindAllVehicles()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *vehicle_service) GetPopularVehicles() (*models.Vehicles, error) {
	data, err := s.repo.FindPopularVehicles()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *vehicle_service) SearchVehicles(r *http.Request) (*models.Vehicles, error) {
	data, err := s.repo.FindVehiclesByName(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *vehicle_service) AddVehicle(data *models.Vehicle) (*models.Vehicle, error) {
	data, err := s.repo.SaveVehicle(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *vehicle_service) UpdateVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error) {
	data, err := s.repo.ChangeVehicle(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *vehicle_service) DeleteVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error) {
	data, err := s.repo.RemoveVehicle(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
