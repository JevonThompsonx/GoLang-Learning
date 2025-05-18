package typechanging

import (
	"fmt"
	"reflect"
	"strconv"
)

func StringConverts() {
	fmt.Println("Let's change some string baybeeeee")
	fmt.Println("Shoutout to how easy imports and exports are in go")
	fmt.Println(`
    String conversions with go are best done with the strconv package
    `)
	fmt.Println(`
    conversions done from int to string: strconv.Itoa
    conversions ||    || string to int: strconv.Atoi
    `)
	fmt.Println(`Test time`)
	const randomNumba int8 = 120
	fmt.Println(`
  const randomNumba int8 = 120
    `)
	stringNumba := strconv.Itoa(int(randomNumba))
	fmt.Println(`
	var stringNumba string = strconv.Itoa(int(randomNumba))
    `)
	fmt.Println("randomNumba = ", randomNumba, "and stringNumba = ", stringNumba)
	fmt.Println("Now let's check the types")
	fmt.Println("randomNumba type = ", reflect.TypeOf(randomNumba))
	fmt.Println("stringNumba type = ", reflect.TypeOf(stringNumba))
}
