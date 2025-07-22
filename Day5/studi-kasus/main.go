package main

import (
	"fmt"
	"mahasiswa/configs"
	"mahasiswa/databases"
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

	// router
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}
