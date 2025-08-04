package main

import (
	"fmt"
	"sync"
)

func main() {
	// // ===============
	// // PROSES ASYNC GO
	// // ===============
	// var wg sync.WaitGroup // Digunakan untuk menunggu proses async

	// fmt.Println("Mulai proses simpan data...")
	// simpanData() // proses utama, blocking

	// wg.Add(1) // Tambah 1 task yang harus ditunggu

	// go func() {
	// 	defer wg.Done() // Sinyal bahwa proses selesai
	// 	kirimEmail()    // proses async
	// }()

	// fmt.Println("Menunggu proses kirim email selesai...")
	// wg.Wait() // Tunggu sampai semua proses async selesai

	// fmt.Println("Selesai semua.")

	// =================
	// PROSES GO ROUTINE
	// =================
	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		wg.Add(1) // Tambah counter untuk setiap goroutine

		go func(n int) {
			defer wg.Done() // Kurangi counter setelah selesai
			fmt.Printf("Proses ke-%d\n", n)
		}(i)
	}

	wg.Wait() // Tunggu semua goroutine selesai
	fmt.Println("Semua proses selesai.")
}

// func simpanData() {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Data berhasil disimpan.")
// }

// func kirimEmail() {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Email berhasil dikirim.")
// }
