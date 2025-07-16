package main

import (
	"belajar-http-server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// // http server dari net/http go
	// http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Halo dari backend sederhana")
	// })

	// fmt.Println("server berjalan di localhost:8080")
	// http.ListenAndServe(":8080", nil)

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run()
}
