package order

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

// func Test_crud(t *testing.T) {
// 	utils.CFG_path = "../cfg/config.yaml"
// 	test_order := Order{
// 		OrderId:      42,
// 		OrderDate:    datatypes.Date(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)),
// 		DeliveryDate: datatypes.Date(time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC)),
// 		UserId:       10,
// 		CakeId:       10,
// 		Amount:       3,
// 		TotalPrice:   42.1,
// 	}
// 	fmt.Printf("%v\n", test_order)
// 	// Create_order(test_order)

// 	var found_order Order
// 	found_order = Get_order("42")
// 	fmt.Printf("%v\n", found_order)
// 	if !reflect.DeepEqual(test_order, found_order) {
// 		t.Error("create order failed.")
// 	}

// 	updated_test_order := Order{
// 		OrderDate:    datatypes.Date(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)),
// 		DeliveryDate: datatypes.Date(time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC)),
// 		UserId:       10,
// 		CakeId:       10,
// 		Amount:       5,
// 		TotalPrice:   50,
// 	}
// 	Update_order("orderId=42", updated_test_order)
// 	found_order = Get_order("42")
// 	updated_test_order.OrderId = 42
// 	if !reflect.DeepEqual(updated_test_order, found_order) {
// 		t.Error("update order failed.")
// 	}
// 	Delete_order("orderId=42")
// 	found_order = Get_order("42")
// 	if !reflect.DeepEqual(Order{}, found_order) {
// 		t.Error("delete order failed.")
// 	}
// }

func setup() (*Suite, error) {
	s := &Suite{}
	var (
		db  *sql.DB
		err error
	)

	// matcher_func := sqlmock.QueryMatcherFunc(func(expectedSQL, actualSQL string) error {
	// 	fmt.Printf("expectedSQL:%s\n", expectedSQL)
	// 	fmt.Printf("actualSQL:%s\n", actualSQL)
	// 	fmt.Printf("expectedSQL:%s\n", expectedSQL)
	// 	fmt.Printf("actualSQL:%s\n", actualSQL)
	// 	if expectedSQL != actualSQL {
	// 		return fmt.Errorf("expectedSQL != actualSQL")
	// 	}
	// 	return nil
	// })
	// db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(matcher_func))

	// db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, fmt.Errorf("Failed to open mock sql db, got error: %v\n", err)
	}

	if db == nil {
		return nil, fmt.Errorf("mock db is null\n")
	}

	if s.mock == nil {
		return nil, fmt.Errorf("sqlmock is null\n")
	}

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to open gorm v2 db, got error: %v\n", err)
	}

	if s.db == nil {
		return nil, fmt.Errorf("gorm db is null\n")
	}
	// defer db.Close()
	return s, nil
}
func Test_create_order(t *testing.T) {
	s, err := setup()
	if err != nil {
		t.Errorf("Failed to set up")
	}

	test_order := Order{
		OrderId:      42,
		OrderDate:    datatypes.Date(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)),
		DeliveryDate: datatypes.Date(time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC)),
		UserId:       10,
		CakeId:       10,
		Amount:       3,
		TotalPrice:   42.1,
	}

	tmp1, _ := test_order.OrderDate.Value()
	tmp2, _ := test_order.DeliveryDate.Value()
	args := []driver.Value{tmp1, tmp2, test_order.UserId, test_order.CakeId, test_order.Amount, test_order.TotalPrice, test_order.OrderId}

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO `Orders` (`orderDate`,`deliveryDate`,`userId`,`cakeId`,`amount`,`totalPrice`,`orderId`) VALUES (?,?,?,?,?,?,?)").WithArgs(args...).WillReturnResult(sqlmock.NewResult(42, 1))
	s.mock.ExpectCommit()

	Create_order(s.db, test_order)

	err = s.mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func Test_update_order(t *testing.T) {
	s, err := setup()
	if err != nil {
		t.Errorf("Failed to set up")
	}

	test_order := Order{
		OrderId:      42,
		OrderDate:    datatypes.Date(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)),
		DeliveryDate: datatypes.Date(time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC)),
		UserId:       10,
		CakeId:       10,
		Amount:       30,
		TotalPrice:   42.1,
	}

	tmp1, _ := test_order.OrderDate.Value()
	tmp2, _ := test_order.DeliveryDate.Value()
	args := []driver.Value{test_order.OrderId, tmp1, tmp2, test_order.UserId, test_order.CakeId, test_order.Amount, test_order.TotalPrice}

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("UPDATE `Orders` SET `orderId`=?,`orderDate`=?,`deliveryDate`=?,`userId`=?,`cakeId`=?,`amount`=?,`totalPrice`=? WHERE orderId=42").WithArgs(args...).WillReturnResult(sqlmock.NewResult(42, 1))
	s.mock.ExpectCommit()

	Update_order(s.db, "orderId=42", test_order)

	err = s.mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func Test_get_order(t *testing.T) {
	s, err := setup()
	if err != nil {
		t.Errorf("Failed to set up")
	}

	test_order := Order{
		OrderId:      42,
		OrderDate:    datatypes.Date(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)),
		DeliveryDate: datatypes.Date(time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC)),
		UserId:       10,
		CakeId:       10,
		Amount:       30,
		TotalPrice:   42.1,
	}

	tmp1, _ := test_order.OrderDate.Value()
	tmp2, _ := test_order.DeliveryDate.Value()
	args := []driver.Value{test_order.OrderId, tmp1, tmp2, test_order.UserId, test_order.CakeId, test_order.Amount, test_order.TotalPrice}

	s.mock.ExpectQuery("SELECT * FROM `Orders` WHERE orderId=?").WithArgs(42).WillReturnRows(sqlmock.NewRows([]string{"OrderId", "OrderDate", "DeliveryDate", "UserId", "CakeId", "Amount", "TotalPrice"}).AddRow(args...))

	found_order := Get_order(s.db, "42")
	if found_order != test_order {
		t.Errorf("Error in querying order")
	}
	err = s.mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}
