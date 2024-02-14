package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
To ensure a single connection to the database, we employ the
Singleton design pattern, initializing the *gorm.DB object
only once.
*/
var DB *gorm.DB
var once sync.Once

func InicializeDB(dns string) error {
	/*With GORM, it's not necessary to use a defer statement
	to close the connection, as the library handles this
	automatically.*/
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}

	/*Synchronization with sync.Once ensures that the
	initialization of the DB object only happens once.*/
	once.Do(func() {
		DB = db
		errorMigrations := db.AutoMigrate(&PlanContratado{}, &EventoTrabajado{})
		if errorMigrations != nil {
			fmt.Errorf("Error occurred while performing database migrations: %w", errorMigrations)
			return
		}
	})

	return nil
}
