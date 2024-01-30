package main

import (
	"github.com/khagendra01/go_app/initializers"
	"github.com/khagendra01/go_app/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
