package ingredient

import "github.com/gin-gonic/gin"

type IngredientSerializer struct {
	C *gin.Context
	Ingredient
}

type IngredientResponse struct {
	IngredientId     uint    `json:"-"`
	IngredientName   string  `json:"ingredientName"`
	IngredientWeight float32 `json:"ingredientWeight"`
	IngredientAmount uint    `json:"ingredientAmount"`
}

func (s *IngredientSerializer) Response() IngredientResponse {
	response := IngredientResponse{
		IngredientId:     s.IngredientId,
		IngredientName:   s.IngredientName,
		IngredientWeight: s.IngredientWeight,
		IngredientAmount: s.IngredientAmount,
	}
	return response
}
