package schemas

type CreateMahasiswaRequest struct {
	Nama   string        `json:"nama" binding:"required"`
	NIM    string        `json:"nim" binding:"required"`
	Umur   int           `json:"umur" binding:"required,min=11,max=30,numeric"`
	Hobi   []string      `json:"hobi"`
	NoHP   *string       `json:"no_hp"`
	Nilai  []int         `json:"nilai"`
	Alamat AlamatRequest `json:"alamat" binding:"required"`
}

type AlamatRequest struct {
	Alamat  string `json:"alamat" binding:"required"`
	Kota    string `json:"kota" binding:"required"`
	KodePos string `json:"kode_pos" binding:"required"`
}

type ListMahasiswaRequest struct {
	Search *string `form:"search"`
	Lulus  *bool   `form:"lulus"`
}

type InputNilaiMahasiswaRequest struct {
	Nilai        int  `json:"nilai" binding:"required,min=0,max=100,numeric"`
	MahasiswaID  uint `json:"mahasiswa_id" binding:"required"`
	MataKuliahID uint `json:"mata_kuliah_id" binding:"required"`
}

type DosenRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type MataKuliahRequest struct {
	Nama    string `json:"nama" binding:"required"`
	DosenID uint   `json:"dosen_id" binding:"required"`
}
