package main

import (
	"github.com/Tholkappiar/go-crud.git/Database"
	"github.com/Tholkappiar/go-crud.git/controllers"
	"github.com/Tholkappiar/go-crud.git/model"
	"github.com/gin-gonic/gin"
)

func init() {
	Database.ConnectToDB()
	Database.DB.AutoMigrate(&model.Blog{}, &model.User{})
}


func main() {
	server := gin.Default();
	server.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is thols",
		})
	})

	server.GET("/blogs", controllers.GetBlog)
	server.POST("/blog", controllers.PostBlog)
	server.PUT("/blog/:id", controllers.UpdateBlog)
	server.DELETE("/blog/:id", controllers.DeleteBlog)


	server.POST("/user/register", controllers.Register)
	server.POST("/user/login", controllers.Login)

	server.Run()
}