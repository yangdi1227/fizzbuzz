package fizzlib

import (
	"strconv"
	"strings"
)

type Context struct {
	FizzNum int
	BuzzNum int
}
func (c *Context)Fizzgame(i int) (ret interface{}) {

	const FIZZ = "fizz"
	const BUZZ = "buzz"
	var result string

	if DevidedByNum(i, c.FizzNum) ||Contains(i, c.FizzNum){

		result += FIZZ
	}

	if DevidedByNum(i, c.BuzzNum) ||Contains(i,c.BuzzNum){
		result += BUZZ
	}

	return RecombinteResult(result, i)
}

func Contains(i , num int) bool {
	return strings.Contains(IntToString(i), IntToString(num))
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func DevidedByNum(i int, num int) bool {
	return i%num == 0
}

func RecombinteResult(str string, i int)(ret interface{}){
	if str == ""{
		return i
	}
	return str
}


