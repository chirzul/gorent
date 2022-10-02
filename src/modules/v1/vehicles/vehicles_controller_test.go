package vehicles

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/libs"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repos = RepoMock{mock.Mock{}}
var service = NewService(&repos)
var ctrl = NewCtrl(service)

var dataMock = models.Vehicles{
	{VehicleID: 1, Name: "Vespa"},
	{VehicleID: 2, Name: "Beat"},
}

func TestCtrlGetAllVehicles(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	repos.mock.On("GetAllVehicles").Return(&dataMock, nil)
	mux.HandleFunc("/test/vehicles", ctrl.GetAllVehicles)

	mux.ServeHTTP(w, httptest.NewRequest("GET", "/test/vehicles", nil))

	var vehicles models.Vehicles
	response := libs.Response{
		Data: &vehicles,
	}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	fmt.Println(response.Data)

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, response.IsError, "isError must be false")
}
