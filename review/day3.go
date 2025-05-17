package review

import "fmt"

func Day3() {
	Println("Day 3 time")
	Println("Today is a bigger day. I have lots to review and some new stuff to try")

	Println("Let's review declaration")
	Println("var is used to declare")
	var variableName string = "Hola. I am variable"
	Println(variableName)
	Println("Variables can also be shortned in their declaration without a type using := which assumes the type")
	variableNam2 := "I am also a string"
	Println(variableNam2)
	variableNam3 := 89
	fmt.Println(variableNam3)
	Println("Const is used to declare variables that don't change")
	const conVar int = 800
	fmt.Println(conVar)

	fmt.Println(`
    Let's leave a note for what to review later: 
    - changing variable types
    - using specific type declarations: int8, float32, uint16 etc
    - using bcrypt to hash passwords 
    - for Loops
    - 
    `)
}
