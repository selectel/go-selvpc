package testutils

import (
	"net/http"

	"github.com/selectel/go-selvpcclient/selvpc"
	"github.com/selectel/go-selvpcclient/selvpc/resell"
)

// NewTestResellV2Client prepares a client for the Resell V2 API tests.
func (testEnv *TestEnv) NewTestResellV2Client() {
	apiVersion := "v2"
	resellClient := &selvpc.ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   testEnv.Server.URL + "/resell/" + apiVersion,
		TokenID:    FakeTokenID,
		UserAgent:  resell.UserAgent,
	}
	testEnv.Client = resellClient
}
