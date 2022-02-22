package order

import (
	"aamau/user"
	"fmt"
	"time"

	"gorm.io/datatypes"
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

func get_delivery_date(conn *gorm.DB, orderDate datatypes.Date, cakeId, amount uint) (*datatypes.Date, error) {
	db, _ := conn.DB()
	rows, err := db.Query(fmt.Sprintf("SELECT dayNeeded FROM Cake WHERE cakeId=%d", cakeId))
	if err != nil {
		fmt.Printf("get dayNeeded failed.")
	}
	defer rows.Close()
	if rows.Next() {
		var dayNeeded int
		if err := rows.Scan(&dayNeeded); err == nil {
			deliveryDate := time.Time(orderDate).AddDate(0, 0, dayNeeded*int(amount))
			resultDate := datatypes.Date(deliveryDate)
			return &resultDate, nil
		}
		fmt.Println(err)
	}
	return nil, err
}

func calculate_total_price(conn *gorm.DB, cakeId, amount uint) float32 {
	db, _ := conn.DB()
	rows, err := db.Query(fmt.Sprintf("SELECT price FROM Cake WHERE cakeId=%d", cakeId))
	if err != nil {
		fmt.Printf("get price failed.")
	}
	defer rows.Close()
	if rows.Next() {
		var price float32
		if err := rows.Scan(&price); err == nil {
			return price * float32(amount)
		}
		fmt.Println(err)
	}
	return -1
}

func handle_order(conn *gorm.DB, new_order RawOrderValidator) (*Order, bool) {
	cakeRecipe := get_cake_recipe_list(conn, int(new_order.order.CakeId))
	enough, err := check_enough_ingredient(conn, cakeRecipe, int(new_order.order.Amount))
	if err != nil {
		fmt.Printf("Failed checking ingredient: %s\n", err)
	} else if !enough {
		fmt.Printf("No enough ingredient: %s\n", err)
	} else {
		if date, err := get_delivery_date(conn, new_order.order.OrderDate, new_order.order.CakeId, new_order.order.Amount); err != nil || date == nil {
			fmt.Printf("Delivery date unavilable: %s\n", err)
		} else {
			new_order.user.UserId = user.Create_user(conn, new_order.user)
			new_order.order.DeliveryDate = *date
			new_order.order.UserId = new_order.user.UserId
			new_order.order.TotalPrice = calculate_total_price(conn, new_order.order.CakeId, new_order.order.Amount)
			return &new_order.order, true
		}
	}
	return nil, false
}
