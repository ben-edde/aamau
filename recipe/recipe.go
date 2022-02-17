package recipe

import (
	"aamau/utils"
	"fmt"
	"strconv"
)

type Recipe struct {
	RecipeId       uint `gorm:"type:uint auto_increment; primary_key;column:recipeId" form:"recipeId" json:"recipeId"`
	CakeId         uint `gorm:"type:uint NOT NULL;column:cakeId" form:"cakeId" json:"cakeId"`
	IngredientId   uint `gorm:"type:uint NOT NULL;column:ingredientId" form:"ingredientId" json:"ingredientId"`
	IntAmtRequired uint `gorm:"type:uint NOT NULL;column:ingredientAmountRequired" form:"ingredientAmountRequired" json:"ingredientAmountRequired"`
}

func Create_recipe(recipe Recipe) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Recipe").Model(&Recipe{}).Create(&recipe)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("create recipe failed.")
	}
}

func Get_recipe(recipeId string) Recipe {
	uid, err := strconv.Atoi(recipeId)
	if err != nil {
		fmt.Println(err)
		return Recipe{}
	}
	conn := utils.Get_connection()
	var recipe Recipe
	result := conn.Debug().Table("Recipe").Model(&Recipe{}).Where("recipeId=?", uid).Find(&recipe)
	if result.Error != nil || result.RowsAffected > 1 {
		fmt.Errorf("get recipe failed.")
	}
	return recipe

}
func Get_all_recipes() []Recipe {
	conn := utils.Get_connection()
	var recipeList []Recipe
	conn.Debug().Table("Recipe").Model(&Recipe{}).Find(&recipeList)
	return recipeList
}
func Update_recipe(up_condition string, recipe Recipe) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Recipe").Model(&Recipe{}).Where(up_condition).Updates(recipe)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("update recipe failed.")
	}
}
func Delete_recipe(rm_condition string) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Recipe").Model(&Recipe{}).Where(rm_condition).Delete(&Recipe{})
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("delete recipe failed.")
	}
}
