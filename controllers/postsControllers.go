package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuttykarthikeyan/initializers"
	"github.com/kuttykarthikeyan/models"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"posts": "Posted successfully",
	})
}
func PostList(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}
func Postid(c *gin.Context) {
	id := c.Param("id")
	var post []models.Post

	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"posts": post,
	})
}
func Postupdate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Model(post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	var updated models.Post
	initializers.DB.First(&updated, id)
	c.JSON(200, gin.H{
		"post": post,
	})

}
