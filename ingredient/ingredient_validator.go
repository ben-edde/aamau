package ingredient

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IngredientInfo struct {
	IngredientName   string `form:"ingredientName" json:"ingredientName" binding:"max=255"`
	IngredientWeight string `form:"ingredientWeight" json:"ingredientWeight" binding:"max=255"`
	IngredientAmount string `form:"ingredientAmount" json:"ingredientAmount" binding:"max=255"`
}

type IngredientValidator struct {
	ingredientInfo IngredientInfo
	ingredient     Ingredient
}

func (self *IngredientValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(&(self.ingredientInfo))
	if err != nil {
		return err
	}
	self.ingredient.IngredientName = self.ingredientInfo.IngredientName

	if ingredientWeight, err := strconv.ParseFloat(self.ingredientInfo.IngredientWeight, 32); err != nil {
		return err
	} else {
		self.ingredient.IngredientWeight = float32(ingredientWeight)
	}

	if ingredientAmount, err := strconv.Atoi(self.ingredientInfo.IngredientAmount); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.ingredient.IngredientAmount = uint(ingredientAmount)
	}

	return nil
}
