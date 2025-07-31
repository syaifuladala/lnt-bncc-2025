package seeders

import (
	"errors"
	"fmt"
	"mahasiswa/configs"
	"mahasiswa/models"

	"gorm.io/gorm"
)

func SeederMataKuliah() {
	err := configs.DB.First(&models.MataKuliah{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var mataKuliah []models.MataKuliah

		var dosenBambang models.Dosen
		err := configs.DB.Where("nama", "Bambang").First(&dosenBambang).Error
		if err != nil {
			fmt.Println("error:", err.Error())
			return
		}

		mataKuliah = append(mataKuliah,
			models.MataKuliah{
				Nama:    "Algoritma Pemrograman",
				DosenID: dosenBambang.ID,
			},
			models.MataKuliah{
				Nama:    "Pemrograman Web",
				DosenID: dosenBambang.ID,
			})

		err = configs.DB.Create(&mataKuliah).Error
		if err != nil {
			fmt.Println("error:", err.Error())
		}
	}
}
