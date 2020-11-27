package heroku

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

const (
	apiName string = "Heroku"
	apiURL  string = "https://api.heroku.com"
)

type Heroku struct {
	bearerToken string
	Client      *http.Client
}

type HerokuError struct {
	Resource string `json:"resource"`
	ID       string `json:"id"`
	Message  string `json:"message"`
}

func NewHeroku(bearerToken string) (*Heroku, error) {
	h := Heroku{}
	h.bearerToken = bearerToken
	h.Client = &http.Client{}

	return &h, nil
}

func (_ *Heroku) baseURL() string {
	return apiURL
}

func (h *Heroku) httpRequest(httpMethod string, url string, body io.Reader, model interface{}) *errortools.Error {
	e := new(errortools.Error)

	request, err := http.NewRequest(httpMethod, url, body)
	e.SetRequest(request)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	// Add authorization token to header
	bearer := fmt.Sprintf("Bearer %s", h.bearerToken)
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := utilities.DoWithRetry(h.Client, request, 10, 3)
	e.SetResponse(response)

	if err != nil {
		if response != nil {
			defer response.Body.Close()

			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				e.SetMessage(err)
				return e
			}

			herokuError := HerokuError{}

			err = json.Unmarshal(b, &herokuError)
			if err != nil {
				e.SetMessage(err)
				return e
			}

			e.SetMessage(herokuError.Message)
			return e
		}

		e.SetMessage(err)
		return e
	}

	if model != nil {
		defer response.Body.Close()

		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			e.SetMessage(err)
			return e
		}

		err = json.Unmarshal(b, &model)
		if err != nil {
			e.SetMessage(err)
			return e
		}
	}

	return nil
}

func (h *Heroku) get(url string, model interface{}) *errortools.Error {
	return h.httpRequest(http.MethodGet, url, nil, model)
}

func (h *Heroku) post(url string, buf *bytes.Buffer, model interface{}) *errortools.Error {
	return h.httpRequest(http.MethodPost, url, buf, model)
}
