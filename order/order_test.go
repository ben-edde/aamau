package order

import (
	"aamau/utils"
	"fmt"
	"reflect"
	"testing"
	"time"

	"gorm.io/datatypes"
)

func Test_crud(t *testing.T) {
	utils.CFG_path = "../cfg/config.yaml"
	test_order := Order{
		OrderId:      42,
		OrderDate:    datatypes.Date(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)),
		DeliveryDate: datatypes.Date(time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC)),
		UserId:       10,
		CakeId:       10,
		Amount:       3,
		TotalPrice:   42.1,
	}
	fmt.Printf("%v\n", test_order)
	Create_order(test_order)

	var found_order Order
	found_order = Get_order("42")
	fmt.Printf("%v\n", found_order)
	if !reflect.DeepEqual(test_order, found_order) {
		t.Error("create order failed.")
	}

	updated_test_order := Order{
		OrderDate:    datatypes.Date(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)),
		DeliveryDate: datatypes.Date(time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC)),
		UserId:       10,
		CakeId:       10,
		Amount:       5,
		TotalPrice:   50,
	}
	Update_order("orderId=42", updated_test_order)
	found_order = Get_order("42")
	updated_test_order.OrderId = 42
	if !reflect.DeepEqual(updated_test_order, found_order) {
		t.Error("update order failed.")
	}
	Delete_order("orderId=42")
	found_order = Get_order("42")
	if !reflect.DeepEqual(Order{}, found_order) {
		t.Error("delete order failed.")
	}
}
