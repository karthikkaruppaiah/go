package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	if 1 == 1 {
		fmt.Println("1")
		fmt.Println(2)
	}
	Temp()
}

// Temp not sure why
func Temp() {
	FileName := "12"
	fmt.Println(FileName)
}
