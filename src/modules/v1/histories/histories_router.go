package histories

import (
	"github.com/chirzul/gorent/src/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/histories").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middlewares.CheckAuth(ctrl.GetAllHistories)).Methods("GET")
	route.HandleFunc("", middlewares.CheckAuth(ctrl.AddHistory)).Methods("POST")
	route.HandleFunc("/search/", middlewares.CheckAuth(ctrl.SearchHistory)).Methods("GET")
	route.HandleFunc("/{history_id}", middlewares.CheckAuth(ctrl.UpdateHistory)).Methods("PUT")
	route.HandleFunc("/{history_id}", middlewares.CheckAuth(ctrl.DeleteHistory)).Methods("DELETE")
}
