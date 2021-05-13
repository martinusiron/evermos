package main

import (
	"fmt"
	"net/http"

	dbx "github.com/go-ozzo/ozzo-dbx"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	_ "github.com/lib/pq"
	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/controllers"
	"github.com/martinusiron/evermos/daos"
	"github.com/martinusiron/evermos/errors"
	"github.com/martinusiron/evermos/services"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	logger := logrus.New()

	db, err := dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	http.Handle("/", buildRouter(logger, db))

	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort()
		return c.Write("OK " + app.Version)
	})

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	api := router.Group("/api")

	customerDAO := daos.NewCustomerDAO()
	itemDAO := daos.NewItemDAO()
	orderDAO := daos.NewOrderDAO()

	controllers.ServeCustomerResource(api, services.NewCustomerService(customerDAO, itemDAO, orderDAO))
	controllers.ServeOrderResource(api, services.NewOrderService(orderDAO))
	controllers.ServeItemResource(api, services.NewItemService(itemDAO))

	return router
}
