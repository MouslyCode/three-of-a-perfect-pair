package router

import (
	"github.com/MouslyCode/three-of-a-perfect-pair/backend/controller"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", controller.GetTasks)
	r.POST("/", controller.CreateTask)
	r.DELETE("/:id", controller.DeleteTask)
	r.PUT("/:id", controller.UpdateTask)
	r.PUT("/:id", controller.CompletedTask)
}
