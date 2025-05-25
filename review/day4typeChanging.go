package review

import (
	"fmt"
	"reflect"
)

func Day4TypeChanging() {
	println(`
   Now to review changing types

    This one isn't fresh in my memory since I haven't reviewed it as much but...

    to convert between numbers we can do 

    var helloNum int8 = 78
	var helloFloat float32 = float32(helloNum)
    `)
	var helloNum int8 = 78
	var helloFloat float32 = float32(helloNum)

	Println("The reflect package can be used to check types")
	Println(`
  fmt.Println(reflect.TypeOf(helloNum))
	fmt.Println(reflect.TypeOf(helloFloat))
    `)
	fmt.Println(reflect.TypeOf(helloNum))
	fmt.Println(reflect.TypeOf(helloFloat))
}
