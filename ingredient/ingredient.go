package ingredient

import (
	"aamau/utils"
	"fmt"
	"strconv"
)

type Ingredient struct {
	IngredientId     uint    `gorm:"type:uint auto_increment; primary_key;column:ingredientId" form:"ingredientId" json:"ingredientId"`
	IngredientName   string  `gorm:"type:varchar(50) NOT NULL;column:ingredientName" form:"ingredientName" json:"ingredientName"`
	IngredientWeight float32 `gorm:"type:float NOT NULL;column:ingredientWeight" form:"ingredientWeight" json:"ingredientWeight"`
	IngredientAmount uint    `gorm:"type:uint NOT NULL;column:ingredientAmount" form:"ingredientAmount" json:"ingredientAmount"`
}

func Create_ingredient(ingredient Ingredient) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Ingredient").Model(&Ingredient{}).Create(&ingredient)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("create ingredient failed.")
	}
}

func Get_ingredient(ingredientId string) Ingredient {
	uid, err := strconv.Atoi(ingredientId)
	if err != nil {
		fmt.Println(err)
		return Ingredient{}
	}
	conn := utils.Get_connection()
	var ingredient Ingredient
	result := conn.Debug().Table("Ingredient").Model(&Ingredient{}).Where("ingredientId=?", uid).Find(&ingredient)
	if result.Error != nil || result.RowsAffected > 1 {
		fmt.Errorf("get ingredient failed.")
	}
	return ingredient

}
func Get_all_ingredients() []Ingredient {
	conn := utils.Get_connection()
	var ingredientList []Ingredient
	conn.Debug().Table("Ingredient").Model(&Ingredient{}).Find(&ingredientList)
	return ingredientList
}
func Update_ingredient(up_condition string, ingredient Ingredient) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Ingredient").Model(&Ingredient{}).Where(up_condition).Updates(ingredient)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("update ingredient failed.")
	}
}
func Delete_ingredient(rm_condition string) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Ingredient").Model(&Ingredient{}).Where(rm_condition).Delete(&Ingredient{})
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("delete ingredient failed.")
	}
}
