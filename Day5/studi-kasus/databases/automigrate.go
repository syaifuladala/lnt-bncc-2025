package databases

import (
	"fmt"
	"mahasiswa/configs"
	"mahasiswa/models"
)

func AutoMigrate() {
	err := configs.DB.AutoMigrate(
		&models.Mahasiswa{},
	)
	if err != nil {
		errorLog := fmt.Sprintf("Gagal Auto Migrate: %s", err.Error())
		panic(errorLog)
	}
}
