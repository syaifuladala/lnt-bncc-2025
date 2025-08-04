package main

import (
	"fmt"
	"mahasiswa/configs"
	"mahasiswa/databases"
	"mahasiswa/databases/seeders"
	"mahasiswa/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Kesalahan:", r)
		}
	}()

	// validasi env
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("No .env file found: %s", err.Error()))
	}

	// koneksi ke db
	configs.SetupMySQL()

	// automigrate
	databases.AutoMigrate()

	// seeder
	seeders.SeederDosen()
	seeders.SeederMataKuliah()
	seeders.SeederPasswordDosen()

	// router
	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run()
}
