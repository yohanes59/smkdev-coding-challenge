package main

import (
	"fmt"
	"math"
)

func main() {
	// inisialisasi variabel number
	var number int
	fmt.Print("input number = ")
	// ambil input number yang di input user
	fmt.Scan(&number)

	fmt.Print(calculateRintCube(number))
}

func calculateRintCube(number int) string {
	// inisialisasi variabel result
	result := ""
	for i := 1; i <= number; i++ {
		// proses hitung dengan package math untuk pangkat 3
		rintValue := math.Pow(float64(i), 3)
		result += fmt.Sprintf("Current Number is : %d and the cube is %.0f \n", i, rintValue)
	}
	return result
}
