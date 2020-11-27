package heroku

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	errortools "github.com/leapforce-libraries/go_errortools"
)

type Dyno struct {
	// app
	App struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"app"`
	// a URL to stream output from for attached processes or null for non-attached processes
	AttachURL *string `json:"attach_url"`
	// command used to start this process
	Command string `json:"command"`
	// when dyno was created
	CreatedAt time.Time `json:"created_at"`
	// unique identifier of this dyno
	ID uuid.UUID `json:"id"`
	// the name of this process on this dyno
	Name string `json:"name"`
	// identity of release
	Release struct {
		ID      uuid.UUID `json:"id"`
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

func (h *Heroku) ListDynos(appIDOrName string) (*[]Dyno, *errortools.Error) {
	url := fmt.Sprintf("%s/apps/%s/dynos", h.baseURL(), appIDOrName)

	dynos := []Dyno{}
	e := h.get(url, &dynos)

	return &dynos, e
}
