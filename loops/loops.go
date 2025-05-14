package loops

import "fmt"

func Loops() {
	fmt.Println("Go uses for as the way to create all loops")
	i := 1
	for i != 11 {
		fmt.Println("This loop will continue until the value of i is 10")
		fmt.Println(i)
		i++
	}
	fmt.Println("Technically the value is 11 now but you get the point", "i =", i)
	for j := 10; j < 20; j = j + 2 {
		fmt.Println("This is a loop with a purpose!")
		fmt.Println("Similar to the last one but")
		fmt.Println(j)
		fmt.Println("The condition for this loop to continue is that j is less than 20")
		fmt.Println("While j is not less than twenty, the value increases by 2")
	}
	fmt.Println("Functions can be set to run a certain number of times using range")
	for count := range 14 {
		count++
		fmt.Println(count)
	}
	// the forever loop
	for {
		fmt.Println("This loop would go on forever")
		break
	}
	fmt.Println("Nesting a condition in the loop:")
	for count := range 20 {
		if count%2 == 0 {
			fmt.Println("Here is the number that will print only if even:", count)
		} else {
			continue
		}
	}
}
