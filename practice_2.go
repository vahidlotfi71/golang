package main

import (
	"fmt"
	//"sort"
)

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

	// oder1 , finalPrice , tax := calculate("single" , 1 , 1)

	// fmt.Printf("price : %v ,tax: %.2f , totalPrice : %f\n", finalPrice, tax , oder1)


//===========================================================

// loges("vahid" , "lotfi" ,1371 , 32 , true , "hasan" , "mehtari" , 28)

//======================================================================

// sumetion1 , multiply1 := calculator(12,12,12,12,121,2,12,12,12,12,12 ,0)
// sumetion2 , multiply2 := calculator(1,2)

// fmt.Printf("sum : %d, mul : %d\n" , sumetion1 , multiply1)
// fmt.Printf("sum : %d, mul : %d" , sumetion2, multiply2)


//==========================================================


// number_1 := []int{12,1,2,3,65,32,53,1,5,8,34,26}
// number_2 := []int{12,1,2,3,65,32,53,1,5,8,34,26}



// sort.Slice(number_1, func(i , j int )bool{
// 	return number_1[i] > number_1[j]	
// })

// fmt.Println("number_1 : ", number_1)
// fmt.Println("number_2 : ", number_2)

// sortingFunction := func (a, b int) bool{
// 	return number_2[a] > number_2[b]
// }

// sort.Slice(number_2 , sortingFunction)

// fmt.Printf("number_1: %d\n", number_1)
// fmt.Printf("number_2: %d\n", number_2)

//====================================================


//به فانکشنن هایی که درون یه فانشن دیگر تعریف می شود این لاین فانکشن هم گفته می شود 


func() {
	fmt.Println("Hello, world!")
}() // با پارانتز تابع را کال می کنیم

my_fumction := func(){
	fmt.Println("Hello vahid how you are")
}

my_fumction()

fmt.Println("sum : ",func(numbers... int)(int){
	var total int 
	for _, value := range numbers {
		total += value
	}
	return total
}(1,2,3,4,5,6	))


}





// func calculator(numbers ...int) (sum int , mul int ) { // جنس نامبر از نوع اسلایس است
// 	mul = 1
// 	for _ , number := range numbers {
// 		sum += number
// 		mul *= number
// 	}
// 	return 

// }





// func loges (values ...interface{}) {
// 	for Index , val := range values {
// 		fmt.Printf("index : %v, value : %v\n", Index , val)
// 		println("index : ", Index , "value : ", val) // خروجی این تابع قابل فهم نمی باشد  حتما باید از پکیج اف ام تی استفاده کنیم

// 	}
// }





// func calculate(roomType string , night int , personNumber int ) (totalPrice float64, price int , tax float64){ // همین جا خورجی هامون را مشخص می کنم

// 	switch roomType {
// 	case "single":
// 		price = night * personNumber * 100
		
// 	case "duble":
// 		price = night * personNumber * 200
	
// 	case "suite":
// 		price = night * personNumber * 400
// 	default :
// 		println("unknown roomType value")
// 		break
// 	}
// 	price = price
// 	tax = float64(price) * 0.09
// 	totalPrice = float64(price) + float64(tax)

// 	return  // دیگر دراینجاخروجی هامون رانمی اوریم

// }




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
