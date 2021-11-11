package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitiateHerokuDB() *gorm.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		"ec2-18-211-243-247.compute-1.amazonaws.com",
		"5432",
		"jkackocdmusjnf",
		"807dceca2ad5749d35ce9948ed0c857c6b50e0bd2050fb604ca10348f9923ffa",
		"ddoqb9eru7odoe",
		// os.Getenv("DATABASE_HOSTNAME"),
		// os.Getenv("DATABASE_PORT"),
		// os.Getenv("DATABASE_USERNAME"),
		// os.Getenv("DATABASE_PASSWORD"),
		// os.Getenv("DATABASE_DBNAME_STAGING"),
	)
	connection, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		return nil
	}

	db, _ := connection.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(5)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(10)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(1 * time.Minute)

	return connection
}
