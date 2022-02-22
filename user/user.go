package user

import (
	"aamau/utils"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type User struct {
	UserId          uint   `gorm:"type:uint auto_increment; primary_key;column:userId" form:"userId" json:"userId"`
	UserName        string `gorm:"type:varchar(50) NOT NULL;column:userName" form:"userName" json:"userName"`
	ContactNo       string `gorm:"type:varchar(10) NOT NULL;column:contactNo" form:"contactNo" json:"contactNo"`
	Email           string `gorm:"type:varchar(100) NOT NULL;column:email" form:"email" json:"email"`
	DeliveryAddress string `gorm:"type:varchar(100) NOT NULL;column:deliveryAddress" form:"deliveryAddress" json:"deliveryAddress"`
}

func Get_user(userId string) User {
	uid, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println(err)
		return User{}
	}
	conn := utils.Get_connection()
	var user User
	result := conn.Debug().Table("User").Model(&User{}).Where("userId=?", uid).Find(&user)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("get user failed.")
	}
	return user
}

func Get_all_users() []User {
	conn := utils.Get_connection()
	var userlist []User
	conn.Debug().Table("User").Model(&User{}).Find(&userlist)
	return userlist
}

func Create_user(conn *gorm.DB, user User) uint {
	result := conn.Debug().Table("User").Model(&User{}).Create(&user)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("create user failed.")
	}
	return user.UserId
}

func Update_user(up_condition string, user User) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("User").Model(&User{}).Where(up_condition).Updates(user)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("update user failed.")
	}
}

func Delete_user(rm_condition string) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("User").Model(&User{}).Where(rm_condition).Delete(&User{})
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("delete user failed.")
	}
}
