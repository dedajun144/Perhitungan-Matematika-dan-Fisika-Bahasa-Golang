package main

import (
	"fmt"
	"math"
)

func persegi(sisi int) int {
	return sisi * sisi
}
func segitiga(alas int, tinggi int) int {
	return alas * tinggi / 2
}
func lingkaran(r float32) float32 {
	return r * r * 3.14
}
func glb(s float64, t float64) float64 {
	return s / t //rumus menghitung glb
}
func glbb() {
	var menu int
	fmt.Println("Gerak Lurus Berubah Beraturan")
	fmt.Println("1. Mencari kecepatan akhir jika waktu diketahui")
	fmt.Println("2. Mencari jarak tempuh")
	fmt.Println("3. Mencari kecepatan akhir jika jarak diketahui")
	fmt.Scan(&menu) //input data untuk memilih menu
	switch menu {
	case 1:
		var (
			v0, t, a float64
		)
		fmt.Println("Masukkan kecepatan awal dalam satuan m/s : ")
		fmt.Scan(&v0)
		fmt.Println("Masukkan waktu dalam satuan sekon : ")
		fmt.Scan(&t)
		fmt.Println("Masukkan percepatan dalam satuan m/s2 : ")
		fmt.Scan(&a)
		vt := v0 + a*t                              //rumus perhitungan gerak lurus berubah beraturan
		fmt.Println("Kecepatan Akhir =", vt, "m/s") //output dari hasil penghitungan data yang telah di input
		break
	case 2:
		var (
			v0, t, a float64
		)
		fmt.Println("Masukkan kecepatan awal dalam satuan m/s : ")
		fmt.Scan(&v0)
		fmt.Println("Masukkan waktu dalam satuan sekon : ")
		fmt.Scan(&t)
		fmt.Println("Masukkan percepatan dalam satuan m/s2 : ")
		fmt.Scan(&a)
		s := v0*t + 0.5*a*t*t                     //rumus perhitungan gerak lurus berubah beraturan
		fmt.Println("Jarak Tempuh =", s, "sekon") //output dari hasil penghitungan data yang telah di input
		break
	case 3:
		var (
			v0, a, s float64
		)
		fmt.Println("Masukkan kecepatan awal dalam satuan m/s : ")
		fmt.Scan(&v0)
		fmt.Println("Masukkan jarak dalam satuan meter : ")
		fmt.Scan(&s)
		fmt.Println("Masukkan percepatan dalam satuan m/s2 : ")
		fmt.Scan(&a)
		vt2 := v0*v0 + 2*a*s                         //rumus perhitungan gerak lurus berubah beraturan
		fmt.Println("Kecepatan Akhir =", vt2, "m/s") //output dari hasil penghitungan data yang telah di input
		break
	}
}
func energi() {
	var c int
	fmt.Print("rumus energi ")
	fmt.Print("1. energi potensial\n")
	fmt.Print("2. energi Kinetik \n ")
	fmt.Print("Masukkan pilihan anda \n")
	fmt.Scan(&c)
	switch c {
	case 1:
		var m, h, ep float64
		const G = 10
		fmt.Print("Masukkan massa : \n")
		fmt.Scan(&m)
		fmt.Print("Masukkan tinggi : \n")
		fmt.Scan(&h)

		ep = m * G * h
		fmt.Print("Energi potensialnya adalah ", ep)
		break
	case 2:
		var m, v, ek float64
		fmt.Print("Rumus Kinetik \n")
		fmt.Print("Masukkan massa : \n")
		fmt.Scan(&m)
		fmt.Print("masukkan kecepatan : \n")
		fmt.Scan(&v)
		ek = 0.5 * m * v * v
		fmt.Print("Maka besar energi kinetik adalah ", ek)
		break
	}
}
func getaranDanFrekuensi() {
	var d int
	fmt.Print("Getaran dan Frekuensi \n")
	fmt.Print("1. Getaran \n")
	fmt.Print("2. Frekuensi \n")
	fmt.Print("Masukkan menu anda ")
	fmt.Scan(&d)
	switch d {
	case 1:
		var n, t, f float64
		fmt.Print("Masukkan jumlah getaran ")
		fmt.Scan(&n)
		fmt.Print("Masukkan waktu")
		fmt.Scan(&t)
		f = n / t
		fmt.Println("Maka frekuensi adalah : ", f)
		break
	case 2:
		var t, n, T float64
		fmt.Print("Masukkan waktu")
		fmt.Scan(&t)
		fmt.Print("Masukkan jumlah getaran ")
		fmt.Scan(&n)
		T = t / n
		fmt.Println("Maka Periode adalah ", T)
		break
	}
}
func massaJenis() {

	var massa, volume, mj float64
	fmt.Print("Masukkan Massa : \n")
	fmt.Scan(&massa)
	fmt.Print("Masukkan volume : \n")
	fmt.Scan(&volume)
	mj = massa / volume
	fmt.Print("Maka massa jenis adalah ", mj)
}
func gayaUsahaTekanandaya() {
	var e int
	fmt.Print("1. Mencari gaya \n")
	fmt.Print("2. Usaha \n")
	fmt.Print("3. Tekanan \n")
	fmt.Print("4. daya \n ")
	fmt.Print("Masukkan pilihan anda : ")
	fmt.Scan(&e)
	switch e {
	case 1:
		var m, percepatan, F float64
		fmt.Println("Masukkan massa ")
		fmt.Scan(&m)
		fmt.Println("Masukkan percepatan ")
		fmt.Scan(&percepatan)
		F = m * percepatan
		fmt.Println("Gaya adalah ", F)
		break
	case 2:
		var F, s, W float64
		fmt.Println("Masukkan Gaya ")
		fmt.Scan(&F)
		fmt.Println("masukkan jarak ")
		fmt.Scan(&s)
		W = F * s
		fmt.Println("Maka Usaha adalah ", W)
		break
	case 3:
		var F, A, p float64
		fmt.Println("masukkan gaya ")
		fmt.Scan(&F)
		fmt.Print("Luas alas penampang ")
		fmt.Scan(&A)
		p = F / A
		fmt.Print("maka tekanan adalah ", p)
		break
	case 4:
		var S, W, t float64
		fmt.Println("Masukkan usaha ")

		fmt.Scan(&W)
		fmt.Println("Masukkan waktu ")
		fmt.Scan(&t)
		S = W / t
		fmt.Println("Jadi daya adalah ", S)
		break
	}
}
func konversi() {
	var C, K, R, Fah float64
	fmt.Println("Masukkan Nilai celcius ")
	fmt.Scan(&C)
	Fah = ((C * 1.8) + 32)
	K = C + 273.15
	R = C * 0.8
	fmt.Println("maka nilai Fahrenheit adalah ", Fah)
	fmt.Println("Maka nilai Kelvin adalah ", K)
	fmt.Println("Maka nilai reamour adalah ", R)
}
func main() {
	var s, pilih int
	fmt.Println("Aplikasi Demak")
	fmt.Print("1. luas persegi\n")
	fmt.Print("2. luas segitiga \n")
	fmt.Print("3. luas lingkaran \n")
	fmt.Print("4. Menmenghitung sudut sinus, cosinus, tangen \n")
	fmt.Print("5. menghitung gerak lurus beraturan \n")
	fmt.Print("6. menghitung gerak lurus berubah beraturan \n")
	fmt.Print("7. menghitung energi potensial, kinetik \n")
	fmt.Print("8. menghitung frekuensi atau getaran \n")
	fmt.Print("9. menghitung frekuensi atau getaran! \n")
	fmt.Print("10. menghitung daya, tekanan, usaha dan gaya \n")
	fmt.Print("11. konversi untuk semua satuan suhu \n")
	fmt.Println("Pilih [1-11] : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		var sisi int
		fmt.Println("Masukkan Sisi : ")
		fmt.Scan(&sisi)
		fmt.Println("Luas Persegi : ", persegi(s))
	case 2:
		var a, b int
		fmt.Println("Masukkan Alas : ")

		fmt.Scan(&a)
		fmt.Println("Masukkan Tinggi : ")
		fmt.Scan(&b)
		fmt.Println("Luas Segitiga : ", segitiga(a, b))
	case 3:
		var r float32
		fmt.Println("Masukkan Jari-jari : ")
		fmt.Scan(&r)
		fmt.Println("Luas Lingkaran : ", lingkaran(r))
	case 4:
		var sin, nilai float64
		fmt.Println("Masukkan sinnya :")
		fmt.Scan(&sin)
		nilai = math.Sin(sin)
		fmt.Println(nilai)
		var cos, nilai2 float64
		fmt.Println("Masukkan cosnya :")
		fmt.Scan(&cos)
		nilai2 = math.Cos(cos)
		fmt.Println(nilai2)
		var tan, nilai3 float64
		fmt.Println("Masukkan Tannya :")
		fmt.Scan(&tan)
		nilai3 = math.Tan(tan)
		fmt.Println(nilai3)
	case 5:
		var (
			s, t float64
		)
		fmt.Println("Masukkan jarak dalam satuan meter : ")
		fmt.Scan(&s) //input data
		fmt.Println("Masukkan waktu dalam satuan sekon : ")
		fmt.Scan(&t)                                 //input data
		fmt.Println("Kecepatan =", glb(s, t), "m/s") //output dari hasil penghitungan data yang telah di input
	case 6:
		glbb()
	case 7:
		energi()
	case 8:
		getaranDanFrekuensi()
	case 9:
		massaJenis()
	case 10:
		gayaUsahaTekanandaya()
	case 11:
		konversi()
	}
}
