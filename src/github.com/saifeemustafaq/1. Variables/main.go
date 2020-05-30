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

	var a bool = false
	fmt.Printf("%v %T\n", a, a)

	b := 1 == 1
	c := 1 == 2
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)

}
