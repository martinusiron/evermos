package services

import (
	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/models"
)

type customerDAO interface {
	Get(rs app.RequestScope, id int) (*models.Customer, error)
	Create(rs app.RequestScope, customer *models.Customer) error
	Update(rs app.RequestScope, id int, customer *models.Customer) error
	Delete(rs app.RequestScope, id int) error
	Count(rs app.RequestScope) (int, error)
	Query(rs app.RequestScope, offset, limit int) ([]models.Customer, error)
}

type CustomerService struct {
	cust_dao  customerDAO
	item_dao  itemDAO
	order_dao orderDAO
}

func NewCustomerService(cust_dao customerDAO, item_dao itemDAO, order_dao orderDAO) *CustomerService {
	return &CustomerService{cust_dao, item_dao, order_dao}
}

func (s *CustomerService) Get(rs app.RequestScope, id int) (*models.Customer, error) {
	return s.cust_dao.Get(rs, id)
}

func (s *CustomerService) Create(rs app.RequestScope, model *models.Customer) (*models.Customer, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.cust_dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.cust_dao.Get(rs, model.Cust_id)
}

func (s *CustomerService) Update(rs app.RequestScope, id int, model *models.Customer) (*models.Customer, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.cust_dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.cust_dao.Get(rs, id)
}

func (s *CustomerService) Delete(rs app.RequestScope, id int) (*models.Customer, error) {
	customer, err := s.cust_dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.cust_dao.Delete(rs, id)
	return customer, err
}

func (s *CustomerService) Count(rs app.RequestScope) (int, error) {
	return s.cust_dao.Count(rs)
}

func (s *CustomerService) Query(rs app.RequestScope, offset, limit int) ([]models.Customer, error) {
	return s.cust_dao.Query(rs, offset, limit)
}

func (s *CustomerService) GetCart(rs app.RequestScope, id int) ([]models.Purchase_Order, error) {
	orders, err := s.order_dao.GetCustomerCart(rs, id)
	if err != nil {
		return nil, err
	}

	return orders, err
}

func (s *CustomerService) GetCartPrice(rs app.RequestScope, id int) int {
	cartPrice := s.order_dao.GetCustomerCartPrice(rs, id)
	return cartPrice
}

func (s *CustomerService) GetOrderTransactions(rs app.RequestScope, id int) ([]models.Purchase_Order, error) {
	orders, err := s.order_dao.GetCustomerCompletedOrders(rs, id)
	if err != nil {
		return nil, err
	}

	return orders, err
}
