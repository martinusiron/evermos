package daos

import (
	"strconv"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/models"
)

type OrderDAO struct{}

func NewOrderDAO() *OrderDAO {
	return &OrderDAO{}
}

func (dao *OrderDAO) Get(rs app.RequestScope, id int) (*models.Purchase_Order, error) {

	cachedOrder, err := models.FindOrder(strconv.Itoa(id))
	if err == models.ErrNoOrder {
		var order models.Purchase_Order
		err := rs.Tx().Select().Model(id, &order)
		models.CacheOrder(&order)
		return &order, err
	} else if err != nil {
		return nil, err
	} else {
		return cachedOrder, err
	}
}

func (dao *OrderDAO) Create(rs app.RequestScope, order *models.Purchase_Order) error {
	order.Purchase_order_id = 0
	return rs.Tx().Model(order).Insert()
}

func (dao *OrderDAO) Update(rs app.RequestScope, id int, order *models.Purchase_Order) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	order.Purchase_order_id = id
	return rs.Tx().Model(order).Exclude("Id").Update()
}

func (dao *OrderDAO) Delete(rs app.RequestScope, id int) error {
	order, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(order).Delete()
}

func (dao *OrderDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("purchase_order").Row(&count)
	return count, err
}

func (dao *OrderDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Purchase_Order, error) {
	order := []models.Purchase_Order{}
	err := rs.Tx().Select().OrderBy("purchase_order_id").Offset(int64(offset)).Limit(int64(limit)).All(&order)
	return order, err
}

func (dao *OrderDAO) GetCustomerCart(rs app.RequestScope, id int) ([]models.Purchase_Order, error) {
	order := []models.Purchase_Order{}
	err := rs.Tx().Select().Where(dbx.HashExp{"cust_id": id, "dispatched": false}).All(&order)
	return order, err
}

func (dao *OrderDAO) GetCustomerCartPrice(rs app.RequestScope, id int) int {
	type OrderPrice struct {
		Item_id    int
		Quantity   int
		Price      int
		Dispatched bool
	}

	var orderPrices []OrderPrice
	err := rs.Tx().NewQuery(
		"select i.item_id, po.quantity, i.price, po.dispatched from purchase_order as po left join item as i on po.item_id = i.item_id where cust_id = 1 and dispatched = false")
	err.All(&orderPrices)

	totalPrice := 0
	for i := 0; i < len(orderPrices); i++ {
		totalPrice += (orderPrices[i].Quantity * orderPrices[i].Price)
	}

	return totalPrice
}

func (dao *OrderDAO) GetCustomerCompletedOrders(rs app.RequestScope, id int) ([]models.Purchase_Order, error) {
	order := []models.Purchase_Order{}
	err := rs.Tx().Select().Where(dbx.HashExp{"cust_id": id, "dispatched": true}).All(&order)
	return order, err
}
