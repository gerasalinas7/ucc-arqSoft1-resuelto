package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"auth-intercambiable-ejercicio/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	_ = godotenv.Load()

	mode := os.Getenv("AUTH_MODE")
	if mode == "" {
		mode = "jwt"
	}

	handler, err := newServer(mode)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Servidor listo en http://localhost:8080 usando auth=%s", mode)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func newServer(mode string) (http.Handler, error) {
	provider, err := auth.New(mode, os.Getenv("AUTH_SECRET"))
	if err != nil {
		return nil, err
	}

	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "json invalido"})
			return
		}

		token, err := provider.Login(req.Username, req.Password)
		if err != nil {
			if errors.Is(err, auth.ErrInvalidCredentials) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales invalidas"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":     token,
			"auth_type": provider.Name(),
		})
	})

	router.GET("/profile", func(c *gin.Context) {
		token := bearerToken(c.GetHeader("Authorization"))
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "falta bearer token"})
			return
		}

		username, err := provider.Validate(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":      username,
			"auth_type": provider.Name(),
		})
	})

	return router, nil
}

func bearerToken(value string) string {
	if !strings.HasPrefix(value, "Bearer ") {
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(value, "Bearer "))
}
