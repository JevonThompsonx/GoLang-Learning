package variables

import "fmt"

func Constants() {
	fmt.Println("This is where we make some constants")
	const constantNum int = 100
	fmt.Println("Here is a constant number:", constantNum)
	// constantNum = 22
	// fmt.Println("I have attempted to change the value which throws an error")
}
