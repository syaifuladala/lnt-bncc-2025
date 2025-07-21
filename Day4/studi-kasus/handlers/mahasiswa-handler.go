package handlers

import (
	"mahasiswa/models"
	"mahasiswa/schemas"
	"mahasiswa/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var mahasiswas []models.Mahasiswa

func ListMahasiswa(c *gin.Context) {
	var request schemas.ListMahasiswaRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []schemas.ListMahasiswaResponse
	for _, mhs := range mahasiswas {
		rataRata := utils.RataRataNilai(mhs.Nilai)
		_, keteranganLulus := utils.CekStatusKelulusan(rataRata)

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
			RataRataNilai:   rataRata,
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
	lulus, keteranganLulus := utils.CekStatusKelulusan(rata2)

	id := len(mahasiswas) + 1

	newMahasiswa := models.Mahasiswa{
		ID:              id,
		Nama:            req.Nama,
		Umur:            req.Umur,
		Hobi:            req.Hobi,
		NoHP:            req.NoHP,
		Alamat:          models.Alamat{Jalan: req.Alamat.Jalan, Kota: req.Alamat.Kota, KodePos: req.Alamat.KodePos},
		Nilai:           req.Nilai,
		Lulus:           lulus,
		KeteranganLulus: keteranganLulus,
	}

	mahasiswas = append(mahasiswas, newMahasiswa)

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

	for _, mhs := range mahasiswas {
		if mhs.ID == idInt {
			c.JSON(http.StatusOK, mhs)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Mahasiswa tidak ditemukan"})
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
	lulus, keteranganLulus := utils.CekStatusKelulusan(rata2)

	for i, mhs := range mahasiswas {
		if mhs.ID == idInt {
			mhs.Nama = req.Nama
			mhs.Umur = req.Umur
			mhs.Hobi = req.Hobi
			mhs.NoHP = req.NoHP
			mhs.Alamat = models.Alamat{Jalan: req.Alamat.Jalan, Kota: req.Alamat.Kota, KodePos: req.Alamat.KodePos}
			mhs.Nilai = req.Nilai
			mhs.Lulus = lulus
			mhs.KeteranganLulus = keteranganLulus

			mahasiswas[i] = mhs

			c.JSON(http.StatusOK, mhs)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Mahasiswa tidak ditemukan"})
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

	for i, mhs := range mahasiswas {
		if mhs.ID == idInt {
			before := mahasiswas[:i]  // Elemen sebelum i
			after := mahasiswas[i+1:] // Elemen setelah i
			mahasiswas = append(before, after...)
			c.JSON(http.StatusOK, nil)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Mahasiswa tidak ditemukan"})
}
