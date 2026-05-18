package handlers

import (
	"net/http"
	"time"

	"backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]string{}
var mySigningKey = []byte("mysecretkey")

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func init() {
	hashedPassword, err := utils.HashPassword("admin")
	if err != nil {
		panic(err)
	}

	users["admin"] = hashedPassword
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func Login(c *gin.Context) {
	var user credentials
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Message: "Invalid request"})
		return
	}

	storedPassword, ok := users[user.Username]
	if !ok {
		c.JSON(http.StatusUnauthorized, errorResponse{Message: "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse{Message: "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Message: "Could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
