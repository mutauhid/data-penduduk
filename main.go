package main

import (
	"data-penduduk/controllers"
	"data-penduduk/database"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}
	psqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	fmt.Println(psqlinfo)

	DB, err = sql.Open("postgres", psqlinfo)
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	} else {
		fmt.Println("DB Connection success")
	}

	database.DbMigrate(DB)
	controllers.Initialize(DB)
	defer DB.Close()

	router := gin.Default()

	// router province
	router.GET("/province", controllers.GetProvince)
	router.POST("/province", controllers.CreateProvince)
	router.PUT("/province/:id", controllers.UpdateProvince)
	router.DELETE("/province/:id", controllers.DeleteProvince)

	//router regency
	router.GET("/regency", controllers.GetRegency)
	router.POST("/regency", controllers.CreateRegency)
	router.PUT("/regency/:id", controllers.UpdateRegency)
	router.DELETE("/regency/:id", controllers.DeleteRegency)

	//router district
	router.GET("/district", controllers.GetDistrict)
	router.POST("/district", controllers.CreateDistrict)
	router.PUT("/district/:id", controllers.UpdateDistrict)
	router.DELETE("/district/:id", controllers.DeleteDistrict)

	//router people
	router.GET("/people", controllers.GetPeople)
	router.POST("/people", controllers.CreatePeople)
	router.PUT("/people/:id", controllers.UpdatePeople)
	router.DELETE("/people/:id", controllers.DeletePeople)

	fmt.Println("server running at http://localhost:8080")
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}
