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

// @Summary Create opening
// @Description Create a new opening
// @Tags opening
// @Accept  json
// @Produce  json
// @Param opening body handler.UpsertOpeningRequest true "Request body"
// @Success 200 {object} handler.OpeningResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /opening [post]
func CreateOpening(ctx *gin.Context) {
	request := handler.UpsertOpeningRequest{}

	ctx.BindJSON(&request)

	fields := request.ValidateCreate()
	if len(fields) > 0 {
		msg := handler.ErrParamRequired(fields)
		handler.SendError(ctx, http.StatusBadRequest, msg)
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	db := config.GetSQLite()

	if err := db.Create(&opening).Error; err != nil {
		msg := fmt.Sprintf("error creating opening: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, msg)
		return
	}

	handler.SendSuccess(ctx, "create opening", opening)

}
