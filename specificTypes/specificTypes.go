package specifictypes

import "fmt"

func Specifictypes() {
	fmt.Println("Hola from specific types")
	fmt.Println("The goal here is to make use of types that are as specific as possible. Known types can save space by being specified")
	fmt.Println(`
    egs:
    int vs int8,int16,int32
    float vs float32,float64
    uint vs uint8,uint16,uint32
    `)
	fmt.Println(`
    ints: basic number type. covers positive and negatives
    uint: positive number only type
    float: decimal point numbers
    `)

	fmt.Println(`
    uints can save space if a number is only going to be positive

    even more space can be saved by making that number a uint8,16,32 when applicable
    `)
	fmt.Println(`
    This is how I'll try to remember it
    uint8: 0 - 255
    uint16: up to 65000 (about)
    uint32: up to 4294967295  
    uint64: up to 18446744073  
    `)
	fmt.Println(`
    I def won't remember all that.
    I'll just try to remeber 8 and 16. If I need more than that, I'll know`)
	fmt.Println(`
    So since uint is all positives, it makes it easy to remeber that ints are about half as much room in both directions
    egs

    uint8: 0-255
    int8: -128-128

    That should make it easy enough

    Let's do the rest

    int16: 32500 (about)
    int32: 2594967295 (about)

    I'm not really doing math here. just guessing because the general number is most important
    `)
}
