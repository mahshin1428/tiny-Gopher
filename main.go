package main


import (
	"fmt"
	"tiny-gopher/mathLib"
)


// execution order: global--> init() -> main() -> other functions
// functions types

//1.named  function
func printHi(){
	fmt.Println("Hello")
}

func main(){
	printHi()
	result := mathLib.Sum(2,4)
	fmt.Println("Sum is:", result)

	//3. annonymous function
	func(a int, b int){
		c:= a + b
		fmt.Println("Sum from annonymous function:", c)
	}(3,5)  //immediatly invoke function expression(IIFE)
}

// 2.init function
func init(){
	fmt.Println("Init function called")
}

