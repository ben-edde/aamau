package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"aamau/cake"
	"aamau/ingredient"
	"aamau/order"
	"aamau/recipe"
	"aamau/user"
	"aamau/utils"

	"github.com/gin-gonic/gin"
)

func start() {
	utils.CFG_path = "cfg/config.yaml"
	exist, err := utils.Check_path_exists(utils.CFG_path)
	if err != nil || !exist {
		fmt.Printf("Config directory path (%s) exists: %v\nError: %v\n", utils.CFG_path, exist, err)
		return
	}
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)

	router := gin.Default()
	api_router_group := router.Group("/api")
	user.UserAPIRegister(api_router_group.Group("/user"))
	cake.CakeAPIRegister(api_router_group.Group("/cake"))
	order.OrderAPIRegister(api_router_group.Group("/order"))
	ingredient.IngredientAPIRegister(api_router_group.Group("/ingredient"))
	recipe.RecipeAPIRegister(api_router_group.Group("/recipe"))
	router.LoadHTMLGlob("view/index.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(listenAddr)
}

func main() {
	start()
}
