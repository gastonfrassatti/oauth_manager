package client

import "github.com/dghubble/sling"

type httpClientConfig struct {
	Base string //
}

func NewSlingConfig() HttpClient {
	//TODO take values from .env
	return &httpClientConfig{
		Base: "https://94ff2c06e96fe23b7913b17145f3c1db.m.pipedream.net",
	}
}

func (config httpClientConfig) SlingClient() *sling.Sling {
	return sling.New().Base(config.Base)
}
