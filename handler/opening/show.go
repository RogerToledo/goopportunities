package opening

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/me/goopportunities/config"
	"github.com/me/goopportunities/handler"
	"github.com/me/goopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a new opening
// @Tags opening
// @Accept  json
// @Produce  json
// @Param id query string true "Opening ID"
// @Success 200 {object} handler.OpeningResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /opening [get]
func ShowOpening(ctx *gin.Context) {
	db := config.GetSQLite()
	opening := schemas.Opening{}
	id := ctx.Query("id")

	if err := db.First(&opening, id).Error; err != nil {
		if strings.Contains(err.Error(), "not found") {
			handler.SendError(ctx, http.StatusNotFound, "opening not found")
			return
		} else {
			handler.SendError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	handler.SendSuccess(ctx, "opening found", opening)
}
