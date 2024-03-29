package heroku

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Dyno struct {
	// app
	App struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"app"`
	// a URL to stream output from for attached processes or null for non-attached processes
	AttachUrl *string `json:"attach_url"`
	// command used to start this process
	Command string `json:"command"`
	// when dyno was created
	CreatedAt time.Time `json:"created_at"`
	// unique identifier of this dyno
	Id uuid.UUID `json:"id"`
	// the name of this process on this dyno
	Name string `json:"name"`
	// identity of release
	Release struct {
		Id      uuid.UUID `json:"id"`
		Version int       `json:"version"`
	} `json:"release"`
	// dyno size (default: “standard-1X”)
	Size string `json:"size"`
	// current status of process (either: crashed, down, idle, starting, or up)
	State string `json:"state"`
	// type of process
	Type string `json:"type"`
	// when process last changed state
	UpdatedAt time.Time `json:"updated_at"`
}

func (service *Service) ListDynos(appIdOrName string) (*[]Dyno, *errortools.Error) {
	dynos := []Dyno{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("apps/%s/dynos", appIdOrName)),
		ResponseModel: &dynos,
	}

	_, _, e := service.httpRequest(&requestConfig)

	return &dynos, e
}
