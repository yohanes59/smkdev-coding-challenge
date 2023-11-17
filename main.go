package main

import (
	"fmt"
)

func main() {
	var input1 int
	fmt.Print("Masukkan Jumlah Angka: ")
	fmt.Scanf("%d", &input1)
	input2 := make([]int, input1)
	fmt.Print("Masukkan Daftar Angka: ")

	for i := 1; i < input1; i++ {
		fmt.Scan(&input2[i])
	}

	fmt.Println(findMissingNumber(input1, input2))
}

func findMissingNumber(total int, numbers []int) int {
	// hitung jumlah total dari 1 sampai n
	totalNumber := total * (total + 1) / 2
	listOfNumber := 0

	// Hitung jumlah dari angka-angka yang sudah diberikan
	for _, number := range numbers {
		listOfNumber += number
	}

	// hitung angka yang hilang
	missingNumber := totalNumber - listOfNumber

	return missingNumber
}
