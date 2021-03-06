package order

import (
	"aamau/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OrderAPIRegister(router *gin.RouterGroup) {
	router.GET("/all", get_all_orders)
	router.GET("/:orderId", get_order_by_id)
	router.POST("/new", new_order)
	router.POST("/", create_order)
	router.DELETE("/:orderId", delete_order_by_id)
	router.PUT("/:orderId", update_order_by_id)
}

func get_order_by_id(c *gin.Context) {
	orderId := c.Param("orderId")
	found_order := Get_order(utils.Get_connection(), orderId)
	orderSerializer := OrderSerializer{c, found_order}
	c.JSON(http.StatusOK, gin.H{"orderId": orderId, "order": orderSerializer.Response()})
}

func get_all_orders(c *gin.Context) {
	response_list := []OrderResponse{}
	found_orders := Get_all_orders()
	for _, order := range found_orders {
		orderSerializer := OrderSerializer{c, order}
		response_list = append(response_list, orderSerializer.Response())
	}
	c.JSON(http.StatusOK, gin.H{"orders": response_list})
}

func create_order(c *gin.Context) {
	orderValidator := OrderValidator{}
	if err := orderValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("%s", err))
		return
	}
	Create_order(utils.Get_connection(), orderValidator.order)
	response := fmt.Sprintf("Created order: %v", orderValidator.order)
	c.JSON(http.StatusOK, gin.H{"content": response})
}

func delete_order_by_id(c *gin.Context) {
	orderId := c.Param("orderId")
	found_order := Get_order(utils.Get_connection(), orderId)
	Delete_order(utils.Get_connection(), fmt.Sprintf("orderId=%s", orderId))
	orderSerializer := OrderSerializer{c, found_order}
	c.JSON(http.StatusOK, gin.H{"orderId": orderId, "deleted_order": orderSerializer.Response()})
}

func update_order_by_id(c *gin.Context) {
	orderId := c.Param("orderId")
	orderValidator := OrderValidator{}
	if err := orderValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("%s", err))
		return
	}
	Update_order(utils.Get_connection(), fmt.Sprintf("orderId=%s", orderId), orderValidator.order)
	orderSerializer := OrderSerializer{c, orderValidator.order}
	c.JSON(http.StatusOK, gin.H{"orderId": orderId, "updated_order": orderSerializer.Response()})
}

func new_order(c *gin.Context) {
	orderValidator := RawOrderValidator{}
	if err := orderValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("%s", err))
		return
	}
	if result_order, accepted := handle_order(utils.Get_connection(), orderValidator); !accepted || result_order == nil {
		c.JSON(http.StatusInternalServerError, "Order rejected.")
		return
	} else {
		Create_order(utils.Get_connection(), *result_order)
		orderSerializer := OrderSerializer{c, *result_order}
		c.JSON(http.StatusOK, gin.H{"response": orderSerializer.Response()})
	}
}
