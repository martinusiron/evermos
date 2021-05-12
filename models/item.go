package models

import (
	"errors"
	"log"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mediocregopher/radix.v2/pool"
)

// Declare a global db variable to store the Redis connection pool.
var db *pool.Pool
var ErrNoItem = errors.New("models: no items found")

func init() {
	var err error
	// Establish a pool of 10 connections to the Redis server listening on
	// port 6379 of the local machine.
	db, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		log.Panic(err)
	}
}

type Item struct {
	Item_id int    `json:"item_id" db:"pk"`
	Name    string `json:"name"`
	Stock   int    `json:"stock"`
	Price   int    `json:"price"`
}

// Validate validates the Item fields.
func (m Item) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 50)),
	)
}

func PopulateItem(reply map[string]string) (*Item, error) {
	var err error
	newItem := new(Item)
	newItem.Item_id, err = strconv.Atoi(reply["item_id"])
	if err != nil {
		return nil, err
	}

	newItem.Name = reply["name"]
	newItem.Stock, err = strconv.Atoi(reply["stock"])
	if err != nil {
		return nil, err
	}
	newItem.Price, err = strconv.Atoi(reply["price"])
	if err != nil {
		return nil, err
	}
	return newItem, nil
}

func FindItem(id string) (*Item, error) {
	reply, err := db.Cmd("HGETALL", "item_id:"+id).Map()
	if err != nil {
		return nil, err
	} else if len(reply) == 0 {
		return nil, ErrNoItem
	}

	return PopulateItem(reply)
}

func CacheItem(item *Item) error {
	resp := db.Cmd("HMSET", "item_id:"+strconv.Itoa(item.Item_id), "item_id", strconv.Itoa(item.Item_id), "name", item.Name, "stock", strconv.Itoa(item.Stock), "price", strconv.Itoa(item.Price))
	if resp.Err != nil {
		log.Fatal(resp.Err)
		return resp.Err
	}
	return nil
}
