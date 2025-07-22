package handlers

import (
	"mahasiswa/models"
	"mahasiswa/repositories"
	"mahasiswa/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListMataKuliah(c *gin.Context) {
	mataKuliah, err := repositories.ListMataKuliah()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mataKuliah)
}

func GetMataKuliahByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mataKuliah, err := repositories.GetMataKuliahByID(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mataKuliah)
}

func CreateMataKuliah(c *gin.Context) {
	var request schemas.MataKuliahRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dosen, err := repositories.GetDosenByID(request.DosenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dosen not found"})
		return
	}

	mataKuliah := models.MataKuliah{
		Nama:    request.Nama,
		DosenID: dosen.ID,
	}

	err = repositories.CreateMataKuliah(mataKuliah)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mataKuliah)
}

func UpdateMataKuliah(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request schemas.MataKuliahRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mataKuliah, err := repositories.GetMataKuliahByID(uint(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mata Kuliah not found"})
		return
	}

	dosen, err := repositories.GetDosenByID(request.DosenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dosen not found"})
		return
	}

	mataKuliah.Nama = request.Nama
	mataKuliah.DosenID = dosen.ID

	err = repositories.UpdateMataKuliah(mataKuliah.ID, mataKuliah)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mataKuliah)
}

func DeleteMataKuliah(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = repositories.DeleteMataKuliah(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
