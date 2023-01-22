package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.GET("/status", func(c *gin.Context) {
		log.Println("status OK")
		c.Status(http.StatusOK)
	})

	return router
}
