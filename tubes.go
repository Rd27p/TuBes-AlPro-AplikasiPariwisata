package main

import "fmt"

const wis int = 100

type wisata struct {
	nama      string
	fasilitas string
	wahana    string
	harga     int
	jarak     float64
}

type arrWisata [wis]wisata

func menu() { // procedure
	fmt.Println("-------------------")
	fmt.Println("Aplikasi pariwisata")
	fmt.Println("-------------------")
	fmt.Println("1. Tambah Tempat Wisata")
	fmt.Println("2. Ubah Tempat Wisata")
	fmt.Println("3. Hapus Tempat Wisata")
	fmt.Println("4. Cari Tempat wisata")
	fmt.Println("5. Tempat Wisata terdekat")
	fmt.Println("6. Tempat Wisata terjauh")
	fmt.Println("7. Tempat Wisata termurah")
	fmt.Println("8. Tempat Wisata termahal")
	fmt.Println("9. Cetak semua tempat wisata")
	fmt.Println("10.Exit")
}

func main() { // main program
	var pilih int
	var a arrWisata
	var n int
	for pilih != 10 {
		menu()
		fmt.Println("Selamat datang di Aplikasi Pariwisata:")
		fmt.Print("Silahkan pilih (1/2/3/4/5/6/7/8/9/10)?")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahTW(&a, &n)
		} else if pilih == 2 {
			editTW(&a, n)
		} else if pilih == 3 {
			hapusTW(&a, n)
		} else if pilih == 4 {
			fmt.Println()
		} else if pilih == 5 {
			terdekat(a, n)
		} else if pilih == 6 {
			terjauh(a, n)
		} else if pilih == 7 {
			termurah(a, n)
		} else if pilih == 8 {
			termahal(a, n)
		} else if pilih == 9 {
			cetakTW(a, n)
		} else if pilih == 10 {
			fmt.Println("Terima kasih telah menggunakan aplikasi kami")
		}
	}
}

func tambahTW(A *arrWisata, n *int) { // pass by reference procedure
	var i int
	var nama string
	for i = 0; i < wis && nama != "none"; i++ {
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
}

func editTW(A *arrWisata, n int) {
	var index int
	var nama, dicari string
	fmt.Println("Silahkan masukkan nama yang ingin diedit")
	fmt.Scan(&dicari)
	index = seqSearch(A, n, dicari)

}
func hapusTW(A *arrWisata, n int) {
	var index int
	var nama, dicari string
	fmt.Println("Silahkan masukkan nama yang ingin diedit")
	fmt.Scan(&dicari)
	index = seqSearch(A, n, dicari)
}

func seqSearch(A *arrWisata, n int, x string) int { // Sequential search, function
	var i, index int
	for i = 0; i < n; i++ {
		if A[i].nama == x {
			index = i
		}
	}
	return index
}

func terdekat(A arrWisata, n int) { // find min, Procedure
	var min float64
	var i, index int
	min = A[0].jarak
	for i = 0; i < n; i++ {
		if A[i].jarak <= min {
			min = A[i].jarak
			index = i
		}
	}
	fmt.Println("Tempat wisata terdekat adalah", A[index].nama, "dengan jarak", A[index].jarak, "KM")
}

func terjauh(A arrWisata, n int) { // find max, Procedure
	var max float64
	var i, index int
	max = A[0].jarak
	for i = 0; i < n; i++ {
		if A[i].jarak >= max {
			max = A[i].jarak
			index = i
		}
	}
	fmt.Println("Tempat wisata terjauh adalah", A[index].nama, "dengan jarak", A[index].jarak, "KM")
}

func termahal(A arrWisata, n int) { // find max, procedure
	var max, index int
	var i int
	max = A[0].harga
	for i = 0; i < n; i++ {
		if A[i].harga >= max {
			max = A[i].harga
			index = i
		}
	}
	fmt.Println("Tempat wisata termahal adalah", A[index].nama, "dengan harga", A[index].harga)
}

func termurah(A arrWisata, n int) { // find min, procedure
	var min, index int
	var i int
	min = A[0].harga
	for i = 0; i < n; i++ {
		if A[i].harga <= min {
			min = A[i].harga
			index = i
		}
	}
	fmt.Println("Tempat wisata termurah adalah", A[index].nama, "dengan harga", A[index].harga)
}

func cetakTW(A arrWisata, n int) { // pass by value procedure
	fmt.Println("Data Tempat Wisata:")
	for i := 0; i < n; i++ {
		fmt.Println(A[i].nama, A[i].fasilitas, A[i].wahana, A[i].harga, A[i].jarak)
	}
}
