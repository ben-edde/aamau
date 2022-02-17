package order

import (
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type OrderSerializer struct {
	C *gin.Context
	Order
}

type OrderResponse struct {
	OrderId      uint           `json:"-"`
	OrderDate    datatypes.Date `json:"orderDate"`
	DeliveryDate datatypes.Date `json:"deliveryDate"`
	UserId       uint           `json:"userId"`
	CakeId       uint           `json:"cakeId"`
	Amount       uint           `json:"amount"`
	TotalPrice   float32        `json:"totalPrice"`
}

func (s *OrderSerializer) Response() OrderResponse {
	response := OrderResponse{
		OrderId:      s.OrderId,
		OrderDate:    s.OrderDate,
		DeliveryDate: s.DeliveryDate,
		UserId:       s.UserId,
		CakeId:       s.CakeId,
		Amount:       s.Amount,
		TotalPrice:   s.TotalPrice,
	}
	return response
}
