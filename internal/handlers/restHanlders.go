package handlers

import (
	"gaston.frassatti/aouth_manager/internal/models"
	"gaston.frassatti/aouth_manager/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return handler{
		service: service,
	}
}

func (h handler) ManageGrants(ctx *gin.Context) {
	var oauthkeys models.OauthKeys

	if err := ctx.ShouldBindJSON(&oauthkeys); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		//TODO Best way to return error?
		return
	}

	grants := h.service.ManageGrants(oauthkeys)
	if grants == (models.Grant{}) {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, grants)
}
