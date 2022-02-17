package cake

import (
	"aamau/utils"
	"fmt"
	"strconv"
)

type Cake struct {
	CakeId    uint    `gorm:"type:uint auto_increment; primary_key;column:cakeId" form:"cakeId" json:"cakeId"`
	CakeName  string  `gorm:"type:varchar(50) NOT NULL;column:cakeName" form:"cakeName" json:"cakeName"`
	DayNeeded uint    `gorm:"type:uint NOT NULL;column:dayNeeded" form:"dayNeeded" json:"dayNeeded"`
	Price     float32 `gorm:"type:float NOT NULL;column:price" form:"price" json:"price"`
}

func Create_cake(cake Cake) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Cake").Model(&Cake{}).Create(&cake)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("create cake failed.")
	}
}
func Get_cake(cakeId string) Cake {
	uid, err := strconv.Atoi(cakeId)
	if err != nil {
		fmt.Println(err)
		return Cake{}
	}
	conn := utils.Get_connection()
	var cake Cake
	result := conn.Debug().Table("Cake").Model(&Cake{}).Where("cakeId=?", uid).Find(&cake)
	if result.Error != nil || result.RowsAffected > 1 {
		fmt.Errorf("get cake failed.")
	}
	return cake

}
func Get_all_cakes() []Cake {
	conn := utils.Get_connection()
	var cakeList []Cake
	conn.Debug().Table("Cake").Model(&Cake{}).Find(&cakeList)
	return cakeList
}
func Update_cake(up_condition string, cake Cake) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Cake").Model(&Cake{}).Where(up_condition).Updates(cake)
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("update cake failed.")
	}
}
func Delete_cake(rm_condition string) {
	conn := utils.Get_connection()
	result := conn.Debug().Table("Cake").Model(&Cake{}).Where(rm_condition).Delete(&Cake{})
	if result.Error != nil || result.RowsAffected != 1 {
		fmt.Errorf("delete cake failed.")
	}
}
