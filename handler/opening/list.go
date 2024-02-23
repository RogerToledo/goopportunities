package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/me/goopportunities/config"
	"github.com/me/goopportunities/handler"
	"github.com/me/goopportunities/schemas"
)

func ListOpening(ctx *gin.Context) {
	db := config.GetSQLite()
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if len(openings) == 0 {
		handler.SendError(ctx, http.StatusNotFound, "No openings found")
		return
	}

	handler.SendSuccess(ctx, "openings find", openings)

}
