package vehicles

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) GetAllVehicles() (*models.Vehicles, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Vehicles), nil
}

func (m *RepoMock) GetPopularVehicles() (*models.Vehicles, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Vehicles), nil
}

func (m *RepoMock) SearchVehicles(name string) (*models.Vehicles, error) {
	args := m.mock.Called(name)
	return args.Get(0).(*models.Vehicles), nil
}

func (m *RepoMock) AddVehicle(data *models.Vehicle) (*models.Vehicle, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.Vehicle), nil
}

func (m *RepoMock) UpdateVehicle(data *models.Vehicle, id string) (*models.Vehicle, error) {
	args := m.mock.Called(data, id)
	return args.Get(0).(*models.Vehicle), nil
}

func (m *RepoMock) DeleteVehicle(data *models.Vehicle, id string) (*models.Vehicle, error) {
	args := m.mock.Called(data, id)
	return args.Get(0).(*models.Vehicle), nil
}

func (m *RepoMock) CheckVehicle(id string) bool {
	return true
}
