package main

import (
	"fmt"
	"log"

	"backend/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/health", handlers.Health)
	router.POST("/login", handlers.Login)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
