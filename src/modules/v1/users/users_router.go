package users

import (
	"github.com/chirzul/gorent/src/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middlewares.CheckAuth(ctrl.GetAllUsers, []string{"admin"})).Methods("GET")
	route.HandleFunc("", ctrl.AddUser).Methods("POST")
	route.HandleFunc("/profile", middlewares.CheckAuth(ctrl.FindByUsername, []string{"user", "admin"})).Methods("GET")
	route.HandleFunc("/profile", middlewares.CheckAuth(ctrl.UpdateUser, []string{"user", "admin"})).Methods("PUT")
	route.HandleFunc("/{username}", middlewares.CheckAuth(ctrl.DeleteUser, []string{"user", "admin"})).Methods("DELETE")
}
