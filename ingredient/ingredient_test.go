package ingredient

import (
	"aamau/utils"
	"reflect"
	"testing"
)

func Test_crud(t *testing.T) {
	utils.CFG_path = "../cfg/config.yaml"
	test_ingredient := Ingredient{
		IngredientId:     42,
		IngredientName:   "test_ingredient",
		IngredientWeight: 3.14,
		IngredientAmount: 30,
	}

	Create_ingredient(test_ingredient)

	var found_ingredient Ingredient
	found_ingredient = Get_ingredient("42")
	if !reflect.DeepEqual(test_ingredient, found_ingredient) {
		t.Error("create ingredient failed.")
	}

	updated_test_ingredient := Ingredient{
		IngredientId:     42,
		IngredientName:   "test_ingredient",
		IngredientWeight: 3.14,
		IngredientAmount: 50,
	}
	Update_ingredient("ingredientId=42", updated_test_ingredient)
	found_ingredient = Get_ingredient("42")
	updated_test_ingredient.IngredientId = 42
	if !reflect.DeepEqual(updated_test_ingredient, found_ingredient) {
		t.Error("update ingredient failed.")
	}

	Delete_ingredient("ingredientId=42")
	found_ingredient = Get_ingredient("42")
	if !reflect.DeepEqual(Ingredient{}, found_ingredient) {
		t.Error("delete ingredient failed.")
	}
}
