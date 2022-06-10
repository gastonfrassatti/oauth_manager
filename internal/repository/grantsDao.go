package repository

import (
	"database/sql"
	"errors"
	"gaston.frassatti/aouth_manager/internal/models"
)

type grantsDao struct {
	db *sql.DB
}

func NewGrantsDAO(db *sql.DB) Dao {
	return &grantsDao{
		db: db,
	}
}

func (dao grantsDao) GetGrants(uuid string) (grants models.Grant, err error) {
	row := dao.db.QueryRow("SELECT * FROM oauth_db.grants WHERE oauth_uuid = ?", uuid)

	err = row.Scan(&grants.Uuid, &grants.AccessToken, &grants.ExpiresDate, &grants.TokenType)
	if errors.Is(err, sql.ErrNoRows) {
		return grants, nil
	}
	return grants, err
}

func (dao grantsDao) Upsert(grants models.Grant) {
	stmt, err := dao.db.Prepare("INSERT INTO oauth_db.grants (oauth_uuid, acces_token, expires_date, token_type) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE acces_token=values( acces_token), expires_date=values(expires_date)")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(grants.Uuid, grants.AccessToken, grants.ExpiresDate, grants.TokenType)
}
