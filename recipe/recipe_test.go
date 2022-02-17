package recipe

import (
	"aamau/utils"
	"reflect"
	"testing"
)

func Test_crud(t *testing.T) {
	utils.CFG_path = "../cfg/config.yaml"
	test_recipe := Recipe{
		RecipeId:       42,
		CakeId:         1,
		IngredientId:   2,
		IntAmtRequired: 3,
	}

	Create_recipe(test_recipe)

	var found_recipe Recipe
	found_recipe = Get_recipe("42")
	if !reflect.DeepEqual(test_recipe, found_recipe) {
		t.Error("create recipe failed.")
	}

	updated_test_recipe := Recipe{
		RecipeId:       42,
		CakeId:         1,
		IngredientId:   2,
		IntAmtRequired: 10,
	}
	Update_recipe("recipeId=42", updated_test_recipe)
	found_recipe = Get_recipe("42")
	updated_test_recipe.RecipeId = 42
	if !reflect.DeepEqual(updated_test_recipe, found_recipe) {
		t.Error("update recipe failed.")
	}

	Delete_recipe("recipeId=42")
	found_recipe = Get_recipe("42")
	if !reflect.DeepEqual(Recipe{}, found_recipe) {
		t.Error("delete recipe failed.")
	}
}
