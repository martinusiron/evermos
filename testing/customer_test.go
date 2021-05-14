package testing

import (
	"net/http"
	"testing"

	"github.com/martinusiron/evermos/controllers"
	"github.com/martinusiron/evermos/daos"
	"github.com/martinusiron/evermos/db"
	"github.com/martinusiron/evermos/services"
)

func TestCustomer(t *testing.T) {
	db.ResetDB()
	router := newRouter()
	controllers.ServeCustomerResource(&router.RouteGroup, services.NewCustomerService(daos.NewCustomerDAO(), daos.NewItemDAO(), daos.NewOrderDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	RequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"first_name","error":"cannot be blank"},{"field":"last_name","error":"cannot be blank"},{"field":"address","error":"cannot be blank"}]}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get an customer", "GET", "/customers/1", "", http.StatusOK, `{"cust_id":1,  "first_name":"John", "last_name":"Doe", "address":"Jakarta"}`},
		{"t2 - get a nonexisting customer", "GET", "/customers/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an customer", "POST", "/customers", `{"first_name":"Mr", "last_name":"Joel", "address":"Jakarta"}`, http.StatusOK, `{"cust_id": 3, "first_name":"Mr", "last_name":"Joel", "address":"Jakarta"}`},
		{"t4 - create an customer with validation error", "POST", "/customers", `{"first_name":"","last_name":"","address":""`, http.StatusBadRequest, RequiredError},
		{"t5 - update an customer", "PUT", "/customers/3", `{"last_name": "Dodi"}`, http.StatusOK, `{"cust_id": 3, "first_name":"Mr", "last_name":"Joel", "address":"Jakarta"}`},
		{"t6 - update an customer with validation error", "PUT", "/customers/2", `{"first_name":""}`, http.StatusBadRequest, RequiredError},
		{"t7 - update a nonexisting customer", "PUT", "/customers/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an customer", "DELETE", "/customers/3", ``, http.StatusOK, `{"cust_id": 3, "first_name":"Mr", "last_name":"Joel", "address":"Jakarta"}`},
		{"t9 - delete a nonexisting customer", "DELETE", "/customers/99999", "", http.StatusNotFound, notFoundError},
	})
}
