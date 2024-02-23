package opening

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/me/goopportunities/config"
	"github.com/me/goopportunities/handler"
	"github.com/me/goopportunities/schemas"
)

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
