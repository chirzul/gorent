package users

import (
	"encoding/json"
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/helpers"
	"github.com/chirzul/gorent/src/interfaces"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(s interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: s}
}

func (c *user_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data, err := c.svc.GetAllUsers()
	if err != nil {
		helpers.Response(w, 400, "error", err)
	} else {
		helpers.Response(w, 200, "success get data", err, data)
	}
}

func (c *user_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		_, err := c.svc.AddUser(&datas)
		if err != nil {
			helpers.Response(w, 400, "", err)
		} else {
			helpers.Response(w, 201, "success add data", err)
		}
	}
}

func (c *user_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		_, err := c.svc.UpdateUser(r, &datas)
		if err != nil {
			helpers.Response(w, 400, "", err)
		} else {
			helpers.Response(w, 200, "success update data", err)
		}
	}
}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas models.User
	_, err := c.svc.DeleteUser(r, &datas)
	if err != nil {
		helpers.Response(w, 400, "", err)
	} else {
		helpers.Response(w, 200, "success delete data", err)
	}
}
