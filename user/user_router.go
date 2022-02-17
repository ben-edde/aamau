package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAPIRegister(router *gin.RouterGroup) {
	router.GET("/all", get_all_users)
	router.GET("/:userId", get_user_by_id)
	router.POST("/", create_user)
	router.DELETE("/:userId", delete_user_by_id)
	router.PUT("/:userId", update_user_by_id)
}

func get_user_by_id(c *gin.Context) {
	userId := c.Param("userId")
	found_user := Get_user(userId)
	userSerializer := UserSerializer{c, found_user}
	c.JSON(http.StatusOK, gin.H{"userId": userId, "user": userSerializer.Response()})
}

func get_all_users(c *gin.Context) {
	response_list := []UserResponse{}
	found_users := Get_all_users()
	for _, user := range found_users {
		userSerializer := UserSerializer{c, user}
		response_list = append(response_list, userSerializer.Response())
	}
	c.JSON(http.StatusOK, gin.H{"users": response_list})
}

func create_user(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
	}
	Create_user(user)
	response := fmt.Sprintf("Created user: %v", user)
	c.JSON(http.StatusOK, gin.H{"content": response})
}

func delete_user_by_id(c *gin.Context) {
	userId := c.Param("userId")
	found_user := Get_user(userId)
	Delete_user(fmt.Sprintf("userId=%s", userId))
	userSerializer := UserSerializer{c, found_user}
	c.JSON(http.StatusOK, gin.H{"userId": userId, "deleted_user": userSerializer.Response()})
}

func update_user_by_id(c *gin.Context) {
	userId := c.Param("userId")
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
	}
	Update_user(fmt.Sprintf("userId=%s", userId), user)
	userSerializer := UserSerializer{c, user}
	c.JSON(http.StatusOK, gin.H{"userId": userId, "updated_user": userSerializer.Response()})
}
