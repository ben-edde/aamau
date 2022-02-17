package recipe

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecipeAPIRegister(router *gin.RouterGroup) {
	router.GET("/all", get_all_recipes)
	router.GET("/:recipeId", get_recipe_by_id)
	router.POST("/", create_recipe)
	router.DELETE("/:recipeId", delete_recipe_by_id)
	router.PUT("/:recipeId", update_recipe_by_id)
}

func get_recipe_by_id(c *gin.Context) {
	recipeId := c.Param("recipeId")
	found_recipe := Get_recipe(recipeId)
	recipeSerializer := RecipeSerializer{c, found_recipe}
	c.JSON(http.StatusOK, gin.H{"recipeId": recipeId, "recipe": recipeSerializer.Response()})
}

func get_all_recipes(c *gin.Context) {
	response_list := []RecipeResponse{}
	found_recipes := Get_all_recipes()
	for _, recipe := range found_recipes {
		recipeSerializer := RecipeSerializer{c, recipe}
		response_list = append(response_list, recipeSerializer.Response())
	}
	c.JSON(http.StatusOK, gin.H{"recipes": response_list})
}

func create_recipe(c *gin.Context) {
	recipeValidator := RecipeValidator{}
	if err := recipeValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("%s", err))
		return
	}
	Create_recipe(recipeValidator.recipe)
	response := fmt.Sprintf("Created recipe: %v", recipeValidator.recipe)
	c.JSON(http.StatusOK, gin.H{"content": response})
}

func delete_recipe_by_id(c *gin.Context) {
	recipeId := c.Param("recipeId")
	found_recipe := Get_recipe(recipeId)
	Delete_recipe(fmt.Sprintf("recipeId=%s", recipeId))
	recipeSerializer := RecipeSerializer{c, found_recipe}
	c.JSON(http.StatusOK, gin.H{"recipeId": recipeId, "deleted_recipe": recipeSerializer.Response()})
}

func update_recipe_by_id(c *gin.Context) {
	recipeId := c.Param("recipeId")
	recipeValidator := RecipeValidator{}
	if err := recipeValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("%s", err))
		return
	}
	Update_recipe(fmt.Sprintf("recipeId=%s", recipeId), recipeValidator.recipe)
	recipeSerializer := RecipeSerializer{c, recipeValidator.recipe}
	c.JSON(http.StatusOK, gin.H{"recipeId": recipeId, "updated_recipe": recipeSerializer.Response()})
}
