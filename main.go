package main

import (
	"Assignment-3/controller"
	"Assignment-3/database"
	"Assignment-3/routes"
)

func main() {
	// Inisialisasi koneksi database
	database.StartDB()

	// Mulai goroutine untuk mengupdate sensor data setiap 15 detik
	go controller.UpdateSensorData()

	// Inisialisasi dan jalankan server API
	r := routes.SetupRoutes()
	r.Run(":8080")
}
