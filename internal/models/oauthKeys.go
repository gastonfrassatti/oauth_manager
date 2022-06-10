package models

type OauthKeys struct {
	Uuid   string `json:"client_uuid" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}
