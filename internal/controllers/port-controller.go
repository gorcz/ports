package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorcz/ports/internal/services"
	"github.com/gorcz/ports/pkg/model/port"
)

type Controller interface {
	UpsertPorts(c *gin.Context)
}

type Port struct {
	service services.Ports
}

func NewPort(portService services.Ports) *Port {
	return &Port{
		service: portService,
	}
}

func (pc *Port) UpsertPorts(c *gin.Context) {
	bodyReader := c.Request.Body
	portIterator, err := port.ParsePortsFromJSONMap(bodyReader)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err = pc.service.UpsertPorts(c, portIterator); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
