package main

import (
	"fmt"
	"fizzbuzz/fizzlib"
)


func main(){
	var firstNum, secondNum int

	fmt.Println("Game Start...")
	fmt.Println("Now input 2 number that less than 10:")

	fmt.Scanln(&firstNum, &secondNum)
	c := &fizzlib.Context{firstNum, secondNum}

	for i:=1; i<=100; i++ {
		fmt.Println(c.Fizzgame(i))
	}
}


