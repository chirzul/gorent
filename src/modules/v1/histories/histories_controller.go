package histories

import (
	"encoding/json"
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/helpers"
	"github.com/chirzul/gorent/src/interfaces"
)

type history_ctrl struct {
	svc interfaces.HistoryService
}

func NewCtrl(s interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{svc: s}
}

func (c *history_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data, err := c.svc.GetAllHistories()
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		helpers.Response(w, 200, "success get data", err, data)
	}
}

func (c *history_ctrl) SearchHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data, err := c.svc.SearchHistory(r)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		helpers.Response(w, 200, "success get data", err, data)
	}
}

func (c *history_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		_, err := c.svc.AddHistory(&datas)
		if err != nil {
			helpers.Response(w, 400, "", err)
		} else {
			helpers.Response(w, 201, "success add data", err)
		}
	}

}

func (c *history_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		_, err := c.svc.UpdateHistory(r, &datas)
		if err != nil {
			helpers.Response(w, 400, "", err)
		} else {
			helpers.Response(w, 200, "success update data", err)
		}
	}
}

func (c *history_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.History
	_, err := c.svc.DeleteHistory(r, &datas)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		helpers.Response(w, 200, "success delete data", err)
	}
}
