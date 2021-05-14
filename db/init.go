package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
	"github.com/martinusiron/evermos/app"
)

var (
	DB *dbx.DB
)

func init() {
	err := app.LoadConfig("./config", "../config")
	if err != nil {
		panic(err)
	}
	DB, err = dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
}

func ResetDB() *dbx.DB {
	if err := runSQLFile(DB, getSQLFile()); err != nil {
		panic(fmt.Errorf("Error while initializing test database: %s", err))
	}
	return DB
}

func getSQLFile() string {
	if _, err := os.Stat("db/db.sql"); err == nil {
		return "db/db.sql"
	}
	return "../db/db.sql"
}

func runSQLFile(db *dbx.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if _, err := db.NewQuery(line).Execute(); err != nil {
			fmt.Println(line)
			return err
		}
	}
	return nil
}
