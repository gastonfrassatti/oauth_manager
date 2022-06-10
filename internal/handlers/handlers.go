package handlers

import "github.com/gin-gonic/gin"

type Handler interface {
	ManageGrants(ctx *gin.Context)
}
