package recipe

import "github.com/gin-gonic/gin"

type RecipeSerializer struct {
	C *gin.Context
	Recipe
}

type RecipeResponse struct {
	RecipeId       uint `json:"-"`
	CakeId         uint `json:"cakeId"`
	IngredientId   uint `json:"ingredientId"`
	IntAmtRequired uint `json:"ingredientAmountRequired"`
}

func (s *RecipeSerializer) Response() RecipeResponse {
	response := RecipeResponse{
		RecipeId:       s.RecipeId,
		CakeId:         s.CakeId,
		IngredientId:   s.IngredientId,
		IntAmtRequired: s.IntAmtRequired,
	}
	return response
}
