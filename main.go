package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func initializeRouter() *gin.Engine {
	router := gin.Default()

	fmt.Println("Registering mutate route...")
	router.POST("/mutate", Mutate)

	fmt.Println("Registering find route...")
	router.POST("/find", Find)

	return router
}

func main() {
	fmt.Println("Preparing service...")
	router := initializeRouter()

	fmt.Println("Starting service...")
	router.Run()
}
