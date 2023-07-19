package main

import (
	"github.com/gin-gonic/gin"
	"student-teach-api/config"
	"student-teach-api/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
