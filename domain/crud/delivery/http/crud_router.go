package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main/domain/crud/repository/postgresql"
	"main/domain/crud/usecase"
)

func registerCRUDRoutes(router *gin.Engine, h *Handler) {
	route := "CRUD"
	routes := router.Group(route)
	routes.POST("/", h.Register)
	routes.GET("/", h.GetAllUser)
	routes.GET("/:id", h.GetUser)
	routes.PUT("/uA/:id", h.UpdateAccountByID)
	routes.PUT("/uP/:id", h.UpdatePasswordByID)
	routes.DELETE("/:id", h.DeleteUserByID)
	routes.POST("/login", h.Login)
}

func Injection(db *gorm.DB, router *gin.Engine) {
	repository := postgresql.NewUserRepositoryPGImpl(db)
	usecase := usecase.NewUserUsecase(repository)
	handler := NewUserHandler(usecase)
	registerCRUDRoutes(router, handler)
}
