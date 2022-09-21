package users

import (
	"encoding/json"
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
	"github.com/chirzul/gorent/src/libs"
	"github.com/gorilla/mux"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(s interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: s}
}

func (c *user_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	c.svc.GetAllUsers().Send(w)
}

func (c *user_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	var datas *models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}
	c.svc.AddUser(datas).Send(w)
}

func (c *user_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var datas *models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}
	c.svc.UpdateUser(datas, vars["username"]).Send(w)
}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var datas *models.User
	vars := mux.Vars(r)
	c.svc.DeleteUser(datas, vars["username"]).Send(w)
}
