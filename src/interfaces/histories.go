package interfaces

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/libs"
)

type HistoryRepo interface {
	GetAllHistories() (*models.Histories, error)
	SearchHistory(query map[string]string) (*models.Histories, error)
	AddHistory(data *models.History) (*models.History, error)
	UpdateHistory(data *models.History, id string) (*models.History, error)
	DeleteHistory(data *models.History, id string) (*models.History, error)
	CheckHistory(id string) bool
}

type HistoryService interface {
	GetAllHistories() *libs.Response
	SearchHistory(query map[string]string) *libs.Response
	AddHistory(data *models.History) *libs.Response
	UpdateHistory(data *models.History, id string) *libs.Response
	DeleteHistory(data *models.History, id string) *libs.Response
}
