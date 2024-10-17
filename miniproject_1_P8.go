package main

import (
	"fmt"
)

type Room struct {
	ID       int
	Type     string
	BedCount int
	Status   bool
	Price    int
}

var Rooms []Room = GenerateRooms()


func main() {
	input :=""

	for input != "exit"{
		fmt.Println("Enter Command : ")
		fmt.Println("1 : Room list")
		fmt.Println("2 : Add room")
		fmt.Println("3 : reserve room")

		fmt.Scan(&input)

		switch input {
		case "1":
			GetRoomList()
		case "2":
			AddRoom()
		case "3":
			ReserveRoom()
		case "Exit":
			println("Exiting ...")
			break
		default :
				println("invalid command")
	}
}
}



func GetRoomList(){
	for _ , room :=range Rooms {
		fmt.Printf("%+v\n", room)
	}
}



func GetRoomFromInput() Room {
	var room Room = Room{Status : false}
	fmt.Println("Enter Room information line by line (ID , Type ,BedCount ,Price")
	fmt.Scan(&room.ID)
	fmt.Scan(&room.Type)
	fmt.Scan(&room.BedCount)
	fmt.Scan(&room.Price)

	return room

}



func AddRoom() {
	room := GetRoomFromInput()
	Rooms = append(Rooms, room)
}



func ReserveRoom(){
	id := 0
	nights := 0
	personCount := 0

	fmt.Println("Enter room ID for reservation: ")
	fmt.Scan(&id)

	room := GetRoom(id)
	if room == nil {
		fmt.Println("Room not found")
		return
	}
	if room.Status == true {
		fmt.Println("Room already reserved")
		return
	}

	fmt.Println("Enter reserve information line by line (nights and personsCount)")
	fmt.Scan(&nights)
	fmt.Scan(&personCount)

	roomPrice, totalPrice  , tax , discoutAmount := calculateRoomPrice(*room , nights ,personCount)
	room.Status = true

	fmt.Printf("RoomPrice: %f, TotalPrice: %f, tax: %f,discoutAmount: %f\n", roomPrice, totalPrice, tax , discoutAmount)
	
}	




func calculateRoomPrice(room Room  , nights int , personCount int )(roomPrice float64 , totalPrice float64  , tax float64 ,discountAmount float64 ){
	discountPercentage := 0.0
	if (nights >= 7 && nights <=15){
		discountPercentage = 0.10
	}else if (nights > 15 && nights <= 20){
		discountPercentage = 0.15
	}else if (nights >21){
		discountPercentage = 0.20}
	
	switch room.Type{
	case "single":
		roomPrice = float64(nights * personCount *room.Price)

	case "double":
		roomPrice = float64(nights * personCount * room.Price) 

	case "suite":
		roomPrice = float64(nights * personCount * room.Price)
		
	}
	
	tax = roomPrice * .9
	discountAmount = roomPrice * discountPercentage
	totalPrice = roomPrice + tax - discountAmount

	return

}



func GetRoom(id int) *Room {
	for i:= 0 ; i <=len(Rooms) ; i++ {
		if Rooms[i].ID == id {
			return &Rooms[i]	
		}
	}
	return nil	
}



func GenerateRooms() []Room {
	rooms := []Room{}

	rooms = append(rooms, Room{ID: 1, Type: "single", BedCount: 1, Status: false, Price: 100})
	rooms = append(rooms, Room{ID: 2, Type: "single", BedCount: 2, Status: false, Price: 200})
	rooms = append(rooms, Room{ID: 3, Type: "single", BedCount: 1, Status: false, Price: 100})
	rooms = append(rooms, Room{ID: 4, Type: "double", BedCount: 2, Status: false, Price: 300})
	rooms = append(rooms, Room{ID: 5, Type: "double", BedCount: 3, Status: false, Price: 400})
	rooms = append(rooms, Room{ID: 6, Type: "double", BedCount: 4, Status: false, Price: 500})
	rooms = append(rooms, Room{ID: 7, Type: "suite", BedCount: 4, Status: false, Price: 700})
	rooms = append(rooms, Room{ID: 8, Type: "suite", BedCount: 5, Status: false, Price: 800})
	rooms = append(rooms, Room{ID: 9, Type: "suite", BedCount: 5, Status: false, Price: 900})
	rooms = append(rooms, Room{ID: 10, Type: "suite", BedCount: 6, Status: false, Price: 1000})

	return rooms
}
