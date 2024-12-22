package main

import (
	"log"

	"github.com/Tholkappiar/go-crud.git/Database"
	"github.com/Tholkappiar/go-crud.git/controllers"
	"github.com/Tholkappiar/go-crud.git/middleware"
	"github.com/Tholkappiar/go-crud.git/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	server.POST("/user/register", controllers.Register)
	server.POST("/user/login", controllers.Login)
	
	server.GET("/blog/:blogId", controllers.GetBlog)
	server.GET("/blogs", controllers.GetBlogs)
	server.POST("/blog", middleware.ExtractEmailFromJWT(), controllers.PostBlog)
	server.PUT("/blog/:id", middleware.ExtractEmailFromJWT(), controllers.UpdateBlog)
	server.DELETE("/blog/:id", middleware.ExtractEmailFromJWT(), controllers.DeleteBlog)

	server.Run()
}
