package review

import "fmt"

func Day5ArraysSlices() {
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
    - eat
    `)
}
