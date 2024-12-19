package controllers

import (
	"net/http"

	db "github.com/Tholkappiar/go-crud.git/DB"
	"github.com/Tholkappiar/go-crud.git/model"
	"github.com/gin-gonic/gin"
)

func GetBlog(c *gin.Context)  {

	var blogs []model.Blog
	result := db.DB.Find(&blogs)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK,gin.H{
		"blogs": blogs,
	})
}

func PostBlog(c *gin.Context) {
	var blog model.Blog

	if err := c.ShouldBindBodyWithJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.DB.Create(&blog)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
	})
}