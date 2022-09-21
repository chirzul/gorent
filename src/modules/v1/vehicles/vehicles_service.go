package vehicles

import (
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

func (s *vehicle_service) SearchVehicles(name string) (*models.Vehicles, error) {
	data, err := s.repo.FindVehiclesByName(name)
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

func (s *vehicle_service) UpdateVehicle(data *models.Vehicle, id string) (*models.Vehicle, error) {
	data, err := s.repo.ChangeVehicle(data, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *vehicle_service) DeleteVehicle(data *models.Vehicle, id string) (*models.Vehicle, error) {
	data, err := s.repo.RemoveVehicle(data, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
