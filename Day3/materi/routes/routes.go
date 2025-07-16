package routes

import (
	"belajar-http-server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/api/mahasiswa", handlers.ListMahasiswa)
	r.POST("/api/mahasiswa", handlers.CreateMahasiswa)
}
