package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("hiii")
	var num1 int=7
	var num2 int = 0

	var divisionResult,divisionRemender,err =intDivision(num1,num2);

	if err != nil{
		fmt.Println(err.Error())
	}

	fmt.Printf("The result is %v with remender %v",divisionResult,divisionRemender)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num1 int,num2 int) (int,int,error){
	var err error

	if(num2==0){
		err=errors.New("can't divide by zero")
		return 0,0,err
	}
	var result int=num1/num2;
	var remender int=num1%num2
	return result,remender,err
}