package main

import (
	"fmt"
	"mahasiswa/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Kesalahan:", r)
		}
	}()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}
