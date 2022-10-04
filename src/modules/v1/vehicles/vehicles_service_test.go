package vehicles

import (
	"fmt"
	"testing"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllVehicles(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Vehicles{
		{VehicleID: 1, Name: "Vespa"},
		{VehicleID: 2, Name: "Beat"},
	}

	repo.mock.On("GetAllVehicles").Return(&dataMock, nil)
	data := service.GetAllVehicles()

	result := data.Data.(*models.Vehicles)

	for i, v := range *result {
		assert.Equal(t, dataMock[i].VehicleID, v.VehicleID, "expect id from data mock")
	}
}

func TestGetPopularVehicle(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Vehicles{
		{VehicleID: 1, Name: "Vespa"},
		{VehicleID: 2, Name: "Beat"},
	}

	repo.mock.On("GetPopularVehicles").Return(&dataMock, nil)
	data := service.GetPopularVehicles()

	result := data.Data.(*models.Vehicles)

	for i, v := range *result {
		assert.Equal(t, dataMock[i].VehicleID, v.VehicleID, "expect id from data mock")
	}
}
func TestAddVehicle(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Vehicle{VehicleID: 1, Name: "Vespa"}

	repo.mock.On("AddVehicle", &dataMock).Return(&dataMock, nil)
	data := service.AddVehicle(&dataMock)

	result := data.Data.(*models.Vehicle)

	assert.Equal(t, dataMock.VehicleID, result.VehicleID, "expect id from data mock")
}

func TestSearchVehicles(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Vehicles{
		{VehicleID: 1, Name: "Vespa"},
		{VehicleID: 2, Name: "Beat"},
	}

	repo.mock.On("SearchVehicles", "vespa").Return(&dataMock, nil)
	data := service.SearchVehicles("vespa")

	result := data.Data.(*models.Vehicles)
	fmt.Println(result)

	for i, v := range *result {
		assert.Equal(t, dataMock[i].VehicleID, v.VehicleID, "expect id from data mock")
	}
}

func TestUpdateVehicle(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Vehicle{VehicleID: 1, Name: "Vespa"}

	repo.mock.On("UpdateVehicle", &dataMock, "1").Return(&dataMock, nil)
	data := service.UpdateVehicle(&dataMock, "1")

	assert.Equal(t, 200, data.Code, "expect id from data mock")
}

func TestDeleteVehicle(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Vehicle{VehicleID: 1, Name: "Vespa"}

	repo.mock.On("DeleteVehicle", &dataMock, "1").Return(&dataMock, nil)
	data := service.DeleteVehicle(&dataMock, "1")

	assert.Equal(t, 200, data.Code, "expect id from data mock")
}
