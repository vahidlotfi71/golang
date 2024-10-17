package main

import "fmt"

func main() {
	// 	print_1()
	// 	println(print_2())
	// 	println(print_3("hello vahid"))
	// 	println(print_4("Hello"))

	// }

	//================================================================

	// order_1, tax_1 := calculateRoomPrice("silge", 2, 2)
	// order_2 , tax_2 := calculateRoomPrice("suite",3 , 4)

	// fmt.Printf("Oder_1: %v , tax: %.2f\n", order_1, tax_1)
	// fmt.Printf("Oder_2: %v ,tax: %.2f\n", order_2, tax_2)

	//====================================================================

	oder1 , finalPrice , tax := calculate("single" , 1 , 1)

	fmt.Printf("price : %v ,tax: %.2f , totalPrice : %f\n", finalPrice, tax , oder1)

}



func calculate(roomType string , night int , personNumber int ) (totalPrice float64, price int , tax float64){ // همین جا خورجی هامون را مشخص می کنم

	switch roomType {
	case "single":
		price = night * personNumber * 100
		
	case "duble":
		price = night * personNumber * 200
	
	case "suite":
		price = night * personNumber * 400
	default :
		println("unknown roomType value")
		break
	}
	price = price
	tax = float64(price) * 0.09
	totalPrice = float64(price) + float64(tax)

	return  // دیگر دراینجاخروجی هامون رانمی اوریم

}




// func calculateRoomPrice(roomType string, night int, personNumber int) (int, float64) {
// 	var price int
// 	var tax float64

// 	switch roomType {
// 	case "silge":
// 			price = night * personNumber * 100

// 	case "dubble":
// 			price = night * personNumber * 200
		
// 	case "suite":
// 			price = night * personNumber * 400
// 	default :
// 			println("Unknown type for roomtype")
// 			break
// 	}

	
// 	tax = float64(price) * 0.09
// 	price = price + (int(tax))
// 	return price, tax
// }	

// func print_1(){
// 	println("hello vahid")
// }

// func print_2() string{
// 	return	"hello vahid"
// }

// func print_3(Enter string) string{
// 	return Enter
// }

// func print_4(Enter string) string{
// 	return Enter + " " +"vahid"
// }
