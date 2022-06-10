package service

import (
	"database/sql"
	"errors"
	"gaston.frassatti/aouth_manager/internal/models"
	"gaston.frassatti/aouth_manager/internal/repository"
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

const dateLayout string = "2006-01-02 15:04:05"

type grantsService struct {
	dao        repository.Dao
	httpClient *sling.Sling
	//Log
}

type requestBody struct {
	grantType    string `json:"grant_type"`
	clientId     string `json:"client_id"`
	clientSecret string `json:"client_secret"`
}

type successResponseBody struct {
	AccessToken string `json:"access_token"`
	ExpiresDate int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func NewGrantsService(dao repository.Dao, client *sling.Sling) Service {
	return &grantsService{
		dao:        dao,
		httpClient: client,
	}
}

func (s grantsService) ManageGrants(oauthKeys models.OauthKeys) models.Grant {
	const pathUrl string = "/mock/api/authorization"
	grants, err := s.dao.GetGrants(oauthKeys.Uuid)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		//TODO return err and endProcess?
	}

	response := successResponseBody{}

	if (models.Grant{}) == grants || tokenIsExpired(grants) {
		res, err := s.httpClient.Post(pathUrl).
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

		grants.Uuid = oauthKeys.Uuid
		grants.AccessToken = response.AccessToken
		grants.ExpiresDate = secondsToDate(response.ExpiresDate)
		grants.TokenType = response.TokenType

		s.dao.Upsert(grants)
	}
	return grants
}

func tokenIsExpired(grants models.Grant) bool {
	t, _ := time.Parse(dateLayout, grants.ExpiresDate)
	return time.Now().After(t)
}

func secondsToDate(secondsToExpire int) (expiresIn string) {
	const secondsOfDay int = 86400
	daysToExpire := secondsToExpire / secondsOfDay
	return time.Now().Local().AddDate(0, 0, daysToExpire).Format(dateLayout)
}
