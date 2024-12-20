package controllers

import (
	"fmt"
	"net/http"

	"github.com/Tholkappiar/go-crud.git/Database"
	"github.com/Tholkappiar/go-crud.git/model"
	"github.com/gin-gonic/gin"
)

func GetBlog(c *gin.Context)  {

	var blogs []model.Blog
	result := Database.DB.Find(&blogs)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"blogs": blogs,
	})
}

func PostBlog(c *gin.Context) {
	var blog model.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

	result := Database.DB.Create(&blog)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
	})
}

func DeleteBlog(c *gin.Context) {

	blogId := c.Param("id")
	var blog model.Blog

	fmt.Println("Blog id : " , blogId)
	if err := Database.DB.First(&blog,blogId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Blog not Found",
		})
		return
	}
	

	if err := Database.DB.Delete(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Deleting Blog",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog Deleted",
	})
}

func UpdateBlog(c *gin.Context) {
	param := c.Param("id")


	var updatedBlog model.Blog
	if err := c.ShouldBindJSON(&updatedBlog); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Provide valid data to update",
		})
		return
	}

	var blog model.Blog
	if err := Database.DB.First(&blog, param).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Blog not Found",
		})
		return
	}

	blog.Title = updatedBlog.Title
	blog.Description = updatedBlog.Description
	if err := Database.DB.Save(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog updated successfully",
		"blog":    blog,
	})
}