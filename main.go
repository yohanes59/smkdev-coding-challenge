package main

import "fmt"

func main() {
	// inisiasi variabel n
	var n int
	fmt.Print("Masukkan Input: ")
	fmt.Scanf("%d", &n)
	fmt.Println(isHappyNumber(n))
}

func isHappyNumber(n int) bool {
	// inisiasi variabel sum
	sum := 0
	// selama nilai input yang dimasukkan lebih besar dari 0
	for n > 0 {
		// ambil digit terakhir dari nilai n dengan menggunakan modulus 10
		e := n % 10
		// menghapus digit terakhir dari n dengan membagi dengan 10
		n = n / 10
		// menghitung kuadrat dari digit yang diambil
		sum += e * e
	}

	if sum == 1 {
		// jika sum = 1, maka nilai sum adalah happy number dan return true
		return true
	} else if sum > 1 && sum <= 4 {
		// jika sum > 1 dan <= 4, maka nilai sum bukan happy number dan return false
		return false
	}
	return isHappyNumber(sum)
}
