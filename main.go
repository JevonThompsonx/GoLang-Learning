package main

import (
	"GoLang-Learning/hello"
	"GoLang-Learning/importMePackage"
	"GoLang-Learning/loops"
	"GoLang-Learning/variables"
	"fmt"
)

func main() {
	// all variables/functions capitalized within other files
	// but packages within the same package (in this case "main")
	// can be used across each other
	// eg:
	hello.Hello()
	variables.Variables()
	// variables & functions from other packages,
	// need to be explicitly imported (see top)
	// and used via their package:
	importMePackage.Wave()
	fmt.Println(variables.ExportMe)
	variables.Constants()
	loops.Loops()
}
