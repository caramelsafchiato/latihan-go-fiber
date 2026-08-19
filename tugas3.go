package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func updateSlice(s *[]string, newItem string) {
	*s = append(*s, newItem)
}

//-perbandingan-

// Pass by Value (Hanya mengirim salinan)
func ubahNilai(x int) {
	x = 100
}

// Pass by Pointer (Mengirim alamat memori asli)
func ubahLewatPointer(x *int) {
	*x = 100
}

func main() {
	fmt.Println("1. Pembuktian Function swap")
	a, b := 10, 20
	fmt.Printf("Sebelum swap: a = %d, b = %d\n", a, b)
	swap(&a, &b) 
	fmt.Printf("Setelah swap: a = %d, b = %d\n\n", a, b)

	fmt.Println("2. Pembuktian Function updateSlice")
	daftar := []string{"Apel", "Jeruk"}
	fmt.Printf("Slice awal  : %v\n", daftar)
	updateSlice(&daftar, "Mangga")
	fmt.Printf("Slice update: %v\n\n", daftar)

	fmt.Println("3. Perbandingan Pass by Value vs Pointer")
	num := 42
	fmt.Printf("Nilai awal num: %d\n", num)
	
	ubahNilai(num)
	fmt.Printf("Setelah ubahNilai (Pass by Value): %d\n", num)
	
	ubahLewatPointer(&num)
	fmt.Printf("Setelah ubahLewatPointer (Pass by Pointer): %d\n", num)
}