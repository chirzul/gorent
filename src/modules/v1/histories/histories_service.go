package histories

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
	"github.com/chirzul/gorent/src/libs"
)

type history_service struct {
	repo interfaces.HistoryRepo
}

func NewService(repo interfaces.HistoryRepo) *history_service {
	return &history_service{repo}
}

func (s *history_service) GetAllHistories() *libs.Response {
	data, err := s.repo.GetAllHistories()
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *history_service) SearchHistory(query map[string]string) *libs.Response {
	data, err := s.repo.SearchHistory(query)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *history_service) AddHistory(data *models.History) *libs.Response {
	data, err := s.repo.AddHistory(data)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *history_service) UpdateHistory(data *models.History, id string) *libs.Response {
	if checkHistory := s.repo.CheckHistory(id); !checkHistory {
		return libs.GetResponse("history not found", 404, true)
	}
	data, err := s.repo.UpdateHistory(data, id)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *history_service) DeleteHistory(data *models.History, id string) *libs.Response {
	if checkHistory := s.repo.CheckHistory(id); !checkHistory {
		return libs.GetResponse("history not found", 404, true)
	}
	data, err := s.repo.DeleteHistory(data, id)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}
