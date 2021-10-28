package api

import (
	"api/internal/application"
	"api/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	prefix        string
	logger        log.IApiLogger
	personService application.IPersonService
}

func NewPersonHandler(prefix string, logger log.IApiLogger, personService application.IPersonService) *PersonHandler {
	return &PersonHandler{prefix, logger, personService}
}

func (ph PersonHandler) SetRoutes(router *gin.RouterGroup) {
	people := router.Group(ph.prefix)
	{
		people.GET("", ph.GetPeople)
	}
}

func (ph *PersonHandler) GetPeople(c *gin.Context) {
	people := ph.personService.GetPeople()

	c.JSON(http.StatusOK, people)
}
