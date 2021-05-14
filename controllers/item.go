package controllers

import (
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/models"
)

type (
	itemService interface {
		Get(rs app.RequestScope, id int) (*models.Item, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Item, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Item) (*models.Item, error)
		Update(rs app.RequestScope, id int, model *models.Item) (*models.Item, error)
		Delete(rs app.RequestScope, id int) (*models.Item, error)
	}

	itemResource struct {
		service itemService
	}
)

func ServeItemResource(rg *routing.RouteGroup, service itemService) {
	r := &itemResource{service}
	rg.Get("/items/<id>", r.get)
	rg.Get("/items", r.query)
	rg.Post("/items", r.create)
	rg.Put("/items/<id>", r.update)
	rg.Delete("/items/<id>", r.delete)
}

func (r *itemResource) get(c *routing.Context) error {
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

func (r *itemResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *itemResource) create(c *routing.Context) error {
	var model models.Item
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *itemResource) update(c *routing.Context) error {
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

func (r *itemResource) delete(c *routing.Context) error {
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
