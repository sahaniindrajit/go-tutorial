package main

import "fmt"

func main() {
	var i int32=7
	//pointer
	var p *int32 = new(int32)
	*p=9
	fmt.Printf("The value p pointer to is: %v", *p)

	fmt.Printf("\nThe value of i is : %v ",i)

	//here we set the p to the address of i
	p=&i

	fmt.Printf("\nThe value p pointer to is: %v", *p)

	


}