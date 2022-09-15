package main

import (
	"fmt"
)

func main() {

	var bill int
	var discount int
	var afterDiscount int

	fmt.Println("voucher DumbMerchBerkah")
	fmt.Println("Total Belanja:")
	fmt.Scanf("%d", &bill)

	fmt.Println("Diskon: 25%")
	fmt.Scanf("%d", &discount)
	if bill >= 50000 {
		discount = 25
	}

	afterDiscount = bill - (bill * discount / 100)
	fmt.Println("Total Pembayaran: ", afterDiscount)

	fmt.Println("=======================================================")

	fmt.Println("voucher DumbMerchMurahBanget")
	fmt.Println("Total Belanja:")
	fmt.Scanf("%d", &bill)

	fmt.Println("Diskon: 50%")
	fmt.Scanf("%d", &discount)
	if bill >= 70000 {
		discount = 50
	}

	afterDiscount = bill - (bill * discount / 100)
	fmt.Println("Total Pembayaran: ", afterDiscount)
}
