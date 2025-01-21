package main

import (
	"fmt"
	"strings"
)

func main() {

	var myString1="rēsusē"
	fmt.Println(myString1)

	var indexed=myString1[0]

	//%T for type
	fmt.Printf("%v,%T \n",indexed,indexed)

	for i,v :=range myString1{
		fmt.Println(i,v)
	}

	//if we take length of string it is the number of bytes not character

	fmt.Printf("\n The length of myString1 is %v",len(myString1))
	
	var strSlince = []string{"s","d","f","g","t"};

	var catstr="fuck"

	for i:=range strSlince{
		catstr+=strSlince[i]
	}

	fmt.Printf("\n %v",catstr)

	//strings are immutable or this meathod is pretty fast

	var strBuilder strings.Builder
	for i:=range strSlince{
		strBuilder.WriteString(strSlince[i])
	}

	var catStr2=strBuilder.String()

	fmt.Printf("\n%v",catStr2)
}