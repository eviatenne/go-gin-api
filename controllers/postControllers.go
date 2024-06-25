package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/initializers"
	"go-gin-api/models"
)

func PostsCreate(c *gin.Context) {
	//GET data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get The Post
	var post []models.Post
	// Get all records
	initializers.DB.Find(&post)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostsShow(c *gin.Context) {

	id := c.Param("id")
	//Get The Post
	var posts []models.Post

	// Get record with specific id
	initializers.DB.First(&posts, id)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsUpdate(c *gin.Context) {

	//Get The id
	id := c.Param("id")

	//Get The body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Find
	var post []models.Post
	initializers.DB.First(&post, id)

	//Update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostsDelete(c *gin.Context) {
	//Get the ID from the URL
	id := c.Param("id")

	// Get record with specific id
	initializers.DB.Delete(&models.Post{}, id)

	//Respond 200
	c.Status(200)

}
