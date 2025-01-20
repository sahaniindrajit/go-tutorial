package main

import "fmt"

func main() {
	fmt.Println("Hello world from golang!")

	var intNum int =4

	fmt.Println(intNum);

	var floatNum float64=2222.22
	fmt.Println(floatNum)

	//two diffrent types cant be added 

	var result float32= float32(floatNum)+float32(intNum);

	fmt.Println(result)

	//integer dividsion results in int and result is rounded down.
	var int1 int=5
	var int2 int=4
	fmt.Println(int1/int2)
	fmt.Println(int1%int2)

	var myString1 string ="Hello World"

	fmt.Println(myString1)

	var myString2 string =`Hello 
	World`

	fmt.Println(myString2)

	var myvar ="text" //type is inferred automattically
	fmt.Println(myvar)

	myvar1:= "text1"
	fmt.Println(myvar1)

	var var2,var3,var4=2,3,4 //or we can do var2,var3,var4:=2,3,4
	fmt.Println(var2,var3,var4)

	const myConst string ="const value" 
	fmt.Println(myConst)
}