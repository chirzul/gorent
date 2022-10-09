package vehicles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
	"github.com/chirzul/gorent/src/libs"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type vehicle_ctrl struct {
	svc interfaces.VehicleService
}

func NewCtrl(s interfaces.VehicleService) *vehicle_ctrl {
	return &vehicle_ctrl{svc: s}
}

func (c *vehicle_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	c.svc.GetAllVehicles().Send(w)
}

func (c *vehicle_ctrl) GetPopularVehicles(w http.ResponseWriter, r *http.Request) {
	c.svc.GetPopularVehicles().Send(w)
}

func (c *vehicle_ctrl) SearchVehicles(w http.ResponseWriter, r *http.Request) {
	vars := strings.ToLower(r.URL.Query().Get("name"))
	c.svc.SearchVehicles(vars).Send(w)
}

func (c *vehicle_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {
	var datas models.Vehicle
	imageName := r.Context().Value("imageName")
	datas.Image = fmt.Sprint(imageName)

	err := schema.NewDecoder().Decode(&datas, r.MultipartForm.Value)
	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}
	c.svc.AddVehicle(&datas).Send(w)
}

func (c *vehicle_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var datas *models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}
	c.svc.UpdateVehicle(datas, vars["vehicle_id"]).Send(w)
}

func (c *vehicle_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	var datas *models.Vehicle
	vars := mux.Vars(r)
	c.svc.DeleteVehicle(datas, vars["vehicle_id"]).Send(w)
}
