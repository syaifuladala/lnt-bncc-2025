package routes

import (
	"mahasiswa/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	apiRoutes := r.Group("/api")
	{
		mahasiswaRoutes := apiRoutes.Group("/mahasiswa")
		{
			mahasiswaRoutes.GET("/", handlers.ListMahasiswa)
			mahasiswaRoutes.GET("/:id", handlers.GetMahasiswaByID)
			mahasiswaRoutes.POST("/", handlers.CreateMahasiswa)
			mahasiswaRoutes.PUT("/:id", handlers.UpdateMahasiswa)
			mahasiswaRoutes.DELETE("/:id", handlers.DeleteMahasiswa)
		}
	}
}
