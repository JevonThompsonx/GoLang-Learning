package review

import "fmt"

func Day1() {
	fmt.Println("Hola from day 1")

	fmt.Println("Here I will review some basics that I have been learning with my about half an hour a day of go")
	fmt.Println("to make this easier, I'll make a Println function in the this package")
	Println("Let's start with variables")
	Println("Variables are defined using var, const or := ")
	var testVariable string = "Hello I am a test var"
	Println("Variables can assume type if using walrus operator :=")
	testVariable2 := 800
	Println("Variables initiated with const, don't by default need to have a type")
	const testVariable3 = 90
	Println(testVariable)
	fmt.Println(testVariable2)
	println(testVariable3)
	fmt.Println("Some math using the const var: testVar/2:", testVariable3/2)
	fmt.Println("Some math using the const var: testVar%2:", testVariable3%2)
	Println("During these operations, the value was converted to an int")
	Println("This next example will make it a float")
	fmt.Println("Some math using the const var: testVar - 2.4:", testVariable3-2.5)
	Println("const is valuable not only for making a constant value but also a value that is flexible without direct modification that will affect other functions using that value")
	Println("Later I should review functinos and for loops but that will have to wait")
}
