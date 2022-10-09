package vehicles

import (
	"github.com/chirzul/gorent/src/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/vehicles").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/popular/", ctrl.GetPopularVehicles).Methods("GET")
	route.HandleFunc("", ctrl.GetAllVehicles).Methods("GET")
	route.HandleFunc("", middlewares.Upload(middlewares.CheckAuth(ctrl.AddVehicle, []string{"admin"}))).Methods("POST")
	route.HandleFunc("/search/", ctrl.SearchVehicles).Methods("GET")
	route.HandleFunc("/{vehicle_id}", middlewares.CheckAuth(ctrl.UpdateVehicle, []string{"admin"})).Methods("PUT")
	route.HandleFunc("/{vehicle_id}", middlewares.CheckAuth(ctrl.DeleteVehicle, []string{"admin"})).Methods("DELETE")
}
