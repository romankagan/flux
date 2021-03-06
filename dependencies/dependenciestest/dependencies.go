package dependenciestest

import (
	"io/ioutil"
	"net/http"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/dependencies/filesystem"
	"github.com/influxdata/flux/dependencies/url"
	"github.com/influxdata/flux/mock"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

var StatusOK int = 200

func defaultTestFunction(req *http.Request) *http.Response {
	body := (*req).Body
	// Test request parameters
	return &http.Response{
		StatusCode: StatusOK,
		Status:     "Body generated by test client",

		// Send response to be tested
		Body: ioutil.NopCloser(body),

		// Must be set to non-nil value or it panics
		Header: make(http.Header),
	}
}

func Default() flux.Deps {
	var deps flux.Deps
	deps.Deps.HTTPClient = &http.Client{
		Transport: RoundTripFunc(defaultTestFunction),
	}
	deps.Deps.SecretService = &mock.SecretService{
		"password": "mysecretpassword",
		"token":    "mysecrettoken",
	}
	deps.Deps.FilesystemService = filesystem.SystemFS
	deps.Deps.URLValidator = url.PassValidator{}
	return deps
}
