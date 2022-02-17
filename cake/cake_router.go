package cake

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CakeAPIRegister(router *gin.RouterGroup) {
	router.GET("/all", get_all_cakes)
	router.GET("/:cakeId", get_cake_by_id)
	router.POST("/", create_cake)
	router.DELETE("/:cakeId", delete_cake_by_id)
	router.PUT("/:cakeId", update_cake_by_id)
}

func get_cake_by_id(c *gin.Context) {
	cakeId := c.Param("cakeId")
	found_cake := Get_cake(cakeId)
	cakeSerializer := CakeSerializer{c, found_cake}
	c.JSON(http.StatusOK, gin.H{"cakeId": cakeId, "cake": cakeSerializer.Response()})
}

func get_all_cakes(c *gin.Context) {
	response_list := []CakeResponse{}
	found_cakes := Get_all_cakes()
	for _, cake := range found_cakes {
		cakeSerializer := CakeSerializer{c, cake}
		response_list = append(response_list, cakeSerializer.Response())
	}
	c.JSON(http.StatusOK, gin.H{"cakes": response_list})
}

func create_cake(c *gin.Context) {
	cakeValidator := CakeValidator{}
	if err := cakeValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("%s", err))
		return
	}

	Create_cake(cakeValidator.cake)
	response := fmt.Sprintf("Created cake: %v", cakeValidator.cake)
	c.JSON(http.StatusOK, gin.H{"content": response})
}

func delete_cake_by_id(c *gin.Context) {
	cakeId := c.Param("cakeId")
	found_cake := Get_cake(cakeId)
	Delete_cake(fmt.Sprintf("cakeId=%s", cakeId))
	cakeSerializer := CakeSerializer{c, found_cake}
	c.JSON(http.StatusOK, gin.H{"cakeId": cakeId, "deleted_cake": cakeSerializer.Response()})
}

func update_cake_by_id(c *gin.Context) {
	cakeId := c.Param("cakeId")
	cakeValidator := CakeValidator{}
	if err := cakeValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("%s", err))
		return
	}
	Update_cake(fmt.Sprintf("cakeId=%s", cakeId), cakeValidator.cake)
	cakeSerializer := CakeSerializer{c, cakeValidator.cake}
	c.JSON(http.StatusOK, gin.H{"cakeId": cakeId, "updated_cake": cakeSerializer.Response()})
}
