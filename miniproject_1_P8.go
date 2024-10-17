package main 

import (
	"fmt"
)

type Room struct {  
	ID int 
	Type string
	Status bool  // true for reserved and false for available
	BedCount int
	Price int
}

var Rooms []Room = GenerateRooms()  // list of rooms

func main() {
	input := "" // گرفتن ورودی

	for input != "exit"{
		fmt.Println("Enter a command: ")
		fmt.Println("1 : Room list")
		fmt.Println("2 : Add Room")
		fmt.Println("3 : reserve Room")

		fmt.Scan(&input)

		switch input {
		case "1" :
			GetRoomList()
		case "2" :
			AddRoom()
		case "3" :
			ReserveRoom()
		case "exit" :
			fmt.Println("Exiting ....")
			break
		default :
			fmt.Println("Invalid command")

		}



	}
}

func GetRoomList(){
	for _ , room := range Rooms{
		fmt.Printf("%+v\n", room)  // بعلاوه وی یعنی همه یاطلاعات ان را نمایش بده 
	}
}

func GetRoomFromInput() Room{  // خروجی این تابع یک رووح است
	var room Room = Room{Status: false} 
	fmt.Println("Enter room information line by line : (ID , Type , BedCount , Price)")

	fmt.Scan(&room.ID)
	fmt.Scan(&room.Type)
	fmt.Scan(&room.BedCount)
	fmt.Scan(&room.Price) 

	return room

}

func AddRoom(){
	room := GetRoomFromInput()
	Rooms = append(Rooms, room)
}

func ReserveRoom(id int , night int , personCount int ){ // وقتی می خواهیم یک روم را رزرو کنیم نیاز داریم بکسری اطلاعات را از کاربر بگیریم 
	room := GetRoom(id)

	if room == nil {
		fmt.Println("No room found")
		return
}

func GetRoom (id int) *Room {
	for i := 0 ; i <= len(Rooms) ; i++{
		if Rooms[i].ID == id {
			return &Rooms[i]
		}
	}
	return nil
}

func CalculateRoomPrice(){}

func GenerateRooms() []Room {  //   اینجا اتاق هامون را می سازیم ویک اسلایس از جنس روم را بر میگرداند
	rooms := []Room{} // رومز یک اسلایس است که اینجا ایجاد کردیم

	rooms = append(rooms,Room{ID: 1 , Type: "single" , Status: false, BedCount: 2 , Price: 100})
	rooms = append(rooms,Room{ID: 2 , Type: "single" , Status: false, BedCount: 3 , Price: 200})
	rooms = append(rooms,Room{ID: 3 , Type: "single" , Status: false, BedCount: 4  , Price: 250})
	rooms = append(rooms,Room{ID: 4 , Type: "duble" , Status: true, BedCount: 2 , Price: 300})
	rooms = append(rooms,Room{ID: 5 , Type: "duble" , Status: false, BedCount: 3 , Price: 400})
	rooms = append(rooms,Room{ID: 6 , Type: "duble" , Status: false, BedCount: 4  , Price: 450})
	rooms = append(rooms,Room{ID: 7 , Type: "suite" , Status: true, BedCount: 2 , Price: 500})
	rooms = append(rooms,Room{ID: 8 , Type: "suite" , Status: false, BedCount: 3 , Price: 600})
	rooms = append(rooms,Room{ID: 9 , Type: "suite" , Status: false, BedCount: 4  , Price: 750})
 
	return rooms
}



