package main

import "fmt"

func GetType(variable any) {
	switch variable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	ch := make(chan int)
	GetType(1)
	GetType("hello world")
	GetType(false)
	GetType(ch)
	GetType([]int{})
}
