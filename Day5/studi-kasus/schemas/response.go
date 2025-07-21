package schemas

type ListMahasiswaResponse struct {
	ID              int     `json:"id"`
	Nama            string  `json:"nama"`
	RataRataNilai   float64 `json:"rata_rata_nilai"`
	KeteranganLulus string  `json:"keterangan_lulus"`
}
