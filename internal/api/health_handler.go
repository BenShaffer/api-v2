package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	prefix string
}

func NewHealthHandler(prefix string) HealthHandler {
	return HealthHandler{prefix}
}

func (hh HealthHandler) SetRoutes(router *gin.Engine) {
	health := router.Group(hh.prefix)
	{
		health.GET("", hh.GetHealth)
	}
}

func (hh HealthHandler) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Ok bool `json:"ok"`
	}{
		Ok: true,
	})
}
