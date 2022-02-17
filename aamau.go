package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

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
	// router.LoadHTMLGlob("view/*")
	user.UserAPIRegister(api_router_group.Group("/user"))
	router.GET("/api/hello", func(c *gin.Context) {
		message := "Hello World!\n"
		content := c.Request.URL.Query().Get("name")
		if content != "" {
			message = fmt.Sprintf("Hello: %s\n", content)
		}
		c.JSON(200, gin.H{
			"message": message,
		})
	})
	router.GET("/api/echo", func(c *gin.Context) {
		message := "Empty\n"
		content := c.Request.URL.Query().Get("content")
		if content != "" {
			message = fmt.Sprintf("Echo: %s\n", content)
		}
		c.JSON(200, gin.H{
			"message": message,
		})
	})

	router.GET("/api/os", func(c *gin.Context) {
		c.String(200, runtime.GOOS)
	})
	// router.GET("/api/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })
	router.Run(listenAddr)
}
func main() {
	start()
}
