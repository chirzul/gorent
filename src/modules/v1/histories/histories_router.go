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

	route.HandleFunc("", middlewares.CheckAuth(ctrl.GetAllHistories, []string{"user", "admin"})).Methods("GET")
	route.HandleFunc("", middlewares.CheckAuth(ctrl.AddHistory, []string{"admin"})).Methods("POST")
	route.HandleFunc("/search/", middlewares.CheckAuth(ctrl.SearchHistory, []string{"user", "admin"})).Methods("GET")
	route.HandleFunc("/{history_id}", middlewares.CheckAuth(ctrl.UpdateHistory, []string{"admin"})).Methods("PUT")
	route.HandleFunc("/{history_id}", middlewares.CheckAuth(ctrl.DeleteHistory, []string{"admin"})).Methods("DELETE")
}
