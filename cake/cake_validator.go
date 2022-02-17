package cake

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CakeInfo struct {
	CakeName  string `form:"cakeName" json:"cakeName" binding:"max=255"`
	DayNeeded string `form:"dayNeeded" json:"dayNeeded" binding:"max=255"`
	Price     string `form:"price" json:"price" binding:"max=255"`
}

type CakeValidator struct {
	cakeInfo CakeInfo
	cake     Cake
}

func (self *CakeValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(&(self.cakeInfo))
	if err != nil {
		return err
	}
	self.cake.CakeName = self.cakeInfo.CakeName
	if dayNeeded, err := strconv.Atoi(self.cakeInfo.DayNeeded); err != nil {
		fmt.Println(err)
		return err
	} else {
		self.cake.DayNeeded = uint(dayNeeded)
	}

	if price, err := strconv.ParseFloat(self.cakeInfo.Price, 32); err != nil {
		return err
	} else {
		self.cake.Price = float32(price)
	}

	return nil
}
