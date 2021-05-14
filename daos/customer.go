package daos

import (
	"strconv"

	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/models"
)

type CustomerDAO struct{}

func NewCustomerDAO() *CustomerDAO {
	return &CustomerDAO{}
}

func (dao *CustomerDAO) Get(rs app.RequestScope, id int) (*models.Customer, error) {
	cachedCustomer, err := models.FindCustomer(strconv.Itoa(id))
	if err == models.ErrorNoCustomer {
		var customer models.Customer
		err := rs.Tx().Select().Model(id, &customer)
		models.CacheCustomer(&customer)
		return &customer, err
	} else if err != nil {
		return nil, err
	} else {
		return cachedCustomer, err
	}
}

func (dao *CustomerDAO) Create(rs app.RequestScope, customer *models.Customer) error {
	customer.Cust_id = 0
	return rs.Tx().Model(customer).Insert()
}

func (dao *CustomerDAO) Update(rs app.RequestScope, id int, customer *models.Customer) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	customer.Cust_id = id
	return rs.Tx().Model(customer).Exclude("Id").Update()
}

func (dao *CustomerDAO) Delete(rs app.RequestScope, id int) error {
	customer, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(customer).Delete()
}

func (dao *CustomerDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("customer").Row(&count)
	return count, err
}

func (dao *CustomerDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Customer, error) {
	customer := []models.Customer{}
	err := rs.Tx().Select().OrderBy("cust_id").Offset(int64(offset)).Limit(int64(limit)).All(&customer)
	return customer, err
}
