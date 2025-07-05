package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/abhiii71/geo-service/geoservice"
	"github.com/gin-gonic/gin"
)

var kdTree *geoservice.KDNode

func main() {
	// Load seed data
	data, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal("Failed to read data.json:", err)
	}

	var locations []geoservice.Location
	if err := json.Unmarshal(data, &locations); err != nil {
		log.Fatal("Invalid data format:", err)
	}

	for _, loc := range locations {
		kdTree = geoservice.Insert(kdTree, loc, 0)
	}

	// Init Gin
	r := gin.Default()
	r.GET("/get-nearby", func(c *gin.Context) {
		geoservice.NearbyHandler(c, kdTree)
	})

	r.Run(":8080")
}
