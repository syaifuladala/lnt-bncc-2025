package main

import (
	"fmt"
	"mahasiswa/configs"
	"mahasiswa/databases"
	"mahasiswa/databases/seeders"
	"mahasiswa/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Kesalahan:", r)
		}
	}()

	// koneksi ke db
	configs.SetupMySQL()

	// automigrate
	databases.AutoMigrate()

	// seeder
	seeders.SeederDosen()
	seeders.SeederMataKuliah()

	// router
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}
