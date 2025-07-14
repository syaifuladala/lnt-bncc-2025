package main

import "fmt"

func main() {
	// fmt.Println("Hello World :)")

	// var (
	// 	nama string
	// 	age  int
	// )
	// nama = "Binus"
	// age = 18

	// namaLengkap := "Binusian"

	// const pi = 3.14

	// fmt.Println("nama:", nama)
	// fmt.Println("nama lengkap:", namaLengkap)
	// fmt.Println("umur:", age)
	// fmt.Println(pi)

	// fmt.Println("===Program Input-Output===")
	// // nama
	// var nama string
	// fmt.Print("Masukkan nama anda: ")
	// fmt.Scanln(&nama)
	// fmt.Println("Halo, ", nama)

	// // umur
	// var age int
	// fmt.Print("Masukkan umur anda: ")
	// fmt.Scanln(&age)
	// fmt.Println("umur saya ", age)

	// fmt.Println("=====SUMMARY======")
	// fmt.Printf("halo perkenalkan nama saya adalah %s, umur saya %d tahun", nama, age)

	// var numbers []int
	// numbers = append(numbers, 10)
	// numbers = append(numbers, 20)
	// fmt.Println(numbers)

	// nilai := 70

	// if nilai > 75 {
	// 	fmt.Println("Nilai anda ", nilai)
	// 	fmt.Println("Lulus")
	// } else if nilai >= 60 {
	// 	fmt.Println("Nilai anda ", nilai)
	// 	fmt.Println("Remidi")
	// } else {
	// 	fmt.Println("Tidak Lulus")
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// var buah = []string{"apel", "jeruk", "mangga"}
	// for _, value := range buah {
	// 	fmt.Println(value)
	// }

	// hari := "rabu"
	// switch hari {
	// case "senin":
	// 	fmt.Println("ini hari senin")
	// case "selasa":
	// 	fmt.Println("ini hari selasa")
	// case "sabtu", "minggu":
	// 	fmt.Println("ini weekend")
	// default:
	// 	fmt.Println("hari biasa")
	// }

	// // pointer
	// var rumahAnjing *string
	// fmt.Println(rumahAnjing)

	// anjing := "puppy"
	// fmt.Println(anjing)

	// rumahAnjing = &anjing
	// fmt.Println(rumahAnjing)

	// fmt.Println(*rumahAnjing)

	// struct
	type Alamat struct {
		Jalan   string
		Kota    string
		KodePos int
	}

	// noHPBinus := "081234567890"
	mahasiswa := struct {
		Nama   string
		Umur   int
		Hobi   []string
		Alamat Alamat
		Aktif  bool
		NoHP   *string
	}{
		Nama: "binus",
		Umur: 18,
		Hobi: []string{"bertani", "meramu"},
		Alamat: Alamat{
			Jalan:   "jalan damai",
			Kota:    "bandung",
			KodePos: 1234,
		},
		Aktif: true,
		// NoHP:  &noHPBinus,
	}

	fmt.Println(mahasiswa)
	fmt.Println("nama saya: ", mahasiswa.Nama)
	fmt.Println("umur saya: ", mahasiswa.Umur)
	fmt.Println("jalan: ", mahasiswa.Alamat.Jalan)
	if mahasiswa.NoHP != nil {
		fmt.Println("no hp: ", *mahasiswa.NoHP)
	} else {
		fmt.Println("no hp: -")
	}
}
