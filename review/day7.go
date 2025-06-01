package review

import (
	"fmt"
	"reflect"
	"strconv"
)

func Day7Func() {
	fmt.Println("Yay a function")
}

func Day7Quick() {
	fmt.Println("This day 6 is fresh, this review will be a quick one")
	fmt.Println("## Basic types")
	var variable1 string = "hello world"
	const variable2 int = 89
	assumedTypeVar := 8.8

	fmt.Println(`### Types expanded
  
  int - basic number type - positive and negative 
    int8
    int16
    int32 
    etc 
  uint - number type - positive and negative
    uint8
    uint16
    uint32 
    etc 
  float - number type with decimal point 
    float32
    float64 
  boolean - binary - true/false 
    `)
	fmt.Println(`Checking types...`,
		reflect.TypeOf(variable1),
		reflect.TypeOf(variable2),
		reflect.TypeOf(assumedTypeVar),
	)

	fmt.Println(` ## conversions`)
	convertToString := strconv.Itoa(variable2) // returns 1 value because converting from an int to a string never fails
	convertMePlease := "90"
	convertToInt, _ := strconv.Atoi(convertMePlease) // returns 2 values because converting to an int can fail
	assumedFloatToInt := uint8(assumedTypeVar)       // converting numbers from float to int
	fmt.Println(``, convertToString, convertToInt, assumedFloatToInt)
	intToFloatVar := float32(variable2) // converting numbers from int to float
	fmt.Println(``, intToFloatVar)
	fmt.Println(`## Advanced types (so far)
  arrays - collection of values 
    syntax: variableName := [indexAmount]type{va,l,u,e,s}
    `)

	arrayVar := [3]uint8{27, 201, 6}
	fmt.Println(arrayVar)
	arrayVarAssumedIndex := [...]float32{562.2323, 543.2323, 2545.12}
	fmt.Println(arrayVarAssumedIndex)
	fmt.Println("## Functions")
	Day7Func()

	fmt.Println("## Loopdi loops")
	x := 21
	for x > 90 {
		x += 7
	}
	for index, value := range arrayVar {
		fmt.Println(` arrarVar index: [`, index, `]`, `value: `, value)
	}
	// for {
	// fmt.Println("Forever func")
	// }
}
