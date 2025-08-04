package schemas

type ListMahasiswaResponse struct {
	ID              uint    `json:"id"`
	Nama            string  `json:"nama"`
	RataRataNilai   float64 `json:"rata_rata_nilai"`
	KeteranganLulus string  `json:"keterangan_lulus"`
}

type ListMataKuliahMahasiswaResponse struct {
	ID         uint                     `json:"id"`
	Nama       string                   `json:"nama"`
	NIM        string                   `json:"nim"`
	MataKuliah []ListMataKuliahResponse `json:"mata_kuliah"`
}

type ListMataKuliahResponse struct {
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	NamaDosen string `json:"nama_dosen"`
	Nilai     int    `json:"nilai"`
}
