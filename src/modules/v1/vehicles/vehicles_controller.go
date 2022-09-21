package vehicles

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/helpers"
	"github.com/chirzul/gorent/src/interfaces"
	"github.com/gorilla/mux"
)

type vehicle_ctrl struct {
	svc interfaces.VehicleService
}

func NewCtrl(s interfaces.VehicleService) *vehicle_ctrl {
	return &vehicle_ctrl{svc: s}
}

func (c *vehicle_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data, err := c.svc.GetAllVehicles()
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		helpers.Response(w, 200, "success get data", err, data)
	}
}

func (c *vehicle_ctrl) GetPopularVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data, err := c.svc.GetPopularVehicles()
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		helpers.Response(w, 200, "success get data", err, data)
	}
}

func (c *vehicle_ctrl) SearchVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := strings.ToLower(r.URL.Query().Get("name"))
	data, err := c.svc.SearchVehicles(vars)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		helpers.Response(w, 200, "success get data", err, data)
	}

}

func (c *vehicle_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		_, err := c.svc.AddVehicle(&datas)
		if err != nil {
			helpers.Response(w, 400, "", err)
		} else {
			helpers.Response(w, 201, "success add data", err)
		}
	}

}

func (c *vehicle_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas models.Vehicle
	vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		_, err := c.svc.UpdateVehicle(&datas, vars["vehicle_id"])
		if err != nil {
			helpers.Response(w, 400, "", err)
		} else {
			helpers.Response(w, 200, "success update data", err)
		}
	}

}

func (c *vehicle_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas *models.Vehicle
	vars := mux.Vars(r)
	_, err := c.svc.DeleteVehicle(datas, vars["vehicle_id"])
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		helpers.Response(w, 200, "success delete data", err)
	}
}
