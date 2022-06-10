package repository

import "gaston.frassatti/aouth_manager/internal/models"

type Dao interface {
	GetGrants(string) (models.Grant, error)
	Upsert(models.Grant)
}
