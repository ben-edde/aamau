package order

import (
	"fmt"
	"strconv"
	"time"

	"aamau/user"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type RawOrderInfo struct {
	OrderDate       string `form:"orderDate" json:"orderDate"`
	CakeId          string `form:"cakeId" json:"cakeId"`
	Amount          string `form:"amount" json:"amount"`
	UserName        string `form:"userName" json:"userName"`
	ContactNo       string `form:"contactNo" json:"contactNo"`
	Email           string `form:"email" json:"email"`
	DeliveryAddress string `form:"deliveryAddress" json:"deliveryAddress"`
}

type OrderInfo struct {
	OrderDate    string `form:"orderDate" json:"orderDate"`
	DeliveryDate string `form:"deliveryDate" json:"deliveryDate"`
	UserId       string `form:"userId" json:"userId"`
	CakeId       string `form:"cakeId" json:"cakeId"`
	Amount       string `form:"amount" json:"amount"`
	TotalPrice   string `form:"totalPrice" json:"totalPrice"`
}

type OrderValidator struct {
	orderInfo OrderInfo
	order     Order
}
type RawOrderValidator struct {
	orderInfo RawOrderInfo
	order     Order
	user      user.User
}

func (self *RawOrderValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(&(self.orderInfo))

	fmt.Printf("OrderInfo: %v\n", self.orderInfo)
	if err != nil {
		return err
	}
	if orderDate, err := strconv.ParseInt(self.orderInfo.OrderDate, 10, 64); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.order.OrderDate = datatypes.Date(time.Unix(int64(orderDate), 0))
	}
	if cakeId, err := strconv.Atoi(self.orderInfo.CakeId); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.order.CakeId = uint(cakeId)
	}
	if amount, err := strconv.Atoi(self.orderInfo.Amount); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.order.Amount = uint(amount)
	}
	self.user.UserName = self.orderInfo.UserName
	self.user.Email = self.orderInfo.Email
	self.user.ContactNo = self.orderInfo.ContactNo
	self.user.DeliveryAddress = self.orderInfo.DeliveryAddress

	return nil
}

func (self *OrderValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(&(self.orderInfo))
	if err != nil {
		return err
	}
	if orderDate, err := strconv.ParseInt(self.orderInfo.OrderDate, 10, 64); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.order.OrderDate = datatypes.Date(time.Unix(int64(orderDate), 0))
	}

	if deliveryDate, err := strconv.ParseInt(self.orderInfo.DeliveryDate, 10, 64); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.order.DeliveryDate = datatypes.Date(time.Unix(int64(deliveryDate), 0))
	}
	if userId, err := strconv.Atoi(self.orderInfo.UserId); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.order.UserId = uint(userId)
	}
	if cakeId, err := strconv.Atoi(self.orderInfo.CakeId); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.order.CakeId = uint(cakeId)
	}

	if amount, err := strconv.Atoi(self.orderInfo.Amount); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.order.Amount = uint(amount)
	}
	if totalPrice, err := strconv.ParseFloat(self.orderInfo.TotalPrice, 32); err != nil {
		return err
	} else {
		self.order.TotalPrice = float32(totalPrice)
	}
	return nil
}
