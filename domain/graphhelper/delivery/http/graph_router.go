package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main/domain/graphhelper/repository"
	"main/domain/graphhelper/usecase"
)

func registerGraphRoutes(router *gin.Engine, g *GraphHelper) {
	route := "loginAzure"
	routes := router.Group(route)
	routes.POST("/device", g.DeviceCodeLogin)
	routes.POST("/password", g.AccountPasswordLogin)
}

func Injection(_ *gorm.DB, router *gin.Engine) {
	graphHelper := repository.NewGraphHelper()
	usecase := usecase.NewGraphHelper(graphHelper)
	handler := NewGraphHandler(usecase)
	registerGraphRoutes(router, handler)
}
