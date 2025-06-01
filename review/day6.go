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

    The {{int8}} at the end indicates what I expect to be returned
    `)
	fmt.Println(`
    Let's test add with 2 + 5 
    fmt.Println(add(2,5)) = `, add(2, 5))

	fmt.Println(`
    also the same is the add2 function:
    add2(9,6)
    `, add2(9, 6))
	fmt.Println(`
    And that is functions! 

    Now we can move on toooo...


    ## Arrays 

    Arrays in go syntax: 

    var variableName = [numberOfValues]type(v,a,lu,es)
    `)
	arrayInt := [6]uint8{3, 8, 2, 5, 86, 3}
	fmt.Println(`
	arrayInt := [6]uint8{3, 8, 2, 5, 86, 3}
    `)
	fmt.Println(arrayInt)
	fmt.Println(`Type of arrayInt: `, reflect.TypeOf(arrayInt))

	fmt.Println(`
    ### assuming array size

    Arrays are generally explicitly defined with an amount of values but this can also be guessed rather than needed to be spelled out. Makes for more succint code 

    eg: 

    arrayAssumption := [...]int16{276, 214, 1000}

    The compiler will figure out that there are three values 
    `)

	fmt.Println(`
    ## Bundling

    There are a few rules to bundling files: 
    - files must have same package name at top of file 
    - file must be in same dir 

    ### Sharing variables/functions 
    - function/variable must begin with capital letter
    - function/varibale must be called as an attachment of the host package

    hello/hello.go 

  "package hello

    Hellovar := "hello"
    "

    main.go 
    "package main 

    import ("fmt", GoLang-Learning/hello)

    fmt.Println(hello.Hellovar)
    "

    Files within the same package, do not need to specify the package they are using to share variables
    hello.go 
    "package main

    Hellovar := "hello"
    "

    main.go 
    "package main 

    import "fmt"

    fmt.Println(Hellovar)
    "
    `)
}

func Day6Loops() {
	fmt.Println(`

    Loop time baybeee 

    ## Loops 


    In go, loops are done with the "for" keyword

    This keyword can be manipulated to do different loops 

    ### infinite loop 
    
    for {
    fmt.Println("I'll go on forever")
    }


    ### conditional loop 

    x := 20 
    for x >= 1 {
    fmt.Println(x)
    x -= 1 
    }


    ### classic loop 
    variable init ; variable condition ; loop action {} 

    for i := 2; i < 100; i += 9 {
      fmt.Print('omg I'm equal to:',i)
    }

    ### range loops

    Whenever a loop needs to go over all the values of an array or collection of some kind, a range loop is best 

    syntax and eg: 
    
    arrayEx := [...]string{"hello", "little", "bird"}
    for index, value := range arrayEx {
      fmt.Println("The value at index", index, "is:", value)
    } 

    ### using break & continue 

    For loops, break and continue will often be needed for when to take action 

    eg: 
    `)
}
