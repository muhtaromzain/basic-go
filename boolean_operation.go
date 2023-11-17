package main

import "fmt"

func main() {
	var benar = true
	var salah = false

	var dan bool = benar && salah
	var atau bool = benar || salah
	var kebalikan bool = !benar

	fmt.Println(dan)
	fmt.Println(atau)
	fmt.Println(kebalikan)

	fmt.Println("Cek kelulusan")

	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80

	var lulus = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}