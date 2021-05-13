package testing

import (
	"net/http"
	"testing"

	"github.com/martinusiron/evermos/controllers"
	"github.com/martinusiron/evermos/daos"
	"github.com/martinusiron/evermos/db"
	"github.com/martinusiron/evermos/services"
)

func TestOrder(t *testing.T) {
	db.ResetDB()
	router := newRouter()
	controllers.ServeOrderResource(&router.RouteGroup, services.NewOrderService(daos.NewOrderDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	quantityRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"quantity","error":"cannot be blank"}]}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get an order", "GET", "/orders/1", "", http.StatusOK, `{"purchase_order_id":1, "cust_id":1, "item_id":1,  "quantity":2, "is_pay":false, "dispatched":false}`},
		{"t2 - get a nonexisting order", "GET", "/orders/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an order", "POST", "/orders", `{"cust_id":1, "item_id":1,  "quantity":2, "is_pay":false, "dispatched":false}`, http.StatusOK, `{"purchase_order_id":2, "cust_id":1, "item_id":1,  "quantity":2, "is_pay":false, "dispatched":false}`},
		{"t4 - create an order with validation error", "POST", "/orders", `{"quantity":}`, http.StatusBadRequest, quantityRequiredError},
	})
}
