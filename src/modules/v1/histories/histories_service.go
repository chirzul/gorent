package histories

import (
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
)

type history_service struct {
	repo interfaces.HistoryRepo
}

func NewService(repo interfaces.HistoryRepo) *history_service {
	return &history_service{repo}
}

func (s *history_service) GetAllHistories() (*models.Histories, error) {
	data, err := s.repo.FindAllHistories()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *history_service) SearchHistory(r *http.Request) (*models.Histories, error) {
	data, err := s.repo.FindHistory(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *history_service) AddHistory(data *models.History) (*models.History, error) {
	data, err := s.repo.SaveHistory(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *history_service) UpdateHistory(r *http.Request, data *models.History) (*models.History, error) {
	data, err := s.repo.ChangeHistory(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *history_service) DeleteHistory(r *http.Request, data *models.History) (*models.History, error) {
	data, err := s.repo.RemoveHistory(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
