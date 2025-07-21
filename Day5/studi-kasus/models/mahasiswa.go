package models

type Alamat struct {
	Jalan   string `json:"jalan"`
	Kota    string `json:"kota"`
	KodePos string `json:"kode_pos"`
}

type Mahasiswa struct {
	ID              int      `json:"id"`
	Nama            string   `json:"nama"`
	Umur            int      `json:"umur"`
	Hobi            []string `json:"hobi"`
	Alamat          Alamat   `json:"alamat"`
	NoHP            *string  `json:"no_hp"`
	Nilai           []int    `json:"nilai"`
	Lulus           bool     `json:"lulus"`
	KeteranganLulus string   `json:"keterangan_lulus"`
}
