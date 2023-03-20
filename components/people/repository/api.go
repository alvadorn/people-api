package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alvadorn/people-api/components/people/domain"
)

const (
	shortDescriptionTag       = "{{Short description|"
	shortDescriptionSkipRunes = 20
	shortDescriptionEndRune   = '}'
)

type apiClient struct {
	client     *http.Client
	baseApiUrl string
}

type apiClientPeopleOperator interface {
	personGetter
}

func NewApiClient(baseAPIUrl string) *apiClient {
	return &apiClient{
		client:     &http.Client{},
		baseApiUrl: baseAPIUrl,
	}
}

type apiData struct {
	Query struct {
		Pages []pageData `json:"pages"`
	} `json:"query"`
}

type pageData struct {
	Title     string `json:"title"`
	Missing   bool   `json:"missing"`
	Revisions []struct {
		Slots struct {
			Main struct {
				Content string `json:"content"`
			} `json:"main"`
		} `json:"slots"`
	} `json:"revisions"`
}

func (ac *apiClient) getByName(ctx context.Context, name string) (*domain.Person, error) {
	requestUrl := fmt.Sprintf(ac.baseApiUrl, name)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", UnexpectedErr, err.Error())
	}

	response, err := ac.client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("%v: %s", UnexpectedErr, err.Error())
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", UnexpectedErr, err.Error())
	}

	var wikiData apiData
	err = json.Unmarshal(data, &wikiData)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", JsonDecodingErr, err.Error())
	}

	page := wikiData.Query.Pages[0]
	if page.Missing {
		return nil, NotFoundErr
	}

	return parseWikiData(&page), nil
}

func parseWikiData(data *pageData) *domain.Person {
	return &domain.Person{
		Name:             data.Title,
		ShortDescription: getShortDescription(data.Revisions[0].Slots.Main.Content),
	}
}

func getShortDescription(content string) string {
	position := strings.Index(content, shortDescriptionTag)
	if position == -1 {
		return ""
	}

	contentRunes := []rune(content)

	runePointer := position + shortDescriptionSkipRunes
	startShortDescription := runePointer

	for ; runePointer < len(contentRunes); runePointer = runePointer + 1 {
		if contentRunes[runePointer] == shortDescriptionEndRune {
			break
		}
	}

	if runePointer == startShortDescription {
		return ""
	}

	return string(contentRunes[startShortDescription:runePointer])
}
