package handlers

import (
	"bytes"
	"fmt"
	"mahasiswa/models"
	"mahasiswa/repositories"
	"mahasiswa/schemas"
	"mahasiswa/utils"
	"net/http"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

	if mahasiswa.Photo != nil {
		urlPath := "localhost:8080/api/mahasiswa/view/" + *mahasiswa.Photo
		mahasiswa.Photo = &urlPath
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

func GetMataKuliahMahasiswaByID(c *gin.Context) {
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

	mataKuliah, err := repositories.GetMataKuliahMahasiswaByID(mahasiswa.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response schemas.ListMataKuliahMahasiswaResponse
	response.ID = mahasiswa.ID
	response.Nama = mahasiswa.Nama
	response.NIM = mahasiswa.NIM
	for _, mk := range mataKuliah {
		response.MataKuliah = append(response.MataKuliah, schemas.ListMataKuliahResponse{
			ID:        mk.MataKuliah.ID,
			Nama:      mk.MataKuliah.Nama,
			NamaDosen: mk.MataKuliah.Dosen.Nama,
			Nilai:     mk.Nilai,
		})
	}

	c.JSON(http.StatusOK, response)
}

func UpdatePhotoMahasiswa(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allowedExt := []string{".jpg", ".png"}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !slices.Contains(allowedExt, ext) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid extension"})
		return
	}

	id := c.Param("id")
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

	uploadDir := "./uploads/"
	saveDirectory := fmt.Sprintf("%s%s_%s", uploadDir, mahasiswa.NIM, file.Filename)

	if err := c.SaveUploadedFile(file, saveDirectory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mahasiswa.Photo = &file.Filename

	err = repositories.UpdateMahasiswa(mahasiswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": file.Filename, "ext": ext})
}

func ViewPhotoMahasiswa(c *gin.Context) {
	filename := c.Param("filename")
	filePath := fmt.Sprintf("./uploads/%s", filename)

	c.File(filePath)
}

func DownloadPhotoMahasiswa(c *gin.Context) {
	filename := c.Param("filename")
	filePath := fmt.Sprintf("./uploads/%s", filename)

	c.FileAttachment(filePath, filename)
}

func ExportMahasiswa(c *gin.Context) {
	f := excelize.NewFile()
	sheetName := "Sheet1"

	f.SetCellValue(sheetName, "A1", "Nama")
	f.SetCellValue(sheetName, "B1", "NIM")

	mahasiswas, err := repositories.ListMahasiswa()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for index, mhs := range mahasiswas {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", index+2), mhs.Nama)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", index+2), mhs.NIM)
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=users.xlsx")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
}
