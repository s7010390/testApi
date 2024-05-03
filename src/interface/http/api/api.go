package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/s7010390/testApi/interface/http/api/v1"
)

func AddRoute(engine *gin.Engine) {
	apiRoute := engine.Group("/api")
	v1.AddRoute(apiRoute)
}
