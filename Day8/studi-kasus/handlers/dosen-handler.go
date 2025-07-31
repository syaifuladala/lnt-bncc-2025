package handlers

import (
	"mahasiswa/models"
	"mahasiswa/repositories"
	"mahasiswa/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListDosen(c *gin.Context) {
	dosen, err := repositories.ListDosen()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dosen)
}

func GetDosenByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	dosen, err := repositories.GetDosenByID(uint(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dosen)
}

func CreateDosen(c *gin.Context) {
	var request schemas.DosenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dosen := models.Dosen{
		Nama: request.Nama,
	}

	err := repositories.CreateDosen(dosen)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, dosen)
}

func UpdateDosen(c *gin.Context) {
	var request schemas.DosenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	dosen, err := repositories.GetDosenByID(uint(idInt))
	if err != nil {
		c.JSON(404, gin.H{"error": "Dosen not found"})
		return
	}

	dosen.Nama = request.Nama

	err = repositories.UpdateDosen(dosen.ID, dosen)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dosen)
}

func DeleteDosen(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	err = repositories.DeleteDosen(uint(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Dosen deleted successfully"})
}
