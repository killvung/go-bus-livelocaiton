package main

import (
	"log"

	"com.killvung/go-demo-bus-live-location/models"
	"com.killvung/go-demo-bus-live-location/routes"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "bus.db")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Vehicle{})

	r := gin.Default()

	routes.InitRoutes(r, db)

	r.Run(":8080")
}
