package main

import "fmt"

func main() {

	//ARRAY

	var intArr [4]int32

	var intArr1 [3] int32=[3] int32{1,2,3}

	var intArr2=[3] int32{2,3,4}

	fmt.Println(intArr1);
	fmt.Println(intArr2);

	intArr[1]=121
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0],&intArr[1],&intArr[2],&intArr[3]);

	//SLICE
	var intSlice []int32=[]int32 {4,5,6}

	fmt.Printf("The length is %v with capacity %v",len(intSlice),cap(intSlice))

	//this would create a new array with old element of array being copied to new array + the new element

	intSlice=append(intSlice, 3,4,5,6)

	fmt.Printf("\nThe length is %v with capacity %v",len(intSlice),cap(intSlice))

	var intSlice2 [] int32= [] int32 {6,7}

	intSlice=append(intSlice, intSlice2...)

	fmt.Println("\n",intSlice)

	var intSlice3 [] int32= make([]int32,3,9) //to make a slice with given length and capacity
	//if you kow what the final capacity would be then pass the capacity as it is faster in append
	fmt.Printf("The length is %v with capacity %v",len(intSlice3),cap(intSlice3))

//MAP

	var myMap=map[string] uint{"as":22,"xs":55,"ki":43}

	fmt.Println("\n", myMap)

	fmt.Println(myMap["xs"])

	var age,ok=myMap["ll"]

	if ok {
		fmt.Printf("The age is %v" , age)
	}else{
		fmt.Println("Invalid name");
	}

	//loops

	//There is no order in loops in maps any element could me printed in any iterartion 
	for name,age := range myMap{
		fmt.Printf("NAME:%v , AGE:%v \n",name,age)
	}
	//index , value
	//for map key,value
	for i,v:=range intArr{
		fmt.Printf("Index: %v, value: %v \n",i,v)
	}


	//There is no inbuild while loop in go so we can do this 

	var i int=0

	for i<10{
		fmt.Println(i)
		i+=1
	}
	//or we can do this 

	var i1=0
	for{
		if i1>=10{
			break
		}
		fmt.Println(i1)
		i1+=1

	}

	//or

	for i:=1; i<=11 ;i++{
		fmt.Println(i)
	}
}
