package main 

func main(){
	print_1()
	println(print_2())
	println(print_3("hello vahid"))
	println(print_4("Hello"))

}

func print_1(){
	println("hello vahid")
}

func print_2() string{
	return	"hello vahid"
}

func print_3(Enter string) string{
	return Enter
}

func print_4(Enter string) string{
	return Enter + " " +"vahid"
}
