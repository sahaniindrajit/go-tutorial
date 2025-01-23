package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"a", "s", "d", "f", "g", "h"}

var results=[]string{}

var wg =sync.WaitGroup{}

var m=sync.RWMutex{}
//normal way of doing in inefficient

    /*func main() {

	t0 := time.Now()

    for i:=0;i<len(dbData);i++{
        dbCall(i)
    }

    fmt.Printf("\n Total exectution time : %v",time.Since(t0))
}

func dbCall(i int){

    var delay float32 =rand.Float32()*2000

    time.Sleep(time.Duration(delay)*time.Millisecond)

    fmt.Println("The result from the data base is: ",dbData[i])
}
    */

//Doing it concurrently

func main() {

    t0 := time.Now()

    for i:=0;i<len(dbData);i++{
        wg.Add(1)
        go dbCall(i)
    }
    wg.Wait()
    fmt.Printf("\nTotal exectution time : %v",time.Since(t0));
    fmt.Printf("\nThe result are %v",results)
}

func dbCall(i int){

    var delay float32 =2000

    time.Sleep(time.Duration(delay)*time.Millisecond)

    fmt.Println("\nThe result from the data base is: ",dbData[i])

    save(dbData[i])
    log()
    wg.Done()
    
}

//m.lock block both read and write if it is active

func save(result string){
    m.Lock()
    results=append(results,result)
    m.Unlock()
}

//m.rlock alows reads and only block write if active
func log(){
    m.RLock()
    fmt.Printf("\nThe current result are %v",results)
    m.RUnlock()
}
