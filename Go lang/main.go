package main

import(
 "fmt"
 "math/Sum"
)

 func getNumbers(num1 int, num2 int)(int, int){
	mod:= num1 * num2
	dev:= num1 / num2

	return mod, dev
 }

func main(){
	a := 20
	b := 30
	mod, dev := getNumbers(a, b)
	fmt.Println(mod)
	fmt.Println(dev)
	res := Sum.Sum(a,b)
	fmt.Println(res)

}