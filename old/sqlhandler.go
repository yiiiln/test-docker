package old

import (
	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Database *gorm.DB

func Connect() {
	var err error
	// 彰化："host=localhost port=5432 dbname=postgres user=postgres password=0000 sslmode=prefer connect_timeout=10"
	// 文心: "host=192.168.1.62 port=5432 dbname=Test0711 user=hyiiilnDB password=Gi012345 sslmode=prefer connect_timeout=10"
	// test cloudSQL: "host=107.167.181.162 port=5432 dbname=yiiiln_test_database user=hyiiiln password=Gi012345 sslmode=prefer connect_timeout=10"

	//host := "192.168.1.62"     // DB_HOS
	//port := "5432"             // DB_PORT
	//databaseName := "Test0711" // DB_NAME
	//username := "hyiiilnDB"    // DB_USER
	//password := "Gi012345"     // DB_PASSWORD
	//mode := "prefer"           // DB_sslmode
	//connectTimeout := "300"    // DB_connect_timeout
	//dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s connect_timeout=%s", host, port, databaseName, username, password, mode, connectTimeout)

	dsn := "host=107.167.181.162 port=5432 dbname=yiiiln_test_database user=hyiiiln password=Gi012345 sslmode=prefer connect_timeout=10"

	Database, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Println("connected to the database error message: ", err.Error())
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

func ConnectCloudSQL() {
	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN()) // "cloudsql-postgres" 可自行設定
	if err != nil {
		// ... handle error
		log.Println("connected to the database error message: ", err.Error())
		panic(err)
	}
	// call cleanup when you're done with the database connection
	defer cleanup()

	Database, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsql-postgres", // "cloudsql-postgres" 可自行設定
		DSN:        "host=total-ensign-373607:asia-east1:ga-erp user=hyiiiln password=Gi012345 dbname=yiiiln_test_database sslmode=disable",
	}), &gorm.Config{})

	// ... etc
}
