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
			mahasiswaRoutes.GET("/:id/mata-kuliah", handlers.GetMataKuliahMahasiswaByID)
			mahasiswaRoutes.POST("/", handlers.CreateMahasiswa)
			mahasiswaRoutes.PUT("/:id", handlers.UpdateMahasiswa)
			mahasiswaRoutes.DELETE("/:id", handlers.DeleteMahasiswa)
		}

		dosenRoutes := apiRoutes.Group("/dosen")
		{
			dosenRoutes.GET("/", handlers.ListDosen)
			dosenRoutes.GET("/:id", handlers.GetDosenByID)
			dosenRoutes.POST("/", handlers.CreateDosen)
			dosenRoutes.PUT("/:id", handlers.UpdateDosen)
			dosenRoutes.DELETE("/:id", handlers.DeleteDosen)
		}

		mataKuliahRoutes := apiRoutes.Group("/mata-kuliah")
		{
			mataKuliahRoutes.GET("/", handlers.ListMataKuliah)
			mataKuliahRoutes.GET("/:id", handlers.GetMataKuliahByID)
			mataKuliahRoutes.POST("/", handlers.CreateMataKuliah)
			mataKuliahRoutes.PUT("/:id", handlers.UpdateMataKuliah)
			mataKuliahRoutes.DELETE("/:id", handlers.DeleteMataKuliah)
		}

		nilaiRoutes := apiRoutes.Group("/nilai")
		{
			nilaiRoutes.POST("", handlers.InputNilaiMahasiswa)
		}
	}
}
