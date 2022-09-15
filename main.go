package main

import (
	"fmt"
)

func main() {

	var bill int
	var discount int
	var max int
	var afterDiscount int

	fmt.Println("voucher DumbMerchBerkah")
	fmt.Println("Total Belanja:")
	fmt.Scanf("%d", &bill)

	fmt.Scanf("%d", &discount)
	if bill >= 50000 {
		discount = 25
	}
	fmt.Println("Diskon:", discount, "%")

	max = bill - (bill - discount)
	if bill >= 50000 {
		max = 20000
	} else if bill <= 50000 {
		max = 0
	}
	fmt.Println("Potongan:", max)

	afterDiscount = bill - (bill * discount / 100)
	fmt.Println("Total Pembayaran: ", afterDiscount)

	fmt.Println("=======================================================")

	fmt.Println("voucher DumbMerchMurahBanget")
	fmt.Println("Total Belanja:")
	fmt.Scanf("%d", &bill)

	fmt.Scanf("%d", &discount)
	if bill >= 70000 {
		discount = 50
	}
	fmt.Println("Diskon:", discount, "%")

	max = bill - (bill - discount)
	if bill >= 70000 {
		max = 45000
	} else if bill <= 70000 {
		max = 0
	}
	fmt.Println("Potongan:", max)

	afterDiscount = bill - (bill * discount / 100)
	fmt.Println("Total Pembayaran: ", afterDiscount)
}
