package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAllUsers).Methods("GET")
	route.HandleFunc("/", ctrl.AddUser).Methods("POST")
	route.HandleFunc("/{username}", ctrl.UpdateUser).Methods("PUT")
	route.HandleFunc("/{username}", ctrl.DeleteUser).Methods("DELETE")
}
