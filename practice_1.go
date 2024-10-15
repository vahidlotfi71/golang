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

	// numbers1 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// numbers2 := &numbers1
	// numbers3 := numbers1[:4]
	// numbers4 := numbers1 // اگر مقدار دنامبر1 را تغییر دهیم مقدار نامبر 4 تغییر نمی کند

	// numbers1[0] = 9999
	// numbers2[1] = 8888

	// fmt.Println("numers1: ", numbers1)
	// fmt.Println("numers2: ", numbers2)
	// fmt.Println("numers3: ", numbers3)
	// fmt.Println("numers4: ", numbers4)

	// changeValue(&numbers1) // حتما باید از امپرسان استفاده کنم

	// fmt.Println("numbers1: ", numbers1)
	// fmt.Println("numbers2: ", numbers2)
	// fmt.Println("numbers3: ", numbers3)
	// fmt.Println("numbers4: ", numbers4)


//================================================================

	// myarray := []int{1, 2, 3, 4}

	// addNumber(&myarray , 12) // حتما باید از امپرسان استفاده کنیم و در خود تابع هم از  استار استفاده شود تا تغییرات روی تابع اعمال شود 
	// 						// ما وقتی یک اسلایس را می دهیم بع تابع تابع یک اسلایس جدید پشت برای خودش ایجاد می کند اگر ما از طریق اشاره گر به تابع یک اسلایس را پاس ندهیم تعییرات اعمال نمشود

	// fmt.Println(myarray)
//==========================================================================

// names := []string {"vahid", "hasan", "mohammad", "ali", "ahmad", "rahman"}

// for _ , name := range names {
// 	name = strings.ToUpper(name)
// // این فر روی اسلایس عمل نمی کند به این نکته توجه شود وقتی ما  وقتی در فر روی اسلایس کار می کنیم و از ایتم استفاده می کنیم این تغییرات اعمال نمی شود باید مستقیم نیمز را تغییر دهیم

// }

// fmt.Println(names)


// for i , name := range names{
// 	names[i] = strings.ToUpper(name)
// 	 // توجه شود وقتی از روی اسلایس کار می کنیم برای ایجاد تغییر مستقیم ار خود اسلایس استفاده کنیم
// 	 // در حلقه فر یک کپی از اسلایس ما گرفته می شود و تغییر روی اونن کپی اعمال می شود ایتم یک کپی از اون است

// }

// fmt.Println(names)

//================================================================

// list_vahid := []int {1,2,3,4,5,6,7,8,}
// list_hasan := []int {11,22,23,24,66,66,66,66,66,22,33,44,55,66}

// cambinate_list := copy(list_vahid, list_hasan) // این تابع ورودی اول مقصد و ورودی دوم منبع است و دو تا ورودی باید هم نوع باشن
// 												// این تابع دادهای منبع را در مقصد می ریزد بعد ریختن اگر ظرفت داشت داده های مقصد رامی اورد 
// 												// ظریفت تابع مقصد به اندازه خودش است اضافه نمیتوانیم واردش کنیم

// fmt.Println(cambinate_list) // تعداد ارایه های کپی شده را نشان می دهد
// fmt.Println(list_vahid) 

//================================================================

// number_1 := []int{1, 2, 3, 4, 5}
// number_2 := []int{10,11,12,13,14}
// number_3 := []int{20,21,22,23,24}

// number_1 = append(number_1 , 6,7,8,9)
// number_2 = append(number_2 , number_1...)
// number_3 = append(number_3,number_2... )
// 								//حتما باید سه نقطه گذاشته شود وگر نه خطا می ده 

// fmt.Println(number_1)
// fmt.Println(number_2)
// fmt.Println(number_3)

//================================================================
//    // تمرین =>  بین عنصر 7 و 8 می خواهم 75 را بگذارم
// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// // یک صفر به اخر اضافه می کنیم وقتی شیف پیدا کردن جلو فضا داشته باشن
// numbers = append(numbers,0)

// _ = copy(numbers[7:],numbers[6:])
// numbers[7] = 75
// fmt.Println(numbers)

//================================================================
// 	// می خواهیم 75 را حذف کنم 

// numbers := []int{1, 2, 3, 4, 5, 6, 7,75, 8, 9}


// numbers = append(numbers[:7],numbers[8:]...)
// fmt.Println(numbers)

//================================================================
// خالی کردن اسلایس

// number := []int{1,2,3,4,5,6,7}
// number = number[:0]
// fmt.Println(number)
// fmt.Println("capacity:" , cap(number)) // هنوز کپ اسلایس اسست

// fmt.Println("--------------------------------")

// number = nil

// fmt.Println(number)
// fmt.Println("capacity:" , cap(number)) 

//================================================================

// names := make(map[string]string)
// names_1 := map[int]string{}
// var names_2 map[int]string = map[int]string{}
// names["121212"] = "vahid"
// names_1[2222] = "hasan"
// names_2[11111] = "ali"

// fmt.Println(names)
// fmt.Println(names_1)
// fmt.Println(names_2)

//================================================================

// type Person struct {
// 	name string
// 	family string
// 	age int 
// }

// persons := make(map[int]Person)

// persons[12] = Person{"vahid", "lotfi", 32}
// persons[13] = Person{"hasan", "mehtari",28}
// persons[14] = Person{"reza", "mesgar", 25}
// persons[15] = Person{"hadi", "farahani",28}

// persons[13] = Person{"Ali", "lotfi",2}  // اپدیت 

// delete(persons , 15)  // تابع دیلیت مپ و کلید من را می گیرد و اون ردیف راحذف می کند

// for _ , person := range persons {
// 	fmt.Println(person)

// }

// vahid := persons[12]  //  می خواهیم چک کنیم ببینم این کد هست یا نه
// hasan := persons[20] // کد نباشد صفر برمی گرداند 


// fmt.Println(vahid)
// fmt.Println(hasan)

// vahid1  := persons[12]
// if true {
// 	fmt.Println(vahid1)
// }else {
// 	println("not found")
// }
	
//================================================================

type Person struct {
	name string
	family string
	age int
}

persons := map[int]Person{}
mySlice := []int{}

persons[1]= Person{"vahid", "lotfi", 32}
mySlice = append(mySlice , 1)
persons[2]= Person{"hsasn", "mehtari", 28}
mySlice = append(mySlice ,2)
persons[3]= Person{"ali", "ahmadi", 22}
mySlice = append(mySlice ,3)
persons[4]= Person{"rahim", "nazari",12}
mySlice = append(mySlice ,4)


for _ , person := range mySlice {
	fmt.Println(persons[person])
}



}






// func changeValue(array *[8]int) { // اینجا باید از ستاره استفاده کنیم تا اشاره کنه به ارایه ما
// 	array[2] = 777
// 	array[3] = 666
// }


// func addNumber(numbers *[]int , number int ){
// 	*numbers = append(*numbers,number)
// }
