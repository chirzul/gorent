package routers

import (
	"errors"

	"github.com/chirzul/gorent/src/databases/orm"
	"github.com/chirzul/gorent/src/modules/v1/histories"
	"github.com/chirzul/gorent/src/modules/v1/users"
	"github.com/chirzul/gorent/src/modules/v1/vehicles"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := orm.New()
	if err != nil {
		return nil, errors.New("gagal init database")
	}

	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	histories.New(mainRoute, db)

	return mainRoute, nil
}
