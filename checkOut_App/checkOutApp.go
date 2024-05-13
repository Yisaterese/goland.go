package main

import (
	"fmt"
	"time"
)

func main() {
	var itemsBought []string
	var quantity []int
	var pricePerUnit []float64

	const vat = 17.50
	var total float64

	fmt.Println("What is the customer's Name?")
	var customerName string
	_, err := fmt.Scanln(&customerName)
	if err != nil {
		return
	}

	for {
		fmt.Println("What did the user buy?")
		var items string
		itemsBought = append(itemsBought, items)
		_, err := fmt.Scanln(&items)
		if err != nil {
			return
		}
		fmt.Println("How many pieces?")
		var piecesBought int
		_, err = fmt.Scanln(&piecesBought)
		if err != nil {
			return
		}
		quantity = append(quantity, piecesBought)

		fmt.Println("How much per unit?")
		var unitPrice float64
		_, err = fmt.Scanln(&unitPrice)
		if err != nil {
			return
		}
		pricePerUnit = append(pricePerUnit, unitPrice)

		fmt.Println("Add more itemsBought? (yes/no)")
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			return
		}
		if response != "yes" {
			break
		}
	}

	fmt.Println("What is your name?")
	var cashierName string
	_, err = fmt.Scanln(&cashierName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("How much discount will they get?")
	var discount float64
	_, err = fmt.Scanln(&discount)
	if err != nil {
		return
	}
	for index := range itemsBought {
		total += float64(quantity[index]) * pricePerUnit[index]
	}

	vatCaculationResult := total * (vat / 100)
	billTotalCalculation := vatCaculationResult + total

	fmt.Printf(`SUPER MART STORES
MAIN BRANCH
LOCATION: 312, HERBERT MACAULAY WAY, SABO YABA LAGOS
TEL: 02345678812
Date: %s
Cashier: %s
Customer's name: %s
===============================================
	Item     QTY     PRICE     TOTAL
-----------------------------------------------
`, time.Now().Format("02/01/06 03:04:05 PM"), cashierName, customerName)

	for i := range itemsBought {
		fmt.Printf("\t%s\t%d\t%.2f\t%.2f\n", itemsBought[i], quantity[i], pricePerUnit[i], float64(quantity[i])*pricePerUnit[i])

	}
	fmt.Printf(`-----------------------------------------------
	Sub Total:%10.2f
	Discount:%10.2f
	VAT @ 17.50%%:%10.2f
===============================================
	Bill Total:%10.2f
===============================================
	THANK YOU FOR YOUR PATRONAGE
`, total, discount, vatCaculationResult, billTotalCalculation)

	fmt.Println("How much did the customer give to you?")
	var amountPaid float64
	_, err = fmt.Scanln(&amountPaid)
	if err != nil {
		return
	}
}
