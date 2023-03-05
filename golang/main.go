package main

import (
	"github.com/gin-gonic/gin"
	"golang/models"
	"golang/routing"
)

func main() {
	app := gin.Default()
	models.TaskStorage = make([]models.Task, 0)
	routing.TaskRouting(app)

	err := app.Run()
	if err != nil {
		print(err)
	}
}
