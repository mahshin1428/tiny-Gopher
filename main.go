package main


import (
	"fmt"
	//"tiny-gopher/mathLib"
)


// execution order: global--> init() -> main() -> other functions
// functions types

//1.named  function
/*

func printHi(){
	fmt.Println("Hello")
}

// 2.init function
func init(){
	fmt.Println("Init function called")
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
*/



//closure function

const a = 10

var p = 20

func outer() func(){
	money := 100
	age := 30
	fmt.Println("age is", age)

	show:= func(){
		money = money + a + p
		fmt.Println("money is:", money)
	}
	return show // doubt silo function ki return korbe? (ans: Go-তে function নিজেও একটা value)
}

func call(){
	innerfunc := outer()
	innerfunc()   //130
	innerfunc()   //160
	innerfunc()   //190

	innerfunc2 := outer()
	innerfunc2()  //130
	innerfunc2()  //160
	innerfunc()   //220
}

func main(){
	call()
}

