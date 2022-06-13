package service

import (
	"database/sql"
	"errors"
	"gaston.frassatti/aouth_manager/internal/client"
	"gaston.frassatti/aouth_manager/internal/handlers"
	"time"
)

const dateLayout string = "2006-01-02 15:04:05"

type db interface {
	GetGrants(uuid string) (grants handlers.Grant, err error)
	Upsert(grants handlers.Grant)
}

type httpClient interface {
	PostAuthorization(path string, oauthKeys handlers.OauthKeys) client.SuccessResponseBody
}

type GrantsService struct {
	db         db
	httpClient httpClient
}

func NewGrantsService(db db, client httpClient) *GrantsService {
	return &GrantsService{
		db:         db,
		httpClient: client,
	}
}

func (s GrantsService) ManageGrants(oauthKeys handlers.OauthKeys) handlers.Grant {
	const pathUrl string = "/mock/api/authorization"
	grants, err := s.db.GetGrants(oauthKeys.Uuid)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		//TODO return err and endProcess?
	}

	if (handlers.Grant{}) == grants || tokenIsExpired(grants) {
		response := s.httpClient.PostAuthorization(pathUrl, oauthKeys)
		//TODO Builder
		grants.Uuid = oauthKeys.Uuid
		grants.AccessToken = response.AccessToken
		grants.ExpiresDate = secondsToDate(response.ExpiresDate)
		grants.TokenType = response.TokenType

		s.db.Upsert(grants)
	}
	return grants
}

func tokenIsExpired(grants handlers.Grant) bool {
	t, _ := time.Parse(dateLayout, grants.ExpiresDate)
	return time.Now().After(t)
}

func secondsToDate(secondsToExpire int) (expiresIn string) {
	const secondsOfDay int = 86400
	daysToExpire := secondsToExpire / secondsOfDay
	return time.Now().Local().AddDate(0, 0, daysToExpire).Format(dateLayout)
}
