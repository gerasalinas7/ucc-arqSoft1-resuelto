package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"vinyl-store/handlers"
)

func main() {
	// Configuración del router
	router := gin.Default()

	// Endpoints
	router.GET("/albums", handlers.GetAllAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.AddAlbum)
	router.PUT("/albums/:id", handlers.UpdateAlbum)
	router.DELETE("/albums/:id", handlers.DeleteAlbum)

	// Iniciar el servidor en el puerto 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
