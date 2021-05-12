package main

import (
	"fmt"
	"net/http"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/errors"
	"github.com/sirupsen/logrus"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "evermos"
)

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	// connect to the database
	db, err := dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}
