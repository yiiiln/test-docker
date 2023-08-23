package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerCRUDRoutes(router *gin.Engine) {
	route := "TestConnect"
	routes := router.Group(route)
	routes.POST("/", TestConnect)
}
func Injection(_ *gorm.DB, router *gin.Engine) {
	registerCRUDRoutes(router)
}
