package user

import "github.com/gin-gonic/gin"

type UserSerializer struct {
	C *gin.Context
	User
}

type UserResponse struct {
	UserId          uint   `json:"-"`
	UserName        string `json:"userName"`
	ContactNo       string `json:"contactNo"`
	Email           string `json:"email"`
	DeliveryAddress string `json:"deliveryAddress"`
}

func (s *UserSerializer) Response() UserResponse {
	response := UserResponse{
		UserId:          s.UserId,
		UserName:        s.UserName,
		ContactNo:       s.ContactNo,
		Email:           s.Email,
		DeliveryAddress: s.DeliveryAddress,
	}
	return response
}
