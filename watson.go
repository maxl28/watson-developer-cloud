package watson

import (
	"net/http"
)

// ToDo: Write documentation
type WatsonOptions struct {
	// Shall we authenticate or not?
	Use_unauthenticated bool
	// Credentials used for authentication, assuming that we even use authentication.
	Username, Password, Api_key string
	// Which version of the API should be used.
	Version string
	// Determines the Watson service endpoint we're using.
	// Not yet implemented
	Alchemy string
	// Not yet implemented
	Use_vcap_services bool
}

// ToDo: Write documentation
type Watson struct {
	Options WatsonOptions
}

// ToDo: Write documentation
func New(options WatsonOptions) Watson {
	if len(options.Version) < 1 {
		options.Version = "v1"
	}

	return Watson{options};
}

// Sign a HTTP request with our Watson credentials. Either basic HTTP Auth with username and password or an API key.
func (w *Watson) SignRequest(req *http.Request) *http.Request {
	req.SetBasicAuth(w.Options.Username, w.Options.Password)

	return req
}