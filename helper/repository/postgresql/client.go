package postgresql

import (
	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"main/model"
)

func InitDB() *gorm.DB {
	var db *gorm.DB
	var err error

	// cloud run
	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN()) // "cloudsql-postgres" 可自行設定
	if err != nil {
		// ... handle error
		log.Println("connected to the database error message: ", err.Error())
		panic(err)
	}
	defer cleanup() // call cleanup when you're done with the database connection
	db, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsql-postgres", // "cloudsql-postgres" 可自行設定
		DSN:        "host=total-ensign-373607:asia-east1:ga-erp user=hyiiiln password=Gi012345 dbname=yiiiln_test_database sslmode=disable TimeZone=Asia/Taipei",
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// local
	//db, err = gorm.Open(postgres.New(postgres.Config{
	//	DSN:                  "host=107.167.181.162 port=5432 dbname=yiiiln_test_database user=hyiiiln password=Gi012345 sslmode=prefer TimeZone=Asia/Taipei",
	//	PreferSimpleProtocol: true,
	//}), &gorm.Config{})

	if err != nil {
		log.Fatalln("connected to the database error message: ", err)
	} else {
		log.Println("Successfully connected to the database")
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
