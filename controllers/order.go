package controllers

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/models"
)

type (
	// orderService specifies the interface for the artist service needed by orderResource.
	orderService interface {
		Get(rs app.RequestScope, id int) (*models.Purchase_Order, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Purchase_Order, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Purchase_Order) (*models.Purchase_Order, error)
		Update(rs app.RequestScope, id int, model *models.Purchase_Order) (*models.Purchase_Order, error)
		Delete(rs app.RequestScope, id int) (*models.Purchase_Order, error)
	}

	// orderResource defines the handlers for the CRUD APIs.
	orderResource struct {
		service orderService
	}
)

func ServeOrderResource(rg *routing.RouteGroup, service orderService) {
	r := &orderResource{service}
	rg.Get("/orders/<id>", r.get)
	rg.Get("/orders", r.query)
	rg.Post("/orders", r.create)
	rg.Put("/orders/<id>", r.update)
	rg.Delete("/orders/<id>", r.delete)
}

func (r *orderResource) get(c *routing.Context) error {
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

func (r *orderResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	orders, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = orders
	return c.Write(paginatedList)
}

func (r *orderResource) create(c *routing.Context) error {
	var model models.Purchase_Order
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *orderResource) update(c *routing.Context) error {
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

func (r *orderResource) delete(c *routing.Context) error {
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
