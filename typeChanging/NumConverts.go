package typechanging

import (
	"GoLang-Learning/review"
	"fmt"
	"reflect"
)

func print(text string) {
	fmt.Println(text)
}

func NumConverts() {
	print("Hello from the typechanging module")
	print("Let's see what we have to review")
	print(review.DAY3NEWSTUFF)

	print("We can change types by using the syntax `var variableName type = type(variable)`")
	print("If we wanna be super verbose about it")
	print("Shorthand `variableName := type(variable)`")
	print("Let's try")
	anIntNumber := 23
	print("-------------------")
	fmt.Println("Type of anIntNumber:", reflect.TypeOf(anIntNumber))
	aFloatNumber := float32(anIntNumber)
	print("")
	print("aFloatNumber := float32(anIntNumber)")
	fmt.Println(aFloatNumber)
	print("")
	fmt.Println("Type of aFloatNumber", reflect.TypeOf(aFloatNumber))
	print("-------------------")
}
