package service

import (
	"gaston.frassatti/aouth_manager/internal/models"
)

type Service interface {
	ManageGrants(models.OauthKeys) models.Grant
}
