package routes

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"com.killvung/go-demo-bus-live-location/models"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/ingest", func(c *gin.Context) {
		defer db.Close()

		startTime := time.Now()
		resp, err := http.Get("https://retro.umoiq.com/service/publicXMLFeed?command=vehicleLocations&a=ttc")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		var data models.Response
		err = xml.Unmarshal(body, &data)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Uncomment me to ingest data cocurrently with Go routine and wg
		// Note: It's slower than sequential solution due to over head
		// var wg sync.WaitGroup
		// wg.Add(len(data.Vehicles))

		// for _, vehicle := range data.Vehicles {
		// 	go func(v models.Vehicle) {
		// 		defer wg.Done()

		// 		if err := db.Create(&v).Error; err != nil {
		// 			log.Println("Error creating vehicle:", err)
		// 		}
		// 	}(vehicle)
		// }
		// wg.Wait()
		for _, v := range data.Vehicles {
			if err := db.Create(&v).Error; err != nil {
				log.Println("Error creating vehicle:", err)
			}
		}

		elapsedTime := time.Since(startTime)

		c.String(http.StatusOK, "Data ingested successfully! Elapsed time: "+elapsedTime.String())
	})
}
