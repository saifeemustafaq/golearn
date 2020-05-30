package main

import (
	"fmt"
	"strconv"
)

var n string = "mustafa7@7" // globally declared variable

var I int = 12 // this will not raise an error

func main() {
	var i int //initialized variable
	i = 1     // declared variable
	fmt.Println(i)

	var j int = 12 //intialized and declared var
	fmt.Println(j)

	k := 123 // auto-detect variable type
	fmt.Println(k)

	var l float32 = 1234
	fmt.Printf("%v %T\n", l, l)

	var m float64 = 1234
	fmt.Printf("%v %T\n", m, m)

	fmt.Printf("%v %T\n", n, n)

	var o int = 23 //initialized variable
	fmt.Printf("%v, %T\n", o, o)

	var p string
	p = strconv.Itoa(o)
	fmt.Printf("%v, %T\n", p, p)

	var a bool = false          //boolean
	fmt.Printf("%v %T\n", a, a) //boolean

	b := 1 == 1                 //boolean
	c := 1 == 2                 //boolean
	fmt.Printf("%v %T\n", b, b) //boolean
	fmt.Printf("%v %T\n", c, c) //boolean

	//int8 //int16 //int32 /int64

	var d uint16 = 42
	fmt.Printf("%v %T\n", d, d)

	//lets do some basic arithmetic operation
	e := 10
	g := 3

	fmt.Printf("e = 11, g = 3\n")
	fmt.Printf("%v e + g\n", e+g)
	fmt.Printf("%v e - g\n", e-g)
	fmt.Printf("%v e * g\n", e*g)
	fmt.Printf("%v e / g\n", e/g)
	fmt.Printf("%v e %% g\n", e%g) //escape sequence for % in %%

}
