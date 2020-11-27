package heroku

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	errortools "github.com/leapforce-libraries/go_errortools"
)

type Apps []App

type App struct {
	// ACM status of this app
	ACM bool `json:"acm"`
	// when app was archived
	ArchivedAt *time.Time `json:"archived_at"`
	// stack
	BuildStack struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"build_stack"`
	// description from buildpack of app
	BuildpackProvidedDescription *string `json:"buildpack_provided_description"`
	// when app was created
	CreatedAt time.Time `json:"created_at"`
	// git repo URL of app
	GitURL string `json:"git_url"`
	// unique identifier
	ID uuid.UUID `json:"id"`
	// describes whether a Private Spaces app is externally routable or not
	InternalRouting *bool `json:"internal_routing"`
	// maintenance status of app
	Maintenance bool `json:"maintenance"`
	// name of app
	Name string `json:"name"`
	// identity of team
	Organization *struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"organization"`
	// identity of an account
	Owner struct {
		Email string    `json:"email"`
		ID    uuid.UUID `json:"id"`
	} `json:"owner"`
	// identity of app region
	Region struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"region"`
	// when app was released
	ReleasedAt *time.Time `json:"released_at"`
	// git repo size in bytes of app
	RepoSize *int `json:"repo_size"`
	// slug size in bytes of app
	SlugSize *int `json:"slug_size"`
	// identity of space
	Space *struct {
		ID     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		Shield bool      `json:"shield"`
	} `json:"space"`
	// identity of app stack
	Stack struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"stack"`
	// identity of team
	Team *struct {
		ID     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		Shield bool      `json:"shield"`
	} `json:"team"`
	// when app was updated
	UpdatedAt time.Time `json:"updated_at"`
	// web URL of app
	WebURL string `json:"web_url"`
}

func (h *Heroku) ListApps() (*[]App, *errortools.Error) {
	url := fmt.Sprintf("%s/%s", h.baseURL(), "apps")

	apps := []App{}
	e := h.get(url, &apps)

	return &apps, e
}
