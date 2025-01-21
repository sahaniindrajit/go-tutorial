package main

import "fmt"

//struct: to define user defined type
type gasEngine struct{
	kml uint8
	litter uint8
	
}

type Electric struct{
	kmkwh uint8
	kmh uint8
}

type engine interface{
	kmLeft() uint8
}

func (e gasEngine) kmLeft() uint8{
	return e.litter*e.kml
}

func (e Electric) kmLeft() uint8{
	return e.kmh*e.kmkwh
}

func canMakeit(e engine,km uint8){
	if km<=e.kmLeft(){
		fmt.Println("you can make it")
	}else{
		fmt.Println("you can't make it")
	}
}

func main (){

	var myEngine gasEngine=gasEngine{20,30};
	myEngine.litter=10
	fmt.Println(myEngine.kml,myEngine.litter)

	canMakeit(myEngine,250)

}