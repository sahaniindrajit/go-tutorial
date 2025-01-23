package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICK_PRICE float32=5

func main() {
	var chickChannel = make(chan string)
	var websites=[]string{"ac.com","gu.in","rt.co"}

	for i:=range websites{
		go checkChickPrice(websites[i],chickChannel)
	}

	sendMessage(chickChannel)
}

func checkChickPrice(websites string,chickChannel chan string) {
	for{
		time.Sleep(time.Second*1)
		var chickPrice=rand.Float32()*20
		if chickPrice<=MAX_CHICK_PRICE{
			chickChannel<-websites
			break
		}
	}
}

func sendMessage(chickChannel chan string){
	fmt.Printf("\nFound a deal on chicken at %v",<-chickChannel)
}