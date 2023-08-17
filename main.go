package main

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	_ "main/docs"
	user "main/domain/crud/delivery/http"
	test "main/domain/test_connect_cloud_run/delivery/http"
	"main/helper/repository/postgresql"

	"github.com/gin-gonic/gin"
)

// @title Gin Swagger Demo
// @version 1.0
// @description Gin Swagger API.
// @contact.name JLHou
// @host localhost:8080
// schemes http
func main() {

	////domain.Connect()
	//old.ConnectCloudSQL()
	//old.Database.AutoMigrate(&old.DbUser{}) // 遷移schema
	//
	//server := gin.Default()
	//server.POST("/c", old.Register)  // 註冊(新增)
	//server.POST("/login", old.Login) // 登入(讀取)
	//server.GET("/r", old.GetUsers)
	//server.PUT("/ua/:ID", old.UpdateAccountByID)
	//server.PUT("/up/:ID", old.UpdatePasswordByID)
	//server.DELETE("/d", old.DeleteUserByAccount)
	//server.GET("/test_connect", old.Test)
	//
	//// http://localhost:8080/swagger/index.html
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	//server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//
	//server.Run(":8080") // localhost:8080

	router := gin.Default()
	postgres := postgresql.InitDB()
	postgresql.MigrateDB(postgres)
	genDependencyInjection(postgres, router)

	swaggerURL := fmt.Sprintf("http://%s:%s/swagger/doc.json", "localhost", "8080")
	url := ginSwagger.URL(swaggerURL)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":8080")
}
func genDependencyInjection(db *gorm.DB, router *gin.Engine) {
	user.Injection(db, router)
	test.Injection(db, router)
}
