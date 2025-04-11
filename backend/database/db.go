package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	//"fmt"
)

var DB *gorm.DB

func ConnectDB(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
		return err
	}

	

	// DROP TABLES
	/*
	DB.Migrator().DropTable(&models.Listing{})

	*/
	


	/*
	err = DB.AutoMigrate(&models.Listing{})    // check
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	*/

	log.Println("Database connected and migrated successfully!")
	return nil
}



func GetDB() *gorm.DB {
	return DB
}