package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type grantsService interface {
	ManageGrants(oauthKeys OauthKeys) Grant
}

type Grant struct {
	Uuid        string
	AccessToken string
	ExpiresDate string
	TokenType   string
}

type OauthKeys struct {
	Uuid   string `json:"client_uuid" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}

type handler struct {
	service grantsService
}

func NewHandler(service grantsService) *handler {
	return &handler{
		service: service,
	}
}

func (h handler) GrantsHandler(ctx *gin.Context) {
	var oauthkeys OauthKeys

	if err := ctx.ShouldBindJSON(&oauthkeys); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		//TODO Best way to return error?
		return
	}

	grants := h.service.ManageGrants(oauthkeys)
	if grants == (Grant{}) {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, grants)
}
