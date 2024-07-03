package controllers

import "database/sql"

var db *sql.DB

func Initialize(database *sql.DB) {
	db = database
}
