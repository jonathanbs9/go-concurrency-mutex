package main

import (
	"fmt"
	"sync"
)
var (
	mutex sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup){
	mutex.Lock()
	fmt.Printf("Depositando a la cuenta : %d , con balance : %d \n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup){
	mutex.Lock()
	fmt.Printf( "Retirando de la cuenta: %d , con balance: %d \n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func main(){
	balance  = 1000
	var  wg sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)

	wg.Wait()

	fmt.Printf("El nuevo balance de la cuenta es : %d \n ", balance)
}
