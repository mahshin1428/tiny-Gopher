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



//struct

type User struct {
	Name string
	Age  int
}


//Receiver function
func printUserDetails(u User) {
	fmt.Println("Name is:", u.Name)
	fmt.Println("Age is:", u.Age)
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
	//fmt.Println("Name is:", user.Name)
	//fmt.Println("Age is:", user.Age)

	//way-2
	user2 := User{
		Name: "Arafat",
		Age:  23,
	}
	//fmt.Println("Name is:", user2.Name)
	//fmt.Println("Age is:", user2.Age)
	printUserDetails(user)
	printUserDetails(user2)
}



// Array in go
func main() {
	//1.Fixed-size array
	var arr [2]int = [2]int{1, 2}
	//2.Fixed-size array with inferred length
	//b := [...]int{4, 5, 6, 7}
	//3. Partial initialization
	//4.Indexed initialization
	//5.Array of strings (or any type)

	x := 20 
	p := &x
	fmt.Println("Value of x:", *p) //20
	*p = 30
	fmt.Println("Value of x:", *p) //30

	fmt.Println("Array is:", arr) //by default 0
}



func main(){
	arr := [7] int{1,2,3,4,5,6,7}
	slice := arr[1:5]
	s := slice[1:3]   // length(number of element) = last index - first index 
	fmt.Println("Array is:", arr)
	fmt.Println("slice is",slice)
	fmt.Println("s is:", s)
	fmt.Println("slice lenth is:", len(slice))
	fmt.Println("slice capacity is:", cap(slice))
	fmt.Println("s length is:", len(s))
	fmt.Println("s capacity is:",cap(s))
	

}
	



//slice 	
func main(){
	var x[] int
	x = append(x, 1)  //eikhane basically j kaj ta hoitache ta hoile first e append function e jonno akta stack frame create hoitache then ager existing array(from heap memory) tare copy
	                  // append function e parameter hishebe antese then dakhe capcity koto jodi full hoi tahole notun akta array create kore then ager existing array er 
					  // element gulo copy kore new array te and new element ta add kore dibe plus pointer ta update korbe
					  // slice er size surute double(100%) kore barbe but 1024 er por 25% kore barbe
	x = append(x, 2)
	x = append(x, 3)

	y:= x
	
	x = append(x, 4) //age 4 number index e x 4 boshaisilo, but pore jokhhon y aki index e value(5) boshailo then ager value the overwrite hoise    
	y = append(y, 5)
	

	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

*/

//slice example

func changeslice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}
func main(){
	x:=[]int{1,2,3,4,5}
	x = append(x, 6)
	x= append(x, 7)

	a := x[4:]
	
	y := changeslice(a)

	fmt.Println("x:", x)
	fmt.Println("Y:", y)
}