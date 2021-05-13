package services

import (
	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/models"
)

type orderDAO interface {
	Get(rs app.RequestScope, id int) (*models.Purchase_Order, error)
	Count(rs app.RequestScope) (int, error)
	Query(rs app.RequestScope, offset, limit int) ([]models.Purchase_Order, error)
	Create(rs app.RequestScope, order *models.Purchase_Order) error
	Update(rs app.RequestScope, id int, order *models.Purchase_Order) error
	Delete(rs app.RequestScope, id int) error
	GetCustomerCart(rs app.RequestScope, id int) ([]models.Purchase_Order, error)
	GetCustomerCartPrice(rs app.RequestScope, id int) int
	GetCustomerCompletedOrders(rs app.RequestScope, id int) ([]models.Purchase_Order, error)
}
type OrderService struct {
	dao orderDAO
}

func NewOrderService(dao orderDAO) *OrderService {
	return &OrderService{dao}
}

func (s *OrderService) Get(rs app.RequestScope, id int) (*models.Purchase_Order, error) {
	return s.dao.Get(rs, id)
}

func (s *OrderService) Create(rs app.RequestScope, model *models.Purchase_Order) (*models.Purchase_Order, error) {
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Purchase_order_id)
}

func (s *OrderService) Update(rs app.RequestScope, id int, model *models.Purchase_Order) (*models.Purchase_Order, error) {
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

func (s *OrderService) Delete(rs app.RequestScope, id int) (*models.Purchase_Order, error) {
	order, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return order, err
}

func (s *OrderService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

func (s *OrderService) Query(rs app.RequestScope, offset, limit int) ([]models.Purchase_Order, error) {
	return s.dao.Query(rs, offset, limit)
}

func (s *OrderService) GetCustomerCart(rs app.RequestScope, id int) ([]models.Purchase_Order, error) {
	return s.dao.GetCustomerCart(rs, id)
}

func (s *OrderService) GetCustomerCartPrice(rs app.RequestScope, id int) int {
	return s.dao.GetCustomerCartPrice(rs, id)
}

func (s *OrderService) GetCustomerCompletedOrders(rs app.RequestScope, id int) ([]models.Purchase_Order, error) {
	return s.dao.GetCustomerCompletedOrders(rs, id)
}
