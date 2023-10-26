package router

import (
	"assignment1/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	orderRouter := r.Group("/orders")
	{
		orderRouter.GET("/", controllers.GetOrders)
		orderRouter.POST("/", controllers.CreateOrder)
		orderRouter.GET("/:id", controllers.GetOrderByID)
		orderRouter.PUT("/:id", controllers.UpdateOrderByID)
		orderRouter.PUT("/", controllers.UpdateOrder)
		orderRouter.DELETE("/:id", controllers.DeleteOrderByID)
	}

	return r
}
