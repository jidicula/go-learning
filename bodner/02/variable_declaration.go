package main

var x int = 10 // most verbose
var x = 10     // default type of integer literal is int

var x int // declares and assigns zero value

var x, y int = 10, 20 // declare multiple variables
var x, y int          // declare multiple vars and assign zero

// declaration list
var (
	x    int
	y        = 20
	z    int = 30
	d, e     = 40, "hello"
	f, g string
)

// type inference
var x = 10
// is equivalent to
x := 10
