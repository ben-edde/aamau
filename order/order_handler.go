package order

import (
	"fmt"

	"gorm.io/gorm"
)

func get_cake_recipe_list(conn *gorm.DB, cakeId int) map[int]int {
	ingredient_amount_map := map[int]int{}
	db, _ := conn.DB()
	rows, err := db.Query(fmt.Sprintf("SELECT ingredientId,ingredientAmountRequired FROM Recipe WHERE cakeId=%d", cakeId))
	if err != nil {
		fmt.Printf("get recipe failed.")
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ingredientId             int
			ingredientAmountRequired int
		)
		if err := rows.Scan(&ingredientId, &ingredientAmountRequired); err != nil {
			fmt.Println(err)
		}
		ingredient_amount_map[ingredientId] = ingredientAmountRequired
	}
	return ingredient_amount_map
}

func check_enough_ingredient(conn *gorm.DB, recipeMap map[int]int, orderAmount int) (bool, error) {
	db, _ := conn.DB()
	for iid, amt := range recipeMap {
		rows, err := db.Query(fmt.Sprintf("SELECT ingredientAmount FROM Ingredient WHERE ingredientId=%d", iid))
		if err != nil {
			fmt.Printf("get ingredient amount failed.")
			return false, err
		}

		for rows.Next() {
			var ingredientAvailable int
			if err := rows.Scan(&ingredientAvailable); err != nil {
				fmt.Println(err)
				return false, err
			}
			if ingredientAvailable < amt*orderAmount {
				fmt.Printf("Lack of ingredient\n")
				return false, nil
			}
		}
		rows.Close()
	}
	return true, nil
}
