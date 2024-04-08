package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuttykarthikeyan/controllers"
	"github.com/kuttykarthikeyan/initializers"
)

func init() {
	initializers.LoadEnvvariables()
	initializers.Connect()
}
func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostsCreate)
	r.GET("/posts", controllers.PostList)
	r.GET("/posts/:id", controllers.Postid)
	r.PUT("/postupdate/:id", controllers.Postupdate)

	r.Run()
}
