package models

import (
	"errors"
	"log"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mediocregopher/radix.v2/pool"
)

var ErrorNoCustomer = errors.New("models: no customers found")

func init() {
	var err error
	db, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		log.Panic(err)
	}
}

type Customer struct {
	Cust_id    int    `json:"cust_id" db:"pk"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Address    string `json:"address"`
}

func (m Customer) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.First_name, validation.Required, validation.Length(0, 50)),
	)
}

func PopulateCustomer(reply map[string]string) (*Customer, error) {
	var err error
	ab := new(Customer)
	ab.Cust_id, err = strconv.Atoi(reply["cust_id"])
	if err != nil {
		return nil, err
	}
	ab.First_name = reply["first_name"]
	ab.Last_name = reply["last_name"]
	ab.Address = reply["address"]
	return ab, nil
}

func FindCustomer(id string) (*Customer, error) {
	reply, err := db.Cmd("HGETALL", "cust_id:"+id).Map()
	if err != nil {
		return nil, err
	} else if len(reply) == 0 {
		return nil, ErrorNoCustomer
	}

	return PopulateCustomer(reply)
}

func CacheCustomer(customer *Customer) error {
	resp := db.Cmd("HMSET", "cust_id:"+strconv.Itoa(customer.Cust_id), "cust_id", strconv.Itoa(customer.Cust_id), "first_name", customer.First_name, "last_name", customer.Last_name, "address", customer.Address)
	if resp.Err != nil {
		log.Fatal(resp.Err)
		return resp.Err
	}
	return nil
}
