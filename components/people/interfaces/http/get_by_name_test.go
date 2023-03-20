package http

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/alvadorn/people-api/components/people/application"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func (s *httpSuite) TestGetByName() {
	s.Run(
		"successfully", func() {
			output := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(output)
			ctx.AddParam("name", "Yoshua Bengio")
			s.useCasesMock.
				On("GetByName", mock.AnythingOfType("*gin.Context"), "Yoshua_Bengio").
				Once().
				Return(&application.PersonDTO{Name: "Yoshua Bengio"}, nil)
			s.httpInterface.getByName(ctx)
			s.Equal(output.Code, http.StatusOK)
			data, _ := io.ReadAll(output.Body)
			s.Equal("{\"name\":\"Yoshua Bengio\",\"short_description\":null}", string(data))
		})

	s.Run(
		"not found", func() {
			output := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(output)
			ctx.AddParam("name", "Yoshua Bengio")
			s.useCasesMock.
				On("GetByName", mock.AnythingOfType("*gin.Context"), "Yoshua_Bengio").
				Once().
				Return(nil, application.NotFoundErr)
			s.httpInterface.getByName(ctx)
			s.Equal(output.Code, http.StatusNotFound)
			data, _ := io.ReadAll(output.Body)
			s.Equal("{\"code\":404,\"message\":\"entity not found\",\"type\":\"not_found_error\"}", string(data))
		})

	s.Run(
		"unexpected error", func() {
			output := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(output)
			ctx.AddParam("name", "Yoshua Bengio")
			s.useCasesMock.
				On("GetByName", mock.AnythingOfType("*gin.Context"), "Yoshua_Bengio").
				Once().
				Return(nil, application.UnexpectedError)
			s.httpInterface.getByName(ctx)
			s.Equal(output.Code, http.StatusInternalServerError)
			data, _ := io.ReadAll(output.Body)
			s.Equal("{\"code\":500,\"message\":\"check logs for details\",\"type\":\"unexpected_error\"}", string(data))
		})
}
