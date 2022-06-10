package client

import "github.com/dghubble/sling"

type HttpClient interface {
	SlingClient() *sling.Sling
}
