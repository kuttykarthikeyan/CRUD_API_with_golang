package main

import (
	"github.com/kuttykarthikeyan/initializers"
	"github.com/kuttykarthikeyan/models"
)

func init() {
	initializers.LoadEnvvariables()
	initializers.Connect()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
