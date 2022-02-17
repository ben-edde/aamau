package cake

import (
	"aamau/utils"
	"reflect"
	"testing"
)

func Test_crud(t *testing.T) {
	utils.CFG_path = "../cfg/config.yaml"
	test_cake := Cake{CakeId: 42, CakeName: "test_cake", DayNeeded: 3, Price: 42.1}

	Create_cake(test_cake)

	var found_cake Cake
	found_cake = Get_cake("42")
	if !reflect.DeepEqual(test_cake, found_cake) {
		t.Error("create cake failed.")
	}

	updated_test_cake := Cake{CakeName: "test_cake", DayNeeded: 1, Price: 42.1}
	Update_cake("cakeId=42", updated_test_cake)
	found_cake = Get_cake("42")
	updated_test_cake.CakeId = 42
	if !reflect.DeepEqual(updated_test_cake, found_cake) {
		t.Error("update cake failed.")
	}

	Delete_cake("cakeId=42")
	found_cake = Get_cake("42")
	if !reflect.DeepEqual(Cake{}, found_cake) {
		t.Error("update cake failed.")
	}
}
