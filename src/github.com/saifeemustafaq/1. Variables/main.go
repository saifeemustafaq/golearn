package main

import (
	"fmt"
	"math"
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
	e := 10.654
	g := 3.543

	fmt.Printf("e = 11, g = 3\n")
	fmt.Printf("%v e + g\n", e+g)
	fmt.Printf("%v e - g\n", e-g)
	fmt.Printf("%v e * g\n", e*g)
	fmt.Printf("%v e / g\n", e/g)
	fmt.Printf("%v e %% g\n", math.Mod(e, g)) //escape sequence for % in %%

	//type conversions

	var a1 int8 = 5
	var b1 int16 = 10
	fmt.Println(int(a1) + int(b1))

	//bit operations

	c1 := 10 //1010
	d1 := 3  //0011

	fmt.Println(c1 & d1)  // 0010 = 2
	fmt.Println(c1 | d1)  // 1011 = 11
	fmt.Println(c1 ^ d1)  // 1001 = 9
	fmt.Println(c1 &^ d1) // 1000

	//bit shifting
	e1 := 8              //2^3
	fmt.Println(e1 << 3) //2^3 * 2^3 = 2^6
	fmt.Println(e1 >> 3) //2^3 / 2^3 = 2^0

	f1 := 3.14
	f1 = 13.7e72
	f1 = 2.1E14

	fmt.Printf("%v %T", f1, f1)

	//complex number
	g1 := 1 + 2i
	h1 := 2 + 5.2i

	fmt.Printf("g1 = 1 + 2i AND h1 = 2 + 5.2i\n")
	fmt.Printf("%v g1 + h1\n", g1+h1)
	fmt.Printf("%v g1 - h1\n", g1-h1)
	fmt.Printf("%v g1 * h1\n", g1*h1)
	fmt.Printf("%v g1 / h1\n", g1/h1)
}
