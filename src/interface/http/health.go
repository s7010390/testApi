package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"service_name": "CSV-AnswerQuestionThinker-service",
		"code":         0,
	})
}
