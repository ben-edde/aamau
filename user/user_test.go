package user

import (
	"aamau/utils"
	"reflect"
	"testing"
)

func Test_get_user(t *testing.T) {
	utils.CFG_path = "../cfg/config.yaml"
	conn := utils.Get_connection()
	var found_users []User
	t.Logf("before: %v\n", found_users)
	// t.Logf("before: %b\n", found_users==nil)
	conn.Debug().Table("User").Model(&User{}).Where("userId=42").Find(&found_users)
	t.Logf("after: %v\n", found_users)
}

func Test_create_delete_user(t *testing.T) {
	utils.CFG_path = "../cfg/config.yaml"
	conn := utils.Get_connection()
	test_user := User{UserId: 42, UserName: "test_user", ContactNo: "1234", Email: "user@email.com", DeliveryAddress: "user home"}

	result := conn.Debug().Table("User").Model(&User{}).Create(&test_user)
	if result.Error != nil {
		t.Error("create user failed.")
	}
	var found_users []User
	result = conn.Debug().Table("User").Model(&User{}).Where("userId=42").Find(&found_users)
	if result.RowsAffected != 1 {
		t.Error("create user failed.")
	}
	if !reflect.DeepEqual(test_user, found_users[0]) {
		t.Error("create user failed.")
	}
	result = conn.Debug().Table("User").Model(&User{}).Where("userId=42").Delete(&User{})
	if result.RowsAffected != 1 {
		t.Error("delete user failed.")
	}
}
