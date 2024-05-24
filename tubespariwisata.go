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

func menu() { // menu procedure
	fmt.Println("-------------------")
	fmt.Println("Aplikasi Pariwisata")
	fmt.Println("-------------------")
	fmt.Println("1.  Tambah Tempat Wisata")
	fmt.Println("2.  Ubah Tempat Wisata")
	fmt.Println("3.  Hapus Tempat Wisata")
	fmt.Println("4.  Cari tempat wisata berdasarkan jarak")
	fmt.Println("5.  Cari tempat wisata berdasarkan harga")
	fmt.Println("6.  Wisata terdekat")
	fmt.Println("7.  Wisata terjauh")
	fmt.Println("8.  Wisata termurah")
	fmt.Println("9.  Wisata termahal")
	fmt.Println("10. Cetak semua tempat wisata")
	fmt.Println("11. Exit")
}

func main() { // main program
	var pilih int
	var a arrWisata
	var n int
	n = 0
	for pilih != 11 {
		menu()
		fmt.Print("Silahkan pilih (1/2/3/4/5/6/7/8/9/10/11)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			inputTW(&a, &n)
		case 2:
			editTW(&a, n)
		case 3:
			hapusTW(&a, &n)
		case 4:
			cariBerdasarkanJarak(a, n)
		case 5:
			cariBerdasarkanHarga(a, n)
		case 6:
			terdekat(a, n)
		case 7:
			terjauh(a, n)
		case 8:
			termurah(a, n)
		case 9:
			termahal(a, n)
		case 10:
			cetakTW(a, n)
		case 11:
			fmt.Println("Terima kasih telah menggunakan aplikasi kami")
		default:
			fmt.Println("Pilihan tidak valid. Silahkan coba lagi.")
		}
	}
}

func inputTW(A *arrWisata, n *int) { // pass by reference procedure
	var nama string
	for *n < wis && nama != "none" {
		fmt.Print("Silahkan masukkan nama Tempat Wisata(masukkan 'none' jika sudah selesai): ")
		fmt.Scan(&A[*n].nama)
		nama = A[*n].nama
		if nama != "none" {
			fmt.Print("Silahkan masukkan fasilitas Tempat Wisata: ")
			fmt.Scan(&A[*n].fasilitas)
			fmt.Print("Silahkan masukkan wahana Tempat Wisata: ")
			fmt.Scan(&A[*n].wahana)
			fmt.Print("Silahkan masukkan harga Tempat Wisata: ")
			fmt.Scan(&A[*n].harga)
			fmt.Print("Silahkan masukkan jarak Tempat Wisata: ")
			fmt.Scan(&A[*n].jarak)
			*n++
		}
	}
}

func editTW(A *arrWisata, n int) { // procedure using sequential search
	var dicari string
	fmt.Print("Silahkan masukkan nama yang ingin diedit: ")
	fmt.Scan(&dicari)
	index := seqSearch(A, n, dicari)
	if index == -1 {
		fmt.Println("Tempat Wisata tidak ditemukan.")
	}
	fmt.Println("Menu Edit")
	fmt.Println("1. Fasilitas")
	fmt.Println("2. Wahana")
	fmt.Println("3. Harga")
	fmt.Println("4. Jarak")

	var pilihan int
	fmt.Print("Pilih data yang ingin diubah (1/2/3/4): ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Print("Silahkan masukkan fasilitas baru: ")
		fmt.Scan(&A[index].fasilitas)
	case 2:
		fmt.Print("Silahkan masukkan wahana baru: ")
		fmt.Scan(&A[index].wahana)
	case 3:
		fmt.Print("Silahkan masukkan harga baru: ")
		fmt.Scan(&A[index].harga)
	case 4:
		fmt.Print("Silahkan masukkan jarak baru: ")
		fmt.Scan(&A[index].jarak)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func seqSearch(A *arrWisata, n int, x string) int { // Sequential search, function
	for i := 0; i < n; i++ {
		if A[i].nama == x {
			return i
		}
	}
	return -1
}

func hapusTW(A *arrWisata, n *int) { // pass by reference Procedure using Sequential search
	var dicari string
	fmt.Print("Silahkan masukkan nama yang ingin dihapus: ")
	fmt.Scan(&dicari)
	index := seqSearch(A, *n, dicari)
	if index == -1 {
		fmt.Println("Tempat Wisata tidak ditemukan.")
	}
	for i := index; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func cariBerdasarkanJarak(A arrWisata, n int) {
	// Selection sort ascending
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if A[j].jarak < A[minIdx].jarak {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}

	var jarak float64
	fmt.Print("Masukkan jarak maksimum: ")
	fmt.Scan(&jarak)

	// Binary search
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if A[mid].jarak <= jarak {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println("Tempat Wisata dengan jarak kurang dari atau sama dengan", jarak, "KM:")
	for i := 0; i < left; i++ {
		fmt.Println(A[i].nama, "-", A[i].jarak, "KM")
	}
}

func cariBerdasarkanHarga(A arrWisata, n int) { // Procedure
	// insertion sort ascending
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j].harga > key.harga {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}

	var harga int
	fmt.Print("Masukkan harga maksimum: ")
	fmt.Scan(&harga)
	fmt.Println("Tempat Wisata dengan harga kurang dari atau sama dengan", harga, ":")
	for i := 0; i < n; i++ {
		if A[i].harga <= harga {
			fmt.Println(A[i].nama, "-", A[i].harga)
		}
	}
}

func terdekat(A arrWisata, n int) { // find min, Procedure
	if n == 0 {
		fmt.Println("Tidak ada data tempat wisata.")
	}
	min := A[0].jarak
	index := 0
	for i := 1; i < n; i++ {
		if A[i].jarak < min {
			min = A[i].jarak
			index = i
		}
	}
	fmt.Println("Tempat wisata terdekat adalah", A[index].nama, "dengan jarak", A[index].jarak, "KM")
}

func terjauh(A arrWisata, n int) { // find max, Procedure
	if n == 0 {
		fmt.Println("Tidak ada data tempat wisata.")
	}
	max := A[0].jarak
	index := 0
	for i := 1; i < n; i++ {
		if A[i].jarak > max {
			max = A[i].jarak
			index = i
		}
	}
	fmt.Println("Tempat wisata terjauh adalah", A[index].nama, "dengan jarak", A[index].jarak, "KM")
}

func termurah(A arrWisata, n int) { // find min, Procedure
	if n == 0 {
		fmt.Println("Tidak ada data tempat wisata.")
	}
	min := A[0].harga
	index := 0
	for i := 1; i < n; i++ {
		if A[i].harga < min {
			min = A[i].harga
			index = i
		}
	}
	fmt.Println("Tempat wisata termurah adalah", A[index].nama, "dengan harga", A[index].harga)
}

func termahal(A arrWisata, n int) { // find max, Procedure
	if n == 0 {
		fmt.Println("Tidak ada data tempat wisata.")
	}
	max := A[0].harga
	index := 0
	for i := 1; i < n; i++ {
		if A[i].harga > max {
			max = A[i].harga
			index = i
		}
	}
	fmt.Println("Tempat wisata termahal adalah", A[index].nama, "dengan harga", A[index].harga)
}

func cetakTW(A arrWisata, n int) { // pass by value procedure
	fmt.Println("Data Tempat Wisata:")
	for i := 0; i < n; i++ {
		fmt.Printf("Nama: %s, Fasilitas: %s, Wahana: %s, Harga: %d, Jarak: %.2f KM\n", A[i].nama, A[i].fasilitas, A[i].wahana, A[i].harga, A[i].jarak)
	}
}
