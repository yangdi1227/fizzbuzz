package fizzlib

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
)

type FizzBuzzTestSuite struct {
	suite.Suite
	c *Context
}

func (suite *FizzBuzzTestSuite) SetupSuite(){
	suite.c = &Context{3,5}
}


func (suite *FizzBuzzTestSuite)Test_input1_return1() {
	ret := suite.c.Fizzgame(1)
	assert.Equal(suite.T(), 1, ret)

}

func (suite *FizzBuzzTestSuite)Test_input2_return2() {
	ret := suite.c.Fizzgame(2)
	assert.Equal(suite.T(), 2, ret)

}
func (suite *FizzBuzzTestSuite)Test_input3_returnfizz() {
	ret := suite.c.Fizzgame(3)
	assert.Equal(suite.T(), "fizz", ret)

}
func (suite *FizzBuzzTestSuite)Test_input5_returnbuzz() {
	ret := suite.c.Fizzgame(5)
	assert.Equal(suite.T(), "buzz", ret)

}
func (suite *FizzBuzzTestSuite)Test_input6_returnfizz() {
	ret := suite.c.Fizzgame(6)
	assert.Equal(suite.T(), "fizz", ret)

}
func (suite *FizzBuzzTestSuite)Test_input10_returnbuzz() {
	ret := suite.c.Fizzgame(10)
	assert.Equal(suite.T(), "buzz", ret)

}
func (suite *FizzBuzzTestSuite)Test_input15_returnfizzbuzz() {
	ret := suite.c.Fizzgame(15)
	assert.Equal(suite.T(), "fizzbuzz", ret)

}
func (suite *FizzBuzzTestSuite)Test_input13_returnfizz() {
	ret := suite.c.Fizzgame(13)
	assert.Equal(suite.T(), "fizz", ret)

}
func (suite *FizzBuzzTestSuite)Test_input35_returnbuzz() {
	ret := suite.c.Fizzgame(35)
	assert.Equal(suite.T(), "fizzbuzz", ret)

}
func (suite *FizzBuzzTestSuite)Test_input53_returnbuzz() {
	ret := suite.c.Fizzgame(53)
	assert.Equal(suite.T(), "fizzbuzz", ret)

}
func (suite *FizzBuzzTestSuite)Test_input52_returnbuzz() {
	ret := suite.c.Fizzgame(52)
	assert.Equal(suite.T(), "buzz", ret)

}
func (suite *FizzBuzzTestSuite)Test_input25_returnbuzz() {
	ret := suite.c.Fizzgame(25)
	assert.Equal(suite.T(), "buzz", ret)

}

func TestFizzBuzzTestSuite(t *testing.T) {
	suite.Run(t, new(FizzBuzzTestSuite))
}

