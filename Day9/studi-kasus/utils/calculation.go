package utils

func RataRataNilai(nilai []int) float64 {
	if len(nilai) == 0 {
		return 0
	}
	total := 0
	for _, n := range nilai {
		total += n
	}

	return float64(total) / float64(len(nilai))
}

func CekStatusKelulusan(rata2 float64) (bool, string) {
	switch {
	case rata2 >= 85:
		return true, "Lulus dengan pujian"
	case rata2 >= 70:
		return true, "Lulus"
	case rata2 >= 50:
		return false, "Remedial"
	default:
		return false, "Tidak Lulus"
	}
}
