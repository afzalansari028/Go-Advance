package main

import "fmt"

type Shop struct {
	Item  string
	Price int
}

var total int
var totalBill int
var itemSlice []string
var quantitySlice []int
var paymentSlice []int

var item string
var quantity int
var payment int

func main() {

	fmt.Println("*********Welcome the grocery shop*******")
	for {
		fmt.Println("---------Below items are available---------")

		fmt.Println("1:Rice ₹40/kg")
		fmt.Println("2:Salt ₹20/kg")
		fmt.Println("3:Daal ₹120/kg")
		fmt.Println("4:Oil ₹180/l")
		fmt.Println("5:Ghee ₹350/l")
		fmt.Println("6:Quit press exit")
		fmt.Println("-------------------------------------------")

		Rice := Shop{Item: "Rice", Price: 40}
		Salt := Shop{Item: "Salt", Price: 20}
		Daal := Shop{Item: "Daal", Price: 120}
		Oil := Shop{Item: "Oil", Price: 180}
		Ghee := Shop{Item: "Ghee", Price: 350}

		fmt.Println("Enter item name for purchase or type exit :")
		fmt.Scan(&item)
		if item == "exit" {
			break
		}
		fmt.Println("Enter quantity:")
		fmt.Scan(&quantity)

		switch item {
		case "Rice":
			payment = Rice.Price * quantity
			calculateBill(payment)
		case "Salt":
			payment = Salt.Price * quantity
			calculateBill(payment)

		case "Daal":
			payment = Daal.Price * quantity
			calculateBill(payment)

		case "Oil":
			payment = Oil.Price * quantity
			calculateBill(payment)

		case "Ghee":
			payment = Ghee.Price * quantity
			calculateBill(payment)
		default:
			fmt.Println("Please Enter valid name of product")
		}
	}
	fmt.Println()
	fmt.Println("---Purchased items are listed below---")
	fmt.Println("Items   Quantity(kg/L)  Amount(Rs)")
	for i := 0; i < len(itemSlice); i++ {
		fmt.Print(itemSlice[i], "           ")
		fmt.Print(quantitySlice[i], "             ")
		fmt.Print(paymentSlice[i])
		fmt.Println()
	}
	fmt.Println("-----------------------------------")
	// fmt.Println("Purchased items are :", itemSlice)
	// fmt.Println("Items Quantity are :", quantitySlice)
	// fmt.Println("Each payment are :", paymentSlice)
	fmt.Println("Total Bill : ", totalBill)
	fmt.Println("-----------------------------------")
	fmt.Println()
}

func calculateBill(payment int) int {
	total = total + payment
	totalBill = total
	itemSlice = append(itemSlice, item)
	quantitySlice = append(quantitySlice, quantity)
	paymentSlice = append(paymentSlice, payment)
	return totalBill
}
