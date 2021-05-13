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
