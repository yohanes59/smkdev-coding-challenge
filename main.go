package main

import (
	"fmt"
)

func main() {
	// inisialisasi variabel term
	var n int
	fmt.Print("Number of terms: ")
	// ambil input nilai term yang di inputkan
	fmt.Scan(&n)

	// panggil function calculateSeriesSum untuk menghitung sebanyak n (nilai term)
	sum := calculateSeriesSum(n)

	// cetak hasil perhitungan
	fmt.Printf("Sum of series up to %d term: %d\n", n, sum)
}

func calculateSeriesSum(n int) int {
	// inisialisasi variabel sum dan term
	sum := 0
	term := 0

	// looping i sebanyak nilai term
	for i := 1; i <= n; i++ {
		// inisialisasi nilai term
		term = term*10 + 2
		sum += term
	}
	return sum
}

/* output
Number of terms: 5
Sum of series up to 5 term: 24690 */
