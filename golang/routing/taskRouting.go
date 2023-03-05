package routing

import (
	"github.com/gin-gonic/gin"
	"golang/controllers"
)

func TaskRouting(app *gin.Engine) {
	app.GET("/tasks", controllers.GetAll)
	app.GET("/tasks/:id", controllers.GetById)
	app.POST("/tasks", controllers.Add)
	app.PUT("/tasks/:id/:idDone", controllers.UpdateDone)
	app.DELETE("/tasks/:id", controllers.DeleteById)
}
