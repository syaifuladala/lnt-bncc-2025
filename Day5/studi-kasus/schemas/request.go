package schemas

type CreateMahasiswaRequest struct {
	Nama   string        `json:"nama" binding:"required"`
	Umur   int           `json:"umur" binding:"required,min=11,max=30,numeric"`
	Hobi   []string      `json:"hobi"`
	NoHP   *string       `json:"no_hp"`
	Nilai  []int         `json:"nilai"`
	Alamat AlamatRequest `json:"alamat" binding:"required"`
}

type AlamatRequest struct {
	Jalan   string `json:"jalan" binding:"required"`
	Kota    string `json:"kota" binding:"required"`
	KodePos string `json:"kode_pos" binding:"required"`
}

type ListMahasiswaRequest struct {
	Search *string `form:"search"`
	Lulus  *bool   `form:"lulus"`
}
