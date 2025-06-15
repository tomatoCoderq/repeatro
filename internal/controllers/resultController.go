package controllers

import (
	"net/http"

	"repeatro/internal/schemes"
	"repeatro/internal/services"

	"repeatro/internal/tools"

	"github.com/gin-gonic/gin"
)

type ResultController struct {
	ResultService services.ResultServiceInterface
}

func CreateNewResultController(resultService *services.ResultService) *ResultController {
	return &ResultController{ResultService: resultService}
}

func (rc ResultController) GetStats(ctx *gin.Context) {
	var interval schemes.Interval
	if err := ctx.ShouldBindJSON(&interval); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	
	userId, err := tools.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	meanGrade, err := rc.ResultService.GetMeanGradeOfPeriod(interval.DtStart, interval.DtEnd, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"Mean Grade": meanGrade})
}
