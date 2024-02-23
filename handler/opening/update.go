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

// @Summary Update opening
// @Description Update a new opening
// @Tags opening
// @Accept  json
// @Produce  json
// @Param id query string true "Opening ID"
// @Param opening body handler.UpsertOpeningRequest true "Request body"
// @Success 200 {object} handler.OpeningResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /opening [put]
func UpdateOpening(ctx *gin.Context) {
	db := config.GetSQLite()
	opening := schemas.Opening{}

	request := handler.UpsertOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.ValidateUpdate(); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	
	id := ctx.Query("id")

	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	if err := db.First(&opening, id).Error; err != nil {
		if strings.Contains(err.Error(), "not found") {
			handler.SendError(ctx, http.StatusNotFound, "opening not found")
			return
		} else {
			handler.SendError(ctx, http.StatusInternalServerError, err.Error())
			return
		}	
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccess(ctx, "update opening", opening)
}
