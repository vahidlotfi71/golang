package main

import (
	"fmt"
	//"strings"
)

func main() {

	// i, j := 12, 14
	// i = i * 2
	// var IP *int = &i // نحوه استفاده از اشاره گر
	// var JP *int = &j

	// KP := &i
	// fmt.Println("IP: ", *IP, " JP: ", JP) // حتما موقع فراخوانی باید از * استفاده شود وگرنه چیز دیگری را چاپ می کندد
	// fmt.Println("kp : ", *KP)

	//================================================================

	// mystr := "vahid lotfi from Iran"

	// for index := 0; index < len(mystr); index++ {
	// 	println("index: ", index)
	// }

	//================================================================

	// var i1 int = 12
	// var i2 int // مقدار 0 میگیرد
	// var s1 string = "vahid"
	// var s2 string // مقدار ""میگیرد

	// var (
	// 	j1 int = 12
	// 	j2 int
	// 	j3 string = "lotfi"
	// 	j4 string
	// )

	// fmt.Println("i1: ", i1)
	// println("i2: ", i2)
	// println("s1: ", s1)
	// println("s2: ", s2)
	// println("j1: ", j1)
	// println("j2: ", j2)
	// println("j3: ", j3)
	// println("j4: ", j4)

	// =================================================================
	// // const = برای مقدار دهی ثابت
	// const google_URL = "https://www.google"  // دیگر هیج جای برنامه نمی توان مقدار کانست را تغییر داد
	// const google_maps = "/maps"

	// const(
	// 	name string = "vahidd"
	// 	family string = "lotif"
	// 	age int = 32
	// )

	// fmt.Println(google_URL +google_maps)
	// fmt.Println("name: " , name , "family: " , family , "age: " , age)

	//================================================================

	// global := 100
	// {
	// 	var local int = 200
	// 	fmt.Println("local: " , local)
	// 	fmt.Println("global: " , global)
	// }

	// {
	// 	global = global + 200
	// 	println("global: " , global)
	// }

	//================================================================

	// mystring  := "hi Im vahid from IRAN "

	// for _ , word := range mystring {
	// 	fmt.Println("word: " ,word)
	// }

	//================================================================

	// var salery float64
	// var minSalery float64 = 5000000
	// var taxPercent float64

	// fmt.Print("Enter your salery : ")
	// fmt.Scan(&salery)

	// if salery <= minSalery {
	// 	taxPercent = 0
	// } else if salery > minSalery && salery < minSalery *2 {
	// 	taxPercent = .05
	// }else if salery >= minSalery * 2 && salery < minSalery * 4 {
	// 	taxPercent = .10
	// } else {
	// 	taxPercent = .20
	// }

	// fmt.Println("salery = ", salery - salery * taxPercent)
	// fmt.Println("taxPercent = ", taxPercent)

	//================================================================

	// var salery float64
	// var taxPercent float64
	// var minSalery float64 = 7000000

	// fmt.Print("Enter your salery : ")
	// fmt.Scan(&salery)

	// switch {
	// case salery <= minSalery:
	// 	taxPercent = 0
	// case salery > minSalery && salery < minSalery *2 :
	// 	taxPercent = .05
	// case salery >= minSalery*2 && salery < minSalery * 4:
	// 	taxPercent = .10
	// default :
	// 	taxPercent = .20
	// }

	// fmt.Printf("salery = %2f\n", salery - taxPercent*salery)
	// fmt.Println("taxPercent = ", taxPercent)

	//================================================================

	// const (
	// 	number1 = 29
	// 	number2 = 30
	// 	number3 = 31
	// )

	// var (
	// 	month int
	// 	totalDays int
	// )
	// fmt.Print("Enter month number for calculate days : ")
	// fmt.Scan(&month)

	// switch month {
	// case 12 :
	// 	totalDays += number1
	// 	fallthrough
	// case 11:
	// 	totalDays += number2
	// 	fallthrough
	// case 10:
	// 	totalDays += number2
	// 	fallthrough
	// case 9:
	// 	totalDays += number2
	// 	fallthrough
	// case 8:
	// 	totalDays += number2
	// 	fallthrough
	// case 7:
	// 	totalDays += number2
	// 	fallthrough
	// case 6:
	// 	totalDays += number3
	// 	fallthrough
	// case 5:
	// 	totalDays += number3
	// 	fallthrough
	// case 4:
	// 	totalDays += number3
	// 	fallthrough
	// case 3:
	// 	totalDays += number3
	// 	fallthrough
	// case 2:
	// 	totalDays += number3
	// 	fallthrough
	// case 1:
	// 	totalDays += number3

	// }

	// fmt.Println("totalDays: ", totalDays)

	//================================================================

	// var notificationType string

	// fmt.Print("Enter notification type'sms','email' : ")
	// fmt.Scan(&notificationType)

	// switch {
	// case strings.Contains(notificationType , "sms"):
	// 	print("sms send notification")
	// case strings.Contains(notificationType , "email"):
	// 	print("email send notification")
	// default:
	// 	print("unkown notification type")
	// }

	//================================================================

	// for {  // بی نهایت
	// 	fmt.Println("Hello vahid!")
	// }

	//================================================================

	// for {
	// 	fmt.Println("Hello vahid!")
	// 	break
	// }

	//================================================================

	// list := []int{1,2,3,4,5,6,7,8,9,10,11,12,13}

	// for index, item :=range list {
	// 	fmt.Printf("index: %d , item: %v\n", index,item)
	// }

	//================================================================
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//================================================================

	// for i:= 0 ; i< 100 ;i++{
	// 	if i == 50 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	//==========================================

	// var array1 [4]int  // یک ارایه به ظول چهار
	// var array2 [3]int = [3]int {1,2,3}
	// array3 := [4]int {1,2}
	// array4 := [...]int {1,2,3,4,5}

	// array3[3] = 11

	// fmt.Println("array1 : " , array1)
	// fmt.Println("array2 : " , array2)
	// fmt.Println("array3 : " , array3)
	// fmt.Println("array4 : " , array4)

	//================================================================

	// names := [5]string {"vahid", "hasan","ali", "rahman"}

	// // fmt.Println("names : " , names)

	// names_key := "ali"

	// for i , name := range names {
	// 	if (names_key == name ){
	// 		println("found name from names list , index : " , i)
	// 		break
	// 	}
	// }
	// fmt.Println("not found name")

	//================================================================

	numbers1 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	numbers2 := &numbers1
	numbers3 := numbers1[:4]
	numbers4 := numbers1 // اگر مقدار دنامبر1 را تغییر دهیم مقدار نامبر 4 تغییر نمی کند

	numbers1[0] = 9999
	numbers2[1] = 8888

	fmt.Println("numers1: ", numbers1)
	fmt.Println("numers2: ", numbers2)
	fmt.Println("numers3: ", numbers3)
	fmt.Println("numers4: ", numbers4)

	changeValue(&numbers1) // حتما باید از امپرسان استفاده کنم

	fmt.Println("numbers1: ", numbers1)
	fmt.Println("numbers2: ", numbers2)
	fmt.Println("numbers3: ", numbers3)
	fmt.Println("numbers4: ", numbers4)

}

func changeValue(array *[8]int) { // اینجا باید از ستاره استفاده کنیم تا اشاره کنه به ارایه ما
	array[2] = 777
	array[3] = 666
}
