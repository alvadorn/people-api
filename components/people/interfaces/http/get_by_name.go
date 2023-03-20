package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/alvadorn/people-api/components/people/application"
	"github.com/gin-gonic/gin"
)

const nameParameterKey = "name"

func (hi *httpInterfaces) getByName(ctx *gin.Context) {
	name := ctx.Param(nameParameterKey)

	nameWithoutSpaces := strings.ReplaceAll(name, " ", "_")
	person, err := hi.peopleUseCases.GetByName(ctx, nameWithoutSpaces)

	if err != nil {
		if errors.Is(err, application.NotFoundErr) {
			ctx.AbortWithStatusJSON(
				http.StatusNotFound, Error{
					Code:    http.StatusNotFound,
					Message: "entity not found",
					Type:    err.Error(),
				})
		} else {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError, Error{
					Code:    http.StatusInternalServerError,
					Message: "check logs for details",
					Type:    err.Error(),
				})
		}
		return
	}

	ctx.JSON(http.StatusOK, person)
}
