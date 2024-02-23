package opening

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/me/goopportunities/config"
	"github.com/me/goopportunities/handler"
	"github.com/me/goopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new opening
// @Tags opening
// @Accept  json
// @Produce  json
// @Param id query string true "Opening ID"
// @Success 200 {object} handler.OpeningResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /opening [delete]
func DeleteOpening(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	opening := schemas.Opening{}

	db := config.GetSQLite()

	if err := db.First(&opening, id).Error; err != nil {
		msg := fmt.Sprintf("error finding opening: %v", err.Error())
		handler.SendError(ctx, http.StatusNotFound, msg)
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		msg := fmt.Sprintf("error deleting opening: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, msg)
		return
	}

	handler.SendSuccess(ctx, "deleted opening", opening)
}
