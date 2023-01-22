package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorcz/ports/internal/controllers"
	"github.com/gorcz/ports/internal/services"
)

func NewRouter(portService services.Ports) *gin.Engine {
	router := gin.New()

	portController := controllers.NewPort(portService)

	router.GET("/status", func(c *gin.Context) {
		log.Println("status OK")
		c.Status(http.StatusOK)
	})

	router.POST("/port", func(c *gin.Context) {
		portController.UpsertPorts(c)
	})

	return router
}
