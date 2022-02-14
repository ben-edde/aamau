package user

import (
	"aamau/utils"
	"fmt"
)

type User struct {
	UserId          uint   `gorm:"type:uint NOT NULL auto_increment; primary_key;column:userId"`
	UserName        string `gorm:"type:varchar(50) NOT NULL;column:userName" `
	ContactNo       string `gorm:"type:varchar(10) NOT NULL;column:contactNo" `
	Email           string `gorm:"type:varchar(100) NOT NULL;column:email" `
	DeliveryAddress string `gorm:"type:varchar(100) NOT NULL;column:deliveryAddress" `
}

func Get_user(uid int) User {
	conn := utils.Get_connection()
	var user User
	result := conn.Debug().Table("User").Model(&User{}).Where("userId=?", uid).Find(&user)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("get user failed.")
	}
	return user
}

func get_users() []User {
	conn := utils.Get_connection()
	var userlist []User
	conn.Debug().Table("User").Model(&User{}).Find(&userlist)
	return userlist
}

func create_user(user User) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("User").Model(&User{}).Create(&user)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("create user failed.")
	}
}

func update_user(up_condition, up_field, up_value string) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("User").Model(&User{}).Where(up_condition).Update(up_field, up_value)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("update user failed.")
	}
}

func delete_user(rm_condition string) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("User").Model(&User{}).Where(rm_condition).Delete(&User{})
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("delete user failed.")
	}
}
