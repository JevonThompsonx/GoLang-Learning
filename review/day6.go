package review

import (
	"fmt"
	"reflect"
	"strconv"
)

func Day6() {
	fmt.Println(`
    Starting day Day6

    # Let's review 

    creating
    ## variables 

    var, const or := 

    var variableName type = value
    eg:
    
    var variable1 string = "hello world"

    const variableName type = value 
    eg: 
    const variable2 int8 = 78

    variableName := value 
    eg: 
    variable3 := 54 

    Those are the different ways that a variable can be created. 
    Now let's specify the the types 

    ## Basic types: 

    string - "within double qoutes"

    int - numbers 
	fmt.Println("randomNumba type = ", reflect.TypeOf(randomNumba))
    float - decimal point numbers 

    bool - true/false binary 
  
    `)

	boolTest := false
	fmt.Println(boolTest)
	fmt.Println(`
      
    ### Basic types broken down 

    #### int - numbers
    int8, int16, int32 etc...

    #### uint - positive numbers only
    uint8, uint16, uint32 etc...

    #### float 
    float32,float64 
    
    ## Type conversions 
    ### Numebers - int/uint/floats
    For conversions between numbers, it's not too hard. It's done by using the desired variable type and the variable 
    
    eg: 
    var variableUint uint8 = 98
    var variableFloat = float32(variableUint)

    hint: the reflect package can be used to check type
    `)
	var variableUint uint8 = 98
	variableFloat := float32(variableUint)

	fmt.Println(`The type of variableUint is:`, reflect.TypeOf(variableUint))
	fmt.Println(`The type of variableFloat is:`, reflect.TypeOf(variableFloat))

	fmt.Println(`
    Let's go over converions between numbers and strings 

    ### Strings := int 

    The strconvert package 

    strconvert.Atoi - string to int
    strconvert.Itoa - int to string 

    variableInt := 23
    variableStr := "hola"


	newString := strconv.Itoa(variableInt)
    `)

	variableInt := 23
	variableStr := "hello"

	newString := strconv.Itoa(variableInt)
	newInt, err := strconv.Atoi(variableStr)

	if err != nil {
		fmt.Println("Conversion failed:", err)
	} else {
		fmt.Println("Converted number:", newInt)
	}
	fmt.Println(`
    Type of newString
    `, reflect.TypeOf(newString))
	fmt.Println(`
    Type of newInt
    `, reflect.TypeOf(newInt))

	fmt.Println(`
    ## functions 

    I've already used them a bit but the basic syntax is: 

    func functionName() {
      function 
        action
          here
    }
    `)
}

func add(x int8, y int8) int8 {
	result := x + y
	return result
}

func add2(x, y int8) int8 {
	result := x + y
	return result
}

func Day6P2() {
	fmt.Println(`
    Back to functions. Let's use this function

func add(x int8, y int8) int8 {
	result := x + y
	return result
}

    To do some simple addition. It works by taking values as x and y 
      It expects x and y to be int8 - this can be rewritten to: 
func add(x , y int8) int8 {
	result := x + y
	return result
}

    The 
    `)
}
