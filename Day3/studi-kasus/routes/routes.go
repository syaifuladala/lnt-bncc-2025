package routes

import (
	"mahasiswa/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/mahasiswa", handlers.ListMahasiswa)
}
