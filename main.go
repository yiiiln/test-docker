package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
	_ "main/docs"
	user "main/domain/crud/delivery/http"
	graph "main/domain/graphhelper/delivery/http"
	test "main/domain/test_connect_cloud_run/delivery/http"
	"main/helper/repository/postgresql"
)

// @title test-connect-api
// @description test-connect.
// @version 1.0.0
//
// @host localhost:8080
// schemes http
//
// @securityDefinitions.oauth2.implicit BearerOAuth2
// @authorizationUrl
// @flow implicit
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

	loadEnv()

	router := gin.Default()
	postgres := postgresql.InitDB()
	//postgresql.MigrateDB(postgres)
	genDependencyInjection(postgres, router)

	swaggerURL := fmt.Sprintf("http://%s:%s/swagger/doc.json", "localhost", "8080")
	url := ginSwagger.URL(swaggerURL)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":8080")
}
func loadEnv() {
	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func genDependencyInjection(db *gorm.DB, router *gin.Engine) {
	user.Injection(db, router)
	test.Injection(db, router)
	graph.Injection(nil, router)
}
