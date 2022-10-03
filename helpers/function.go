package helpers

import (
	"fmt"
	"strings"
	"unicode"
)

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func PrintData(biodata []Biodata, absen int) {
	pembatas()
	fmt.Println("Absen\t\t\t\t: ", absen)
	fmt.Println("Nama\t\t\t\t: ", biodata[absen].Nama)
	fmt.Println("Alamat\t\t\t\t: ", biodata[absen].Alamat)
	fmt.Println("Pekerjaan\t\t\t: ", biodata[absen].Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang\t: ", biodata[absen].Alasan)
}

func pembatas() {
	fmt.Println(strings.Repeat("=", 50))
}
