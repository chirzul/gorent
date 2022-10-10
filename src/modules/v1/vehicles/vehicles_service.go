package vehicles

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
	"github.com/chirzul/gorent/src/libs"
)

type vehicle_service struct {
	repo interfaces.VehicleRepo
}

func NewService(repo interfaces.VehicleRepo) *vehicle_service {
	return &vehicle_service{repo}
}

func (s *vehicle_service) GetAllVehicles() *libs.Response {
	data, err := s.repo.GetAllVehicles()
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *vehicle_service) GetPopularVehicles() *libs.Response {
	data, err := s.repo.GetPopularVehicles()
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *vehicle_service) SearchVehicles(name string) *libs.Response {
	data, err := s.repo.SearchVehicles(name)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *vehicle_service) GetVehiclesByCategory(category string) *libs.Response {
	data, err := s.repo.GetVehiclesByCategory(category)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *vehicle_service) AddVehicle(data *models.Vehicle) *libs.Response {
	data, err := s.repo.AddVehicle(data)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *vehicle_service) UpdateVehicle(data *models.Vehicle, id string) *libs.Response {
	if checkVehicle := s.repo.CheckVehicle(id); !checkVehicle {
		return libs.GetResponse("vehicle not found", 404, true)
	}
	data, err := s.repo.UpdateVehicle(data, id)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *vehicle_service) DeleteVehicle(data *models.Vehicle, id string) *libs.Response {
	if checkVehicle := s.repo.CheckVehicle(id); !checkVehicle {
		return libs.GetResponse("vehicle not found", 404, true)
	}
	data, err := s.repo.DeleteVehicle(data, id)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(data, 200, false)
}
