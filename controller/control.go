package controller

import (
	"fmt"
	"math/rand"
	"time"

	"Assignment-3/database"
	"Assignment-3/models"

	"github.com/gin-gonic/gin"
)

var sensorData models.SensorData

func UpdateSensorData() {
	for {
		newWater := rand.Intn(100) + 1
		newWind := rand.Intn(100) + 1
		waterStatus := getStatus(newWater, "water")
		windStatus := getStatus(newWind, "wind")

		// Simpan data ke database
		sensorData := models.SensorData{
			Water:  newWater,
			Wind:   newWind,
			Status: fmt.Sprintf("Water: %s, Wind: %s", waterStatus, windStatus),
		}
		database.GetDB().Create(&sensorData)

		// Output log status
		fmt.Printf("Updated Sensor Data:\nWater: %d meters\nWind: %d meters/second\nStatus Water: %s\nStatus Wind: %s\n", newWater, newWind, waterStatus, windStatus)

		time.Sleep(15 * time.Second)
	}
}

func getStatus(value int, sensorType string) string {
	if sensorType == "water" {
		switch {
		case value < 5:
			return "Aman"
		case value >= 5 && value <= 8:
			return "Siaga"
		default:
			return "Bahaya"
		}
	} else if sensorType == "wind" {
		switch {
		case value < 6:
			return "Aman"
		case value >= 6 && value <= 15:
			return "Siaga"
		default:
			return "Bahaya"
		}
	}
	return ""
}

func GetSensorData(c *gin.Context) {
	// Ambil data sensor dari database
	var sensorData models.SensorData
	result := database.GetDB().Last(&sensorData)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch sensor data"})
		return
	}

	c.JSON(200, sensorData)
}
