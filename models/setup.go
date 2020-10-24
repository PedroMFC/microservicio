package models

import (
	//"gorm.io/driver/postgres"
	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
var DB *gorm.DB

/*func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}*/

func ConnectDatabase() {
	db, err := gorm.Open( "postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")


	if err != nil {
		panic("Failed to connect to database!")
	}

	defer db.Close()
	db.AutoMigrate(&Valoracion{})

	DB = db
}
