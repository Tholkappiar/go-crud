package main

import (
	db "github.com/Tholkappiar/go-crud.git/DB"
	"github.com/Tholkappiar/go-crud.git/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectToDB()
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

	server.Run()
}