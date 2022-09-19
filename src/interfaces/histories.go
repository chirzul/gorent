package interfaces

import (
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
)

type HistoryRepo interface {
	FindAllHistories() (*models.Histories, error)
	FindHistoriesByVehicle(r *http.Request) (*models.Histories, error)
	SaveHistory(data *models.History) (*models.History, error)
	ChangeHistory(r *http.Request, data *models.History) (*models.History, error)
	RemoveHistory(r *http.Request, data *models.History) (*models.History, error)
}

type HistoryService interface {
	GetAllHistories() (*models.Histories, error)
	SearchHistoriesByVehicle(r *http.Request) (*models.Histories, error)
	AddHistory(data *models.History) (*models.History, error)
	UpdateHistory(r *http.Request, data *models.History) (*models.History, error)
	DeleteHistory(r *http.Request, data *models.History) (*models.History, error)
}
