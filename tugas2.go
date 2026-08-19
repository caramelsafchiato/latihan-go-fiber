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
}