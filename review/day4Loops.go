package review

import "fmt"

func Day4Loops() {
	fmt.Println("time to review some foooorrrrr loops baybee")
	Println(`
    I will review everything else first then move on to for loops

    Let's start off with variables.
    
    ## Variables 

    to init a variable const, var or := can be used
    
    ### egs: 
    
    `)
	var variable1 string = "I am a variable"
	const variable2 int8 = 78
	fmt.Println(variable1)
	fmt.Println(variable2)

	println(`
    eg of the walrus operator :=

    variable3 := 78

    The walrus operator is used to assume the type based on the info provided IN this case, the info provided says "this is an int"
    `)

	fmt.Println(`
    Now let's move on to variable types

    ## Variable types

    ### Basic types

    The basic types that variables can be are: 
    int, string, boolean 

    I'll learn some more as we go 

    #### Basic types expanded

    The basic types can be broken down further to save memory. When variables are initalized by their default basic types, some extra memory is given if the variable declaration wasn't specific. Specifics are int8, int16 etc vs int 

    egs: 

    var variable4 int8 = 24

    For when to use these types, it matters how large the variables are. For ints

    int8 = -128 to 128 
    int16 = about -32000 to 32000 
    etc 

    I will know when there's time to get more specific than that. I likely will never need to use more than those 2 if we're being honest. 

    uints are ints that are only positive. 

    uint8 = 0 to 256 
    uint16 = 0 to 65000

    these aren't exact since I can't remember the int16s and uint16s but again, that doesn't matter much right now. 

    floats can also be more specific: 
    float32 and float64
    
    I'll likely only need float32

    More info: 

    | Category     | Types                                                    |
| ------------ | -------------------------------------------------------- |
| Boolean      | bool                                                  |
| Integers     | int, int8, int16, int32, int64                 |
| Unsigned Int | uint, uint8, uint16, uint32, uint64, uintptr |
| Floats       | float32, float66                                   |
| Complex      | complex64, complex128                                |
| String       | string                                                 |
| Aliases      | byte, rune                                           |

  
    ### Booleans 

    Booleans are values that are only true of false. On a binary level it's just a zero or one

    boolVar := false


    Now we can move on to some for loops! 
    `)

	fmt.Println("Time for the main event in this review. For loooooooooops")
}
