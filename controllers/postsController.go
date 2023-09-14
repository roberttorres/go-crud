package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body) //Remove this later to see what happens

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsShow(c *gin.Context) {

	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})

}

func PostsIndex(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	initializers.DB.Delete(&post, id)

	c.Status(200)

}
