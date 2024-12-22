package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

var client *supabase.Client

func init() {
	err1 := godotenv.Load()
    if err1 != nil {
        log.Fatalf("err loading: %v", err1)
    }
    var err error
    client, err = supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), &supabase.ClientOptions{})
    if err != nil {
        fmt.Println("Cannot initialize Supabase client:", err)
    }
    if client == nil {
        fmt.Println("Supabase client is nil")
    } else {
        fmt.Println("Supabase client initialized successfully")
    }
}

func Register(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}

	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Supabase client is not initialized",
		})
		return
	}

	
	userSession, err := client.Auth.Signup(types.SignupRequest{Email: body.Email,Password: body.Password})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sign up user with Supabase",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered",
		"session": userSession,
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid inputs",
		})
		return
	}

	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Supabase client is not initialized",
		})
		return
	}

	userSession, err := client.SignInWithEmailPassword(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session": userSession,
	})
}
