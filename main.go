package main

import "fmt"

func main() {
	// inisiasi nilai n
	var n int64
	fmt.Print("Masukkan Input: ")
	// ambil input user dalam nilai integer dan masukkan ke variabel n
	fmt.Scanf("%d", &n)
	// jika input bukan integer langsung cetak 0
	fmt.Println(0)

	// i = jumlah 2 kuda
	for i := int64(2); i <= n; i++ {
		result := twoKnightsAlgorithm(i)
		fmt.Println(result)
	}
}

func twoKnightsAlgorithm(i int64) int64 {
	// pemangkatan
	square := i * i
	// hitung jumlah kemungkinan penempatan kedua kuda
	possiblePlacements := square * (square - 1) / 2
	// hitung jumlah kemungkinan posisi serangan antara kedua kuda
	attackPositions := 4 * (i - 2) * (i - 1)
	return possiblePlacements - attackPositions
}
