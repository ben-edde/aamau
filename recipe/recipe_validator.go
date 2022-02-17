package recipe

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RecipeInfo struct {
	CakeId         string `form:"cakeId" json:"cakeId" binding:"max=255"`
	IngredientId   string `form:"ingredientId" json:"ingredientId" binding:"max=255"`
	IntAmtRequired string `form:"ingredientAmountRequired" json:"ingredientAmountRequired" binding:"max=255"`
}

type RecipeValidator struct {
	recipeInfo RecipeInfo
	recipe     Recipe
}

func (self *RecipeValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(&(self.recipeInfo))
	if err != nil {
		return err
	}
	fmt.Printf("self.recipeInfo: %v\n", self.recipeInfo)
	if cakeId, err := strconv.Atoi(self.recipeInfo.CakeId); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.recipe.CakeId = uint(cakeId)
	}

	fmt.Printf("self.recipeInfo: %v\n", self.recipeInfo)
	if ingredientId, err := strconv.Atoi(self.recipeInfo.IngredientId); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.recipe.IngredientId = uint(ingredientId)
	}

	if intAmtRequired, err := strconv.Atoi(self.recipeInfo.IntAmtRequired); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.recipe.IntAmtRequired = uint(intAmtRequired)
	}

	fmt.Printf("self.recipeInfo: %v\n", self.recipeInfo)
	return nil
}
