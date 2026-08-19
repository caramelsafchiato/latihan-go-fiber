package main

import "fmt"

func main() {
	var nama string = "Alisya Aninda Safwa"
	var umur int = 20
	var ipk float64 = 3.32
	var isAktif bool = true
	hobi := []string{"Bermain Roblox", ",", "Menkritik pemerintah"} 

	fmt.Println("=== Output Variabel ===")
	fmt.Printf("Nama: %s\n", nama)
	fmt.Printf("Umur: %d\n", umur)
	fmt.Printf("IPK: %.2f\n", ipk)
	fmt.Printf("Status Aktif: %t\n", isAktif)
	fmt.Printf("Hobi: %v\n\n", hobi)

	//map data mhs
	nilaiMahasiswa := make(map[string]int)

	nilaiMahasiswa["Lailia"] = 92
	nilaiMahasiswa["Valerina"] = 88
	nilaiMahasiswa["Alisya"] = 93
	fmt.Println("=== Operasi Map Data Mahasiswa ===")
	fmt.Println("Data mahasiswa berhasil ditambahkan.")

	namaCari := "Alisya"
	if nilai, ada := nilaiMahasiswa[namaCari]; ada {
		fmt.Printf("Data ditemukan -> %s: %d\n", namaCari, nilai)
	} else {
		fmt.Printf("Data %s tidak ditemukan.\n", namaCari)
	}

	fmt.Println("\nMenghapus data Valerina dari map...")
	delete(nilaiMahasiswa, "Valerina")

	fmt.Println("\nSeluruh data mahasiswa yang tersisa:")
	for namaMhs, nilai := range nilaiMahasiswa {
		fmt.Printf("- %s: %d\n", namaMhs, nilai)
	}
}

