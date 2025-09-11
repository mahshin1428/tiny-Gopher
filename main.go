package main

import (
	"fmt"
	"tiny-gopher/mathLib"
)

func printHi(){
	fmt.Println("Hello")
}

func main(){
	printHi()
	result := mathLib.Sum(2,4)
	fmt.Println("Sum is:", result)
}