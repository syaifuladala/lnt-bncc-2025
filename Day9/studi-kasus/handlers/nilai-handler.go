package handlers

import (
	"mahasiswa/models"
	"mahasiswa/repositories"
	"mahasiswa/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InputNilaiMahasiswa(c *gin.Context) {
	var request schemas.InputNilaiMahasiswaRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var nilai models.NilaiMahasiswa

	nilai.Nilai = request.Nilai

	mhs, err := repositories.GetMahasiswaByID(request.MahasiswaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nilai.MahasiswaID = mhs.ID
	nilai.MataKuliahID = request.MataKuliahID

	err = repositories.InputNilaiMahasiswa(nilai)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nilai})
}
