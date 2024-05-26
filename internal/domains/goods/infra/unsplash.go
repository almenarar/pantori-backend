package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type UnsplashResponse struct {
	Results []struct {
		Urls struct {
			Regular string `json:"regular"`
		} `json:"urls"`
	} `json:"results"`
}

type TranlateResponse struct {
	Sentences []struct {
		Trans string `json:"trans"`
	} `json:"sentences"`
}

type UnsplashParams struct {
	BaseURL   string
	AccessKey string
}

type unsplash struct {
	params UnsplashParams
}

func NewUnsplash(params UnsplashParams) *unsplash {
	return &unsplash{
		params: params,
	}
}

func (uns *unsplash) GetImageURL(description string) string {
	translated, err := uns.translate(description)
	if err != nil {
		log.Error().Stack().Msg(err.Error())
		translated = description
	}

	baseURL := uns.params.BaseURL
	queryParams := url.Values{}
	queryParams.Set("query", translated)
	queryParams.Set("client_id", uns.params.AccessKey)
	queryParams.Set("orientation", "landscape")

	requestURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
	response, err := http.Get(requestURL)
	if err != nil {
		log.Error().Stack().Msg(err.Error())
		return ""
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Stack().Msg(err.Error())
		return ""
	}

	var unsplashResponse UnsplashResponse
	err = json.Unmarshal(body, &unsplashResponse)
	if err != nil {
		log.Error().Stack().Msg(err.Error())
		return ""
	}

	if len(unsplashResponse.Results) > 0 {
		return unsplashResponse.Results[0].Urls.Regular
	}

	log.Error().Stack().Msg(fmt.Sprintf("no image found for description: %s", description))
	return ""
}

// TODO receive languages, check if it needs to translate
func (uns *unsplash) translate(description string) (string, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=%s&hl=%s&dt=t&dt=bd&dj=1&source=input&q=%s", "en", "pt", url.QueryEscape(description))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if body == nil {
		return "", errors.New("empty body")
	}

	var data TranlateResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.Sentences[0].Trans, nil
}
