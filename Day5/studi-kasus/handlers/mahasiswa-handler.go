package handlers

import (
	"mahasiswa/models"
	"mahasiswa/repositories"
	"mahasiswa/schemas"
	"mahasiswa/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListMahasiswa(c *gin.Context) {
	var request schemas.ListMahasiswaRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mahasiswas, err := repositories.ListMahasiswa()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []schemas.ListMahasiswaResponse
	for _, mhs := range mahasiswas {
		_, keteranganLulus := utils.CekStatusKelulusan(mhs.RataRata)

		if request.Search != nil {
			if !strings.Contains(mhs.Nama, *request.Search) {
				continue
			}
		}

		if request.Lulus != nil {
			if mhs.Lulus != *request.Lulus {
				continue
			}
		}

		response = append(response, schemas.ListMahasiswaResponse{
			ID:              mhs.ID,
			Nama:            mhs.Nama,
			RataRataNilai:   mhs.RataRata,
			KeteranganLulus: keteranganLulus,
		})
	}
	c.JSON(200, response)
}

func CreateMahasiswa(c *gin.Context) {
	var req schemas.CreateMahasiswaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := utils.CekNoHP(req.NoHP)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = utils.ValidasiNilai(req.Nilai)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rata2 := utils.RataRataNilai(req.Nilai)
	lulus, _ := utils.CekStatusKelulusan(rata2)

	var hobiString string
	if len(req.Hobi) > 0 {
		hobiString = strings.Join(req.Hobi, ", ")
	}

	newMahasiswa := models.Mahasiswa{
		Nama:     req.Nama,
		NIM:      req.NIM,
		Umur:     req.Umur,
		Hobi:     hobiString,
		NoHP:     req.NoHP,
		Alamat:   req.Alamat.Alamat,
		Kota:     req.Alamat.Kota,
		KodePos:  req.Alamat.KodePos,
		Lulus:    lulus,
		RataRata: rata2,
	}

	err = repositories.CreateMahasiswa(&newMahasiswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newMahasiswa)
}

func GetMahasiswaByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak boleh kosong"})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	mahasiswa, err := repositories.GetMahasiswaByID(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

func UpdateMahasiswa(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak boleh kosong"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}
	var req schemas.CreateMahasiswaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = utils.CekNoHP(req.NoHP)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = utils.ValidasiNilai(req.Nilai)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rata2 := utils.RataRataNilai(req.Nilai)
	lulus, _ := utils.CekStatusKelulusan(rata2)

	mahasiswa, err := repositories.GetMahasiswaByID(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mahasiswa.Nama = req.Nama
	mahasiswa.NIM = req.NIM
	mahasiswa.Umur = req.Umur

	var hobiString string
	if len(req.Hobi) > 0 {
		hobiString = strings.Join(req.Hobi, ", ")
	}
	mahasiswa.Hobi = hobiString

	mahasiswa.NoHP = nil
	if req.NoHP != nil {
		mahasiswa.NoHP = req.NoHP
	}

	mahasiswa.Alamat = req.Alamat.Alamat
	mahasiswa.Kota = req.Alamat.Kota
	mahasiswa.KodePos = req.Alamat.KodePos

	mahasiswa.Lulus = lulus
	mahasiswa.RataRata = rata2

	err = repositories.UpdateMahasiswa(mahasiswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mahasiswa)
}

func DeleteMahasiswa(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak boleh kosong"})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	err = repositories.DeleteMahasiswa(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
