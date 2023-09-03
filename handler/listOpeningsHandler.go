package handler

import (
	"github.com/Robertotheto/golang-api/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list opening", openings)
}
