package order

import (
	"aamau/utils"
	"fmt"
	"strconv"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Order struct {
	OrderId      uint           `gorm:"type:uint auto_increment; primary_key;column:orderId" form:"orderId" json:"orderId"`
	OrderDate    datatypes.Date `gorm:"type:date NOT NULL;column:orderDate" form:"orderDate" json:"orderDate"`
	DeliveryDate datatypes.Date `gorm:"type:date NOT NULL;column:deliveryDate" form:"deliveryDate" json:"deliveryDate"`
	UserId       uint           `gorm:"type:uint NOT NULL;column:userId" form:"userId" json:"userId"`
	CakeId       uint           `gorm:"type:uint NOT NULL;column:cakeId" form:"cakeId" json:"cakeId"`
	Amount       uint           `gorm:"type:uint NOT NULL;column:amount" form:"amount" json:"amount"`
	TotalPrice   float32        `gorm:"type:float NOT NULL;column:totalPrice" form:"totalPrice" json:"totalPrice"`
}

func Create_order(conn *gorm.DB, order Order) {
	result := conn.Debug().Table("Orders").Model(&Order{}).Create(&order)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Printf("create order failed: %s\n", result.Error)
	}
}

func Get_order(orderId string) Order {
	uid, err := strconv.Atoi(orderId)
	if err != nil {
		fmt.Println(err)
		return Order{}
	}
	conn := utils.Get_connection()
	var order Order
	result := conn.Debug().Table("Orders").Where("orderId=?", uid).Find(&order)
	if result.Error != nil || result.RowsAffected > 1 {
		fmt.Printf("get order failed: %s\n", result.Error)
	}
	return order
}
func Get_all_orders() []Order {
	conn := utils.Get_connection()
	var orderList []Order
	conn.Debug().Table("Orders").Model(&Order{}).Find(&orderList)
	return orderList
}
func Update_order(conn *gorm.DB, up_condition string, order Order) {
	result := conn.Debug().Table("Orders").Model(&Order{}).Where(up_condition).Updates(order)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Printf("update order failed: %s\n", result.Error)
	}
}
func Delete_order(rm_condition string) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Orders").Model(&Order{}).Where(rm_condition).Delete(&Order{})
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("delete order failed.")
	}
}
