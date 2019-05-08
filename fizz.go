package fizzbuzz

import (
	"strconv"
	"strings"
)

func Fizzgame(fizzNum, buzzNum, i int) (ret interface{}) {

	const FIZZBUZZ = "fizzbuzz"
	const FIZZ = "fizz"
	const BUZZ = "buzz"

	if DevidedByNum(i, fizzNum) && DevidedByNum(i, buzzNum) {
		return FIZZBUZZ
	}
	if DevidedByNum(i, fizzNum) ||Contains(i, fizzNum){
		return FIZZ
	}

	if DevidedByNum(i, buzzNum) ||Contains(i, buzzNum){
		return BUZZ
	}

	return i
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



