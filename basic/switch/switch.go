package main

import "fmt"

func main() {
	ChooseFruit("banana")
	ChooseFruit("apple")
	ChooseFruit("test")
}

func ChooseFruit(fruit string) {
	switch fruit {
	case "apple":
		fmt.Println("This is an apple")
	case "banana":
		fmt.Println("This is a banana")
	default:
		fmt.Printf("Don't know what it is: %s \n", fruit)
	}
}
