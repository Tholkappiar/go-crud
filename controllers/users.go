package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Tholkappiar/go-crud.git/Database"
	"github.com/Tholkappiar/go-crud.git/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)


func Register(c *gin.Context) {
	
	var body struct {
		Email string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Failed to read body",
		})
		return
	}
	
	hash , err := bcrypt.GenerateFromPassword([]byte(body.Password),10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Failed to Hash Password",
		})
		return
	}
	
	user := model.User{Email: body.Email,Password: string(hash)}

	if err := Database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Error while creating user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H {
		"message": "Successfully registered",
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Invalid inputs",
		})
		return
	}


	var user model.User
	
	if err := Database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H {
			"message": "Invalid Email or Password",
		})
		return
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": input.Email,
		"exp": time.Now().Add(24 * 60 * 60).Unix(),
	})
	
	SECRET_KEY := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H {
			"message": "Error while creating the token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"token": tokenString,
	})
}