package heroku

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type AddOn struct {
	// actions
	Actions []struct {
		// identifier of the action to take that is sent via SSO
		Action string `json:"action"`
		// a unique identifier
		Id uuid.UUID `json:"id"`
		// the display text shown in Dashboard
		Label string `json:"label"`
		// if the action requires the user to own the app
		RequiresOwner bool `json:"requires_owner"`
		// absolute URL to use instead of an action
		Url string `json:"url"`
	} `json:"actions"`
	// identity of addon_service
	AddOnService struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"addon_service"`
	// app
	App struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"app"`
	// billed price
	BilledPrice *struct {
		// price in cents per unit of plan
		Cents int `json:"cents"`
		// price is negotiated in a contract outside of monthly add-on billing
		Contract bool `json:"contract"`
		// unit of price for plan
		Unit string `json:"unit"`
	}
	// billing_entity
	BillingEntity struct {
		// unique identifier of the billing entity
		Id uuid.UUID `json:"id"`
		// name of the billing entity
		Name string `json:"name"`
		// type of Object of the billing entity; new types allowed at any time
		Type string `json:"type"`
	} `json:"billing_entity"`
	// config vars exposed to the owning app by this add-on
	ConfigVars []string `json:"config_vars"`
	// when add-on was created
	CreatedAt time.Time `json:"created_at"`
	// unique identifier of add-on
	Id uuid.UUID `json:"id"`
	// globally unique name of the add-on
	Name string `json:"name"`
	// identity of plan
	Plan struct {
		// unique identifier of this plan
		Id uuid.UUID `json:"id"`
		// unique name of this plan
		Name string `json:"name"`
	} `json:"plan"`
	// id of this add-on with its provider
	ProviderId string `json:"provider_id"`
	// state in the add-on’s lifecycle
	State string `json:"state"`
	// when add-on was updated
	UpdatedAt time.Time `json:"updated_at"`
	//URL for logging into web interface of add-on (e.g. a dashboard)
	WebUrl *string `json:"web_url"`
}

func (service *Service) ListAddOns(appIdOrName string) (*[]AddOn, *errortools.Error) {
	addOns := []AddOn{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("apps/%s/addons", appIdOrName)),
		ResponseModel: &addOns,
	}

	_, _, e := service.httpRequest(&requestConfig)

	return &addOns, e
}
