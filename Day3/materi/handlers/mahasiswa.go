package handlers

import "github.com/gin-gonic/gin"

func ListMahasiswa(c *gin.Context) {
	c.String(200, "Menampilkan list mahasiswa")
}

func CreateMahasiswa(c *gin.Context) {
	c.String(200, "Create data mahasiswa baru")
}
