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

/*
//closure function

const a = 10

var p = 20

func outer() func() {
	money := 100
	age := 30
	fmt.Println("age is", age)

	show := func() {
		money = money + a + p // as money(a r p global tai eder niye tension nai) variable ta lagtese(as outter function dies) tai escape analysis kore heap er moddhe rakhe dibe
		fmt.Println("money is:", money)
	}
	return show // doubt silo function ki return korbe? (ans: Go-তে function নিজেও একটা value)
	//show er value reo heap e store kore rakhe
}

func call() {
	innerfunc := outer()
	innerfunc() //130
	innerfunc() //160
	innerfunc() //190
	/*f = {
	                              fn_ptr: &show_code,
	                               env_ptr: &{ money = 100 }
	                             }
								 eikhane akta variable er jonno different differnt stuct create korte
								 and protteke nijer stack joto bar call hoi shei onujayi update kortese


	innerfunc2 := outer()
	innerfunc2() //130
	innerfunc2() //160
	innerfunc()  //220  ei jonnei eikahne 220 dicche as agei 180 porjonto update kora silo
}

func main() {
	call()
}

*/

//struct

type User struct {
	Name string
	Age  int
}

func main() {
	//way-1
	// like var p int
	var user User

	//like p = 10
	user = User{ //instance or object of struct
		Name: "Mahshin",
		Age:  22,
	}
	fmt.Println("Name is:", user.Name)
	fmt.Println("Age is:", user.Age)

	//way-2
	user2 := User{
		Name: "Arafat",
		Age:  23,
	}
	fmt.Println("Name is:", user2.Name)
	fmt.Println("Age is:", user2.Age)
}
