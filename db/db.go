package db

import (
	"fmt"
	"pb4/config"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
)

const dbErrorMessage = "Error connecting to database"

var db *gorm.DB

func Initdatabase() {

	migrateConnection, err := migrate.New("file://db/migrate", config.GetConfig().Database.URL)

	if err != nil {
		fmt.Println("Error creating tables")
		return
	}

	version := config.GetConfig().Database.Version
	currentVersion, _, _ := migrateConnection.Version()

	if version != currentVersion {
		err = migrateConnection.Migrate(version)
		if err != nil {
			fmt.Println("Error creating tables")
			return
		}
	}

	migrateConnection.Close()

	db, err = gorm.Open("postgres", config.GetConfig().Database.URL)

	if err != nil {
		fmt.Println(dbErrorMessage)
	}

	db.LogMode(config.GetConfig().Database.LogMode)

	fmt.Println(db)
}

func GetDB() *gorm.DB {
	return db
}
