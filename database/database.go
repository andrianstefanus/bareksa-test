package database

import "gorm.io/gorm"

type Database struct {
	HerokuDB *gorm.DB
}

// Databases this var will be used everywhere
var Databases Database

func init() {
	// initiate connection to Heroku DB
	Databases.HerokuDB = InitiateHerokuDB()
}
