package models

type Alamat struct {
	Jalan   string
	Kota    string
	KodePos string
}

type Mahasiswa struct {
	Nama            string
	Umur            int
	Hobi            []string
	Alamat          Alamat
	NoHP            *string
	Nilai           []int
	Lulus           bool
	KeteranganLulus string
}
