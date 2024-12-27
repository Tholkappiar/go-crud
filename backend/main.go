package main

import (
	"log"
	"time"

	"github.com/Tholkappiar/go-crud.git/Database"
	"github.com/Tholkappiar/go-crud.git/controllers"
	"github.com/Tholkappiar/go-crud.git/middleware"
	"github.com/Tholkappiar/go-crud.git/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("err loading: %v", err)
    }
    
    Database.ConnectToDB()
    
    if err := Database.DB.AutoMigrate(&model.Blog{}); err != nil {
        log.Printf("Migration error: %v", err)
    }
    
}

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:  	  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	  }))

	server.POST("/user/register", controllers.Register)
	server.POST("/user/login", controllers.Login)
	
	server.GET("/blog/:blogId", controllers.GetBlog)
	server.GET("/blogs", controllers.GetBlogs)
	server.POST("/blog", middleware.ExtractEmailFromJWT(), controllers.PostBlog)
	server.PUT("/blog/:id", middleware.ExtractEmailFromJWT(), controllers.UpdateBlog)
	server.DELETE("/blog/:id", middleware.ExtractEmailFromJWT(), controllers.DeleteBlog)

	server.Run()
}
