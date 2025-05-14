package main

import (
	"GoLang-Learning/importMePackage"
	"fmt"
)

func main() {
	// all variables/functions capitalized within other files
	// but packages within the same package (in this case "main")
	// can be used across each other
	// eg:
	Hello()
	Variables()
	// variables & functions from other packages,
	// need to be explicitly imported (see top)
	// and used via their package:
	importMePackage.Wave()
	fmt.Println(ExportMe)
}
