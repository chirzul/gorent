package histories

import (
	"encoding/json"
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
	"github.com/chirzul/gorent/src/libs"
	"github.com/gorilla/mux"
)

type history_ctrl struct {
	svc interfaces.HistoryService
}

func NewCtrl(s interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{svc: s}
}

func (c *history_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	c.svc.GetAllHistories().Send(w)
}

func (c *history_ctrl) SearchHistory(w http.ResponseWriter, r *http.Request) {
	varsVehicle := r.URL.Query().Get("vehicle_id")
	varsUser := r.URL.Query().Get("user_id")
	query := map[string]string{
		"vehicle_id": varsVehicle,
		"user_id":    varsUser,
	}
	c.svc.SearchHistory(query).Send(w)
}

func (c *history_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {
	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}
	c.svc.AddHistory(&datas).Send(w)
}

func (c *history_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var datas *models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}
	c.svc.UpdateHistory(datas, vars["history_id"]).Send(w)
}

func (c *history_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	var datas *models.History
	vars := mux.Vars(r)
	c.svc.DeleteHistory(datas, vars["history_id"]).Send(w)
}
