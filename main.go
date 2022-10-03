package main

import (
	"assignment-1/helpers"
	"fmt"
	"os"
	"strconv"
)

func main() {
	biodata := helpers.DataSiswa()

	if len(os.Args) != 1 {
		for i := 1; i < len(os.Args); i++ {
			if helpers.IsLetter(os.Args[i]) {
				fmt.Printf("Harap masukkan nomor absen yang valid! Absen bernilai [%v] tidak valid!\n", os.Args[i])
				continue
			}
			absen, _ := strconv.Atoi(os.Args[i])

			if absen <= 0 {
				fmt.Printf("Harap masukkan nomor absen yang valid! Absen bernilai [%v] tidak valid!\n", os.Args[i])
				continue
			}

			if absen >= len(biodata) {
				fmt.Printf("Biodata siswa dengan nomor absen [%v] tidak ditemukan karena jumlah siswa hanya %v!", absen, len(biodata))
				continue
			}
			helpers.PrintData(biodata, absen)
		}
	} else {
		fmt.Println("Masukkan nomor absen siswa setelah keyword main.go!")
		fmt.Println("Contoh: go run main.go 4 5")
	}
}
