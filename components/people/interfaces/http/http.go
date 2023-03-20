package http

import (
	"fmt"

	"github.com/alvadorn/people-api/components/people/application"
	"github.com/gin-gonic/gin"
)

type httpInterfaces struct {
	peopleUseCases application.UseCases
}

type HTTP interface {
	RegisterRoutes(r *gin.RouterGroup)
}

func New(cases application.UseCases) *httpInterfaces {
	return &httpInterfaces{peopleUseCases: cases}
}

func (hi *httpInterfaces) RegisterRoutes(r *gin.RouterGroup) {
	peopleGroup := r.Group("/people")
	peopleGroup.GET(fmt.Sprint("/:", nameParameterKey), hi.getByName)
}
