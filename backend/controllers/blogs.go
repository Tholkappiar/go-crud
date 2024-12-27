package controllers

import (
	"fmt"
	"net/http"

	"github.com/Tholkappiar/go-crud.git/Database"
	"github.com/Tholkappiar/go-crud.git/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserIDByEmail(c *gin.Context) (uuid.UUID, error) {
	email, exists := c.Get("email")
	if !exists {
		return uuid.UUID{}, fmt.Errorf("email is missing")
	}

	strEmail, ok := email.(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("invalid email")
	}

	var user model.User
	if err := Database.DB.First(&user, "email = ?", strEmail).Error; err != nil {
		return uuid.UUID{}, fmt.Errorf("user not found")
	}

	return user.ID, nil
}

func GetBlogs(c *gin.Context) {
	var blogs []model.Blog
	result := Database.DB.Find(&blogs)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blogs": blogs,
	})
}

func GetBlog(c *gin.Context) {
	blogId := c.Param("blogId");
	fmt.Println(blogId)
	var blog model.Blog
	result := Database.DB.First(&blog, "id = ? ",  blogId)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blogs": blog,
	})
}

func PostBlog(c *gin.Context) {
	var blog model.Blog

	userID, err := GetUserIDByEmail(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := Database.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.UserID = userID

	result := Database.DB.Create(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Statement.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"blog":    blog,
	})
}

func DeleteBlog(c *gin.Context) {
	blogId := c.Param("id")
	var blog model.Blog

	if err := Database.DB.First(&blog, blogId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	userID, err := GetUserIDByEmail(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if blog.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this blog post"})
		return
	}

	if err := Database.DB.Delete(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog deleted successfully",
	})
}

func UpdateBlog(c *gin.Context) {
	blogId := c.Param("id")
	var updatedBlog model.Blog

	if err := c.ShouldBindJSON(&updatedBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var blog model.Blog
	if err := Database.DB.First(&blog, blogId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	userID, err := GetUserIDByEmail(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(blog.UserID)
	fmt.Println(userID)
	if blog.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this blog"})
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
