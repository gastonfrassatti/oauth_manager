package client

import (
	"gaston.frassatti/aouth_manager/internal/handlers"
	"github.com/dghubble/sling"
	"net/http"
)

type slingClient struct {
	sling *sling.Sling
}

type requestBody struct {
	grantType    string `json:"grant_type"`
	clientId     string `json:"client_id"`
	clientSecret string `json:"client_secret"`
}

type SuccessResponseBody struct {
	AccessToken string `json:"access_token"`
	ExpiresDate int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func NewHttpClient(base string) *slingClient {
	return &slingClient{
		sling: sling.New().Base(base),
	}
}

func (client slingClient) PostAuthorization(path string, oauthKeys handlers.OauthKeys) SuccessResponseBody {
	response := SuccessResponseBody{}
	res, err := client.sling.Post(path).
		BodyForm(requestBody{
			grantType:    "client_credentials",
			clientId:     oauthKeys.Uuid,
			clientSecret: oauthKeys.Secret,
		}).
		ReceiveSuccess(&response)

	if err != nil {
		//TODO Log error
	}
	if res.StatusCode != http.StatusOK {
		//TODO Log error
	}
	return response
}
