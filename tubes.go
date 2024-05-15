package main

import "fmt"

const NMAX int = 100

type wisata struct {
	nama      string
	fasilitas string
	wahana    string
	harga     int
	jarak     float64
}

type arrWisata [NMAX]wisata

func menu() { // procedure
	fmt.Println("-------------------")
	fmt.Println("Aplikasi pariwisata")
	fmt.Println("-------------------")
	fmt.Println("1. Tambah Tempat Wisata")
	fmt.Println("2. Ubah Tempat Wisata")
	fmt.Println("3. Hapus Tempat Wisata")
	fmt.Println("4. Cari Tempat Wisata berdasarkan jarak")
	fmt.Println("5. Cari Tempat Wisata berdasarkan harga")
	fmt.Println("6. Tempat Wisata terdekat")
	fmt.Println("7. Tempat Wisata terjauh")
	fmt.Println("8. Tempat Wisata termurah")
	fmt.Println("9.Tempat Wisata termahal")
	fmt.Println("10.Cetak semua tempat wisata")
	fmt.Println("11. Exit")
}

func main() { // main program
	var pilih int
	var a arrWisata 
	var n, x int
	for pilih != 11 {
		menu()
		fmt.Println("Selamat datang di Aplikasi Pariwisata:")
		fmt.Print("Silahkan pilih (1/2/3/4/5/6/7/8/9/10/11)?")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahTW(&a, &n)
		} else if pilih == 	2 {
			ubahTW()	
		} else if pilih == 	3 {
			hapusTW()	
		} else if pilih == 4 {
			cariJarak(a, n, x)
		} else if pilih == 5 {
			cariHarga(a, n, x)
		} else if pilih == 6 {
			terdekat(a, n)
		} else if pilih == 7 {
			terjauh(a, n)
		} else if pilih == 8 {
			termahal(a, n)
		} else if pilih == 9 {
			termurah(a,n)
		} else if pilih == 10 {
			cetakTambahTW(a, n)
		} else if pilih == 11 {
			fmt.Println("Terima kasih telah menggunakan aplikasi kami")
		}
	}
}

func tambahTW(A *arrWisata, n *int) { // pass by reference procedure
	var i int
	var nama string
	for i = 0; i < NMAX && nama != "none"; i++ {
		fmt.Println("Silahkan masukkan nama Tempat Wisata")
		fmt.Scan(&A[i].nama)
		nama = A[i].nama
		if nama != "none" {
			fmt.Println("Silahkan masukkan fasilitas Tempat Wisata")
			fmt.Scan(&A[i].fasilitas)
			fmt.Println("Silahkan masukkan wahana Tempat Wisata")
			fmt.Scan(&A[i].wahana)
			fmt.Println("Silahkan masukkan harga Tempat Wisata")
			fmt.Scan(&A[i].harga)
			fmt.Println("Silahkan masukkan jarak Tempat Wisata")
			fmt.Scan(&A[i].jarak)
			*n++
		}
	}
	if nama == "none" {
		menu()
	}
}

func cetakTambahTW(A arrWisata, n int) { // pass by value procedure
	fmt.Println("Data Tempat Wisata:")
	for i := 0; i < n-1; i++ {
		fmt.Println(A[i].nama, A[i].fasilitas, A[i].wahana, A[i].harga, A[i].jarak)
	}
	menu()
}

func cariHarga(A arrWisata, n int, x int) int{ // Sequential search, function
	var ketemu, i int
	ketemu = -1
	i = 0
	for i <= n && ketemu == -1 {
        if A[i].harga == x {
            ketemu = i
        }
		i++
    }
	return ketemu
}

func cariJarak(A arrWisata, n int, x float64) float64{ // binary search, function
	
}

func terdekat(A wisata, n int) { // find min, Procedure
	var min float64
	var i int
	min = A[0].jarak
	for i = 0; i < n; i++ {
        if A[i].jarak < min {
            min = A[i]
        }
    }
	fmt.Println("Tempat wisata terdekat adalah", ,"dengan jarak", )
}

func terjauh(A wisata, n int) {  // find max, Procedure
	var max float64
	var i int
	max = A[0].jarak
	for i = 0; i < n; i++ {
        if A[i].jarak > max {
            max = A[i]
        }
    }
	fmt.Println("Tempat wisata terjauh adalah",  ,"dengan jarak",)
}

func termurah(A wisata, n int) { // find min, procedure
	var min int
    var i int
    min = A[0].harga
    for i = 0; i < n; i++ {
        if A[i].harga < min {
            min = A[i].harga
        }
    }
    fmt.Println("Tempat wisata termurah adalah", ,"dengan harga", )
}

func termahal(A wisata, n int) { // find max, procedure
	var max int
    var i int
    max = A[0].harga
    for i = 0; i < n; i++ {
        if A[i].harga > max {
            max = A[i].tinggi
        }
    }
    fmt.Println("Tempat wisata termahal adalah", ,"dengan harga", )
}