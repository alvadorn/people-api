package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvadorn/people-api/components/people/application"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type httpSuite struct {
	suite.Suite

	httpInterface *httpInterfaces
	useCasesMock  *application.UseCasesMock
}

func TestHttpSuite(t *testing.T) {
	suite.Run(t, new(httpSuite))
}

func (s *httpSuite) SetupSuite() {
	s.useCasesMock = application.NewUseCases(s.T())
	s.httpInterface = &httpInterfaces{peopleUseCases: s.useCasesMock}
}

func (s *httpSuite) TestRegisterRoutes() {
	output := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(output)
	rg := engine.Group("")
	s.httpInterface.RegisterRoutes(rg)

	s.Equal(http.MethodGet, engine.Routes()[0].Method)
	s.Equal("/people/:name", engine.Routes()[0].Path)
}
