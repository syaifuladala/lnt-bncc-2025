package seeders

import (
	"errors"
	"mahasiswa/configs"
	"mahasiswa/models"

	"gorm.io/gorm"
)

func SeederDosen() {
	err := configs.DB.First(&models.Dosen{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var dosen []models.Dosen

		dosen = append(dosen,
			models.Dosen{
				Nama: "Budi",
			},
			models.Dosen{
				Nama: "Bambang",
			},
			models.Dosen{
				Nama: "Rina",
			})

		configs.DB.Create(&dosen)
	}
}
