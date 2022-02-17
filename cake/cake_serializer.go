package cake

import "github.com/gin-gonic/gin"

type CakeSerializer struct {
	C *gin.Context
	Cake
}

type CakeResponse struct {
	CakeId    uint    `json:"-"`
	CakeName  string  `json:"cakeName"`
	DayNeeded uint    `json:"dayNeeded"`
	Price     float32 `json:"price"`
}

func (s *CakeSerializer) Response() CakeResponse {
	response := CakeResponse{
		CakeId:    s.CakeId,
		CakeName:  s.CakeName,
		DayNeeded: s.DayNeeded,
		Price:     s.Price,
	}
	return response
}
