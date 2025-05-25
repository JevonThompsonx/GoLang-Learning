package review

import (
	"fmt"
	"reflect"
)

func Day5() {
	fmt.Println(`I've seen this popup a few times so I should add it to what I know.

    Let's learn about go arrays/slices next time a well as do normal review

    From what I can tell so far, the basic structure is: 
    variableName := []arrayType {

    arrayContents
    arrayContentes
    }

    egs 
    
    numberArray := []int {64,1212,65}
    
    `)
	numberArray := []uint8{64, 121, 65}
	fmt.Println(numberArray)
	Println(`
    Or 

	stringSlice := []string{"hello", "I string"}
    `)
	stringSlice := []string{"hello", "I string"}
	fmt.Println(stringSlice)
	Println(`Anyways. more next time

  The list is 
    - Learn about slices
    - Dive more into testing
    - typechanging

    `)

	Println(`
    Alright, it's review time
    
    # variables

    Variables can be created using the var or const keyword 

    var variableName string = "hello I is string"

    var variableNam type    =  "value here"


    const variableName2 int8 = 89
    const variableName  type = value

    as a shortcut to making the type implicit, use := the walrus operator

    variableName3 := 90

    this creates a regular variable, not a const variable 

    var - regular variables that can be changed 
      - using explicit/implicit types means that the 'type' of variable can't be change but the value can 

    const - constant variables that cannot be change
      - the value of this never changes 
      - ever 

    := functions like using var 

    # Importing/exporting 

    In go, variables are imported or exported using packages and capitalization 

    All files require a "package packageName" at the beginning that groups files in the same folder 

    All files in the same package can share variables if: 
      - the variables is capitalized
      - they are in the same parent folder 

    eg: 

    "review.go" file 
    package main 

    import "fmt"

    func Review() {
      fmt.Println("Hello from review")
    }

    "main.go" file

    package main

    import  "parentFolderName/review"
    func main(){
      review.Review()
    }

    since we've used a function, let's go over it 

    # functions 

    Functions are created using the func keyword 

    func functionName() {
    what
    the
    function 
    does
    here
    }

    it's that simple 

    Let's go over basic operators 

    - subtract 
    + add 
    / division 
    * multiplication

    `)
	fmt.Println(`2 + 2 = `, 2+2)
	fmt.Println(`9 * 5 = `, 9*5)
	Println(`
  Thee is also the modulo operator to find the remainder of a division

    `)
	fmt.Println(`9 % 4 = `, 9%4)
	Println(`Let's use that to make a even/odd tester`)
}

func Day5evenOdd() {
	fmt.Println("100%4=", evenOdd(100, 4))
	fmt.Println("100/4 = ", 100/4)
}

func evenOdd(number1, number2 int8) int8 {
	return number1 % number2
}

func Day5P2() {
	Println(`
    Feeling pretty comfy with functions so let's move on

    Next is type specifity. 

    Although it isn't necessary, specifying the types can save memory

    eg 

    int is a broad number type, int8 is a specific number type

    int vs int8, int16, int32 etc...
    uint vs uint8, uint16, uint32 etc...
    string is just string
    bool is just bool 
    float vs float32, float64

    I've already been over these to the point of them being really in my head so I'll move on

    Most important to remember is that int is all numbers and uint is positive only. 
    Not most important but cool because I didn't know that was a thing
    
    ## type conversion 

    Since we're doing types, let's do type conversions 


  for numbers, type conversion is as simple as wrapping the value in the type
  variableInt := 89
  variableFloat := float32(89)
  
    We can check the type using reflect:
      reflect.TypeOf(variableInt)
    `)
	variableInt := 89
	variableFloat := float32(variableInt)
	fmt.Println(`type of variableInt = `, reflect.TypeOf(variableInt))
	fmt.Println(`type of variableFloat = `, reflect.TypeOf(variableFloat))
}

func Day5ArraysSlices() {
}
