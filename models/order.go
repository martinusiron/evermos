package models

import (
	"errors"
	"log"
	"strconv"

	"github.com/mediocregopher/radix.v2/pool"
)

var ErrNoOrder = errors.New("models: no order found")

func init() {
	var err error
	db, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		log.Panic(err)
	}
}

type Purchase_Order struct {
	Purchase_order_id int  `json:"purchase_order_id" db:"pk"`
	Cust_id           int  `json:"cust_id"`
	Item_id           int  `json:"item_id"`
	Quantity          int  `json:"quantity"`
	Is_Pay            bool `json:"is_pay"`
	Dispatched        bool `json:"dispatched"`
}

func PopulateOrder(reply map[string]string) (*Purchase_Order, error) {
	var err error
	ab := new(Purchase_Order)
	ab.Purchase_order_id, err = strconv.Atoi(reply["purchase_order_id"])
	if err != nil {
		return nil, err
	}
	ab.Cust_id, err = strconv.Atoi(reply["cust_id"])
	if err != nil {
		return nil, err
	}
	ab.Item_id, err = strconv.Atoi(reply["item_id"])
	if err != nil {
		return nil, err
	}
	ab.Quantity, err = strconv.Atoi(reply["quantity"])
	if err != nil {
		return nil, err
	}
	ab.Is_Pay, err = strconv.ParseBool(reply["is_pay"])
	ab.Dispatched, err = strconv.ParseBool(reply["dispatched"])
	return ab, nil
}

func FindOrder(id string) (*Purchase_Order, error) {
	reply, err := db.Cmd("HGETALL", "purchase_order_id:"+id).Map()
	if err != nil {
		return nil, err
	} else if len(reply) == 0 {
		return nil, ErrNoOrder
	}

	return PopulateOrder(reply)
}

func CacheOrder(order *Purchase_Order) error {
	resp := db.Cmd("HMSET", "purchase_order_id:"+strconv.Itoa(order.Purchase_order_id), "purchase_order_id", strconv.Itoa(order.Purchase_order_id), "cust_id", strconv.Itoa(order.Cust_id), "item_id", order.Item_id, "quantity", strconv.Itoa(order.Quantity), "is_pay", strconv.FormatBool(order.Is_Pay), "dispatched", strconv.FormatBool(order.Dispatched))
	if resp.Err != nil {
		log.Fatal(resp.Err)
		return resp.Err
	}
	return nil
}
