package vehicles

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/vehicles").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", ctrl.GetAllVehicles).Methods("GET")
	route.HandleFunc("", ctrl.AddVehicle).Methods("POST")
	route.HandleFunc("/popular/", ctrl.GetPopularVehicles).Methods("GET")
	route.HandleFunc("/search/", ctrl.SearchVehicles).Methods("GET")
	route.HandleFunc("/{vehicle_id}", ctrl.UpdateVehicle).Methods("PUT")
	route.HandleFunc("/{vehicle_id}", ctrl.DeleteVehicle).Methods("DELETE")
}
