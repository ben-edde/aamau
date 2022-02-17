package ingredient

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IngredientAPIRegister(router *gin.RouterGroup) {
	router.GET("/all", get_all_ingredients)
	router.GET("/:ingredientId", get_ingredient_by_id)
	router.POST("/", create_ingredient)
	router.DELETE("/:ingredientId", delete_ingredient_by_id)
	router.PUT("/:ingredientId", update_ingredient_by_id)
}

func get_ingredient_by_id(c *gin.Context) {
	ingredientId := c.Param("ingredientId")
	found_ingredient := Get_ingredient(ingredientId)
	ingredientSerializer := IngredientSerializer{c, found_ingredient}
	c.JSON(http.StatusOK, gin.H{"ingredientId": ingredientId, "ingredient": ingredientSerializer.Response()})
}

func get_all_ingredients(c *gin.Context) {
	response_list := []IngredientResponse{}
	found_ingredients := Get_all_ingredients()
	for _, ingredient := range found_ingredients {
		ingredientSerializer := IngredientSerializer{c, ingredient}
		response_list = append(response_list, ingredientSerializer.Response())
	}
	c.JSON(http.StatusOK, gin.H{"ingredients": response_list})
}

func create_ingredient(c *gin.Context) {
	ingredientValidator := IngredientValidator{}
	if err := ingredientValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("%s", err))
		return
	}
	Create_ingredient(ingredientValidator.ingredient)
	response := fmt.Sprintf("Created ingredient: %v", ingredientValidator.ingredient)
	c.JSON(http.StatusOK, gin.H{"content": response})
}

func delete_ingredient_by_id(c *gin.Context) {
	ingredientId := c.Param("ingredientId")
	found_ingredient := Get_ingredient(ingredientId)
	Delete_ingredient(fmt.Sprintf("ingredientId=%s", ingredientId))
	ingredientSerializer := IngredientSerializer{c, found_ingredient}
	c.JSON(http.StatusOK, gin.H{"ingredientId": ingredientId, "deleted_ingredient": ingredientSerializer.Response()})
}

func update_ingredient_by_id(c *gin.Context) {
	ingredientId := c.Param("ingredientId")
	ingredientValidator := IngredientValidator{}
	if err := ingredientValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("%s", err))
		return
	}
	Update_ingredient(fmt.Sprintf("ingredientId=%s", ingredientId), ingredientValidator.ingredient)
	ingredientSerializer := IngredientSerializer{c, ingredientValidator.ingredient}
	c.JSON(http.StatusOK, gin.H{"ingredientId": ingredientId, "updated_ingredient": ingredientSerializer.Response()})
}
