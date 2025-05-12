package main

import "fmt"

func Variables() {
	fmt.Println("1 + 4 = ", +1+4)
	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(!true)
	fmt.Println("7/3=", 7/3)
	// comment like this
	fmt.Println("Variables are declared using var")
	var firstNum, secondNum int = 22, 24
	fmt.Println("first number", firstNum, "and second num", secondNum)
	greet := "Hello there buddy"
	fmt.Println(greet)
	fmt.Println("Generally, it's best practice to use type inference for types that can be infered easily")
	fmt.Println("In my code editor right now, if I create a variable, my linter/formatter fixes it automatically which is nice")
	fmt.Println("this makes declaring variables as easy as name := value")
	thirdNum := 89
	fmt.Println("Third num", thirdNum)
}
