package services

import (
	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/models"
)

type itemDAO interface {
	Get(rs app.RequestScope, id int) (*models.Item, error)
	Create(rs app.RequestScope, item *models.Item) error
	Update(rs app.RequestScope, id int, item *models.Item) error
	Delete(rs app.RequestScope, id int) error
	Count(rs app.RequestScope) (int, error)
	Query(rs app.RequestScope, offset, limit int) ([]models.Item, error)
}

type ItemService struct {
	dao itemDAO
}

func NewItemService(dao itemDAO) *ItemService {
	return &ItemService{dao}
}

func (s *ItemService) Get(rs app.RequestScope, id int) (*models.Item, error) {
	return s.dao.Get(rs, id)
}

func (s *ItemService) Create(rs app.RequestScope, model *models.Item) (*models.Item, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Item_id)
}

func (s *ItemService) Update(rs app.RequestScope, id int, model *models.Item) (*models.Item, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

func (s *ItemService) Delete(rs app.RequestScope, id int) (*models.Item, error) {
	item, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return item, err
}

func (s *ItemService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

func (s *ItemService) Query(rs app.RequestScope, offset, limit int) ([]models.Item, error) {
	return s.dao.Query(rs, offset, limit)
}
