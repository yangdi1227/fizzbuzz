package fizzbuzz

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_input1_return1(t *testing.T) {
	ret := Fizzgame(3,5,1)
	assert.Equal(t, 1, ret)

}

func Test_input2_return2(t *testing.T) {
	ret := Fizzgame(3,5,2	)
	assert.Equal(t, 2, ret)

}
func Test_input3_returnfizz(t *testing.T) {
	ret := Fizzgame(3,5,3)
	assert.Equal(t, "fizz", ret)

}
func Test_input5_returnbuzz(t *testing.T) {
	ret := Fizzgame(3,5,5)
	assert.Equal(t, "buzz", ret)

}
func Test_input6_returnfizz(t *testing.T) {
	ret := Fizzgame(3,5,6)
	assert.Equal(t, "fizz", ret)

}
func Test_input10_returnbuzz(t *testing.T) {
	ret := Fizzgame(3,5,10)
	assert.Equal(t, "buzz", ret)

}
func Test_input15_returnfizzbuzz(t *testing.T) {
	ret := Fizzgame(3,5,15)
	assert.Equal(t, "fizzbuzz", ret)

}
func Test_input13_returnfizz(t *testing.T) {
	ret := Fizzgame(3,5,13)
	assert.Equal(t, "fizz", ret)

}
func Test_input35_returnbuzz(t *testing.T) {
	ret := Fizzgame(3,5,35)
	assert.Equal(t, "fizz", ret)

}
func Test_input53_returnbuzz(t *testing.T) {
	ret := Fizzgame(3,5,53)
	assert.Equal(t, "fizz", ret)

}
func Test_input52_returnbuzz(t *testing.T) {
	ret := Fizzgame(3,5,52)
	assert.Equal(t, "buzz", ret)

}
func Test_input25_returnbuzz(t *testing.T) {
	ret := Fizzgame(3,5,25)
	assert.Equal(t, "buzz", ret)

}

