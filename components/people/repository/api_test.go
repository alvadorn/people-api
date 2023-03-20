package repository

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvadorn/people-api/components/people/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type apiRepositorySuite struct {
	suite.Suite
	mux        *http.ServeMux
	server     *httptest.Server
	repository *apiClient
	ctx        context.Context
}

func TestNewApiRepository(t *testing.T) {
	t.Run(
		"instantiate redis repository successfully", func(t *testing.T) {
			repo := NewApiClient("api_url")
			assert.NotNil(t, repo)
		})
}

func (s *apiRepositorySuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *apiRepositorySuite) SetupTest() {
	s.mux = http.NewServeMux()
	s.server = httptest.NewServer(s.mux)
	s.repository = NewApiClient(s.server.URL + "/%s")
}

func (s *apiRepositorySuite) SetupSubTest() {
	s.SetupTest()
}

func (s *apiRepositorySuite) TestGetByName() {
	s.Run(
		"retrieve by name successfully", func() {
			s.mux.HandleFunc(
				"/yoshua_bengio", func(w http.ResponseWriter, r *http.Request) {
					w.Write(
						[]byte(`{ "query": { "pages": [{
					"title": "Yoshua Bengio",
					"revisions": [{
						"slots": { "main": {
							"content": "{{Short description|computer scientist}}"
							} } }]
					}]}}`))
				})
			output, err := s.repository.getByName(s.ctx, "yoshua_bengio")
			s.Nil(err)
			s.Equal(
				&domain.Person{
					Name:             "Yoshua Bengio",
					ShortDescription: "computer scientist",
				}, output)
		})

	s.Run(
		"retrieve by name successfully but without short description", func() {
			s.mux.HandleFunc(
				"/yoshua_bengio", func(w http.ResponseWriter, r *http.Request) {
					w.Write(
						[]byte(`{ "query": { "pages": [{
					"title": "Yoshua Bengio",
					"revisions": [{
						"slots": { "main": {
							"content": ""
							} } }]
					}]}}`))
				})
			output, err := s.repository.getByName(s.ctx, "yoshua_bengio")
			s.Nil(err)
			s.Equal(
				&domain.Person{
					Name:             "Yoshua Bengio",
					ShortDescription: "",
				}, output)
		})

	s.Run(
		"does not find person", func() {
			s.mux.HandleFunc(
				"/yoshua_bengio", func(w http.ResponseWriter, r *http.Request) {
					w.Write(
						[]byte(`{ "query": { "pages": [{
					"title": "Yoshua Bengio",
					"missing": true
					}]}}`))
				})
			output, err := s.repository.getByName(s.ctx, "yoshua_bengio")
			s.Nil(output)
			s.Error(err)
			s.ErrorContains(err, "not_found_error")
		})
}

func TestApiClientRepositorySuite(t *testing.T) {
	suite.Run(t, new(apiRepositorySuite))
}
