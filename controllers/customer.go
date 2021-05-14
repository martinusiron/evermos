package controllers

import (
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/models"
)

type (
	// customerService specifies the interface for the customer service needed by customerResource.
	customerService interface {
		Get(rs app.RequestScope, id int) (*models.Customer, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Customer, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Customer) (*models.Customer, error)
		Update(rs app.RequestScope, id int, model *models.Customer) (*models.Customer, error)
		Delete(rs app.RequestScope, id int) (*models.Customer, error)

		GetCart(rs app.RequestScope, id int) ([]models.Purchase_Order, error)
		GetCartPrice(rs app.RequestScope, id int) int
		GetOrderTransactions(rs app.RequestScope, id int) ([]models.Purchase_Order, error)
	}

	customerResource struct {
		service customerService
	}
)

func ServeCustomerResource(rg *routing.RouteGroup, service customerService) {
	r := &customerResource{service}
	rg.Get("/customers/<id>", r.get)
	rg.Get("/customers", r.query)
	rg.Post("/customers", r.create)
	rg.Put("/customers/<id>", r.update)
	rg.Delete("/customers/<id>", r.delete)

	rg.Get("/customers/<id>/cart", r.getCart)
	rg.Get("/customers/<id>/cart/price", r.getCartPrice)
	rg.Get("/customers/<id>/transactions", r.getOrderTransactions)
}

func (r *customerResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *customerResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	customers, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = customers
	return c.Write(paginatedList)
}

func (r *customerResource) create(c *routing.Context) error {
	var model models.Customer
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *customerResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *customerResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *customerResource) getCart(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	response, err := r.service.GetCart(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *customerResource) getCartPrice(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	response := r.service.GetCartPrice(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *customerResource) getOrderTransactions(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	response, err := r.service.GetOrderTransactions(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
