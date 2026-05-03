package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(x int, y int) int { return x + y }
func sub(x int, y int) int { return x - y }
func mul(x int, y int) int { return x * y }
func div(x int, y int) int {
	if y == 0 {
		fmt.Println("DIVISION BY ZERO!")
		os.Exit(1)
	}
	return x / y
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Invalid execution!")
		fmt.Println("./calculator <X> <Operation> <Y> \n(x and y are integers and operation is one of this (+-x/))")
		os.Exit(1)
	}

	type operation func(int, int) int // func declare that is used in the ops dict

	ops := map[string]operation{ // dict key is string value is a func
		"+": add,
		"-": sub,
		"x": mul,
		"/": div,
	}
	x, err := strconv.Atoi(os.Args[1]) // x is the first argument converted from str to int
	if err != nil {                    // if strconv.Atoi catch any errors in conversion err will not be nil (None in python)
		fmt.Println(os.Args[1], "is not a integer type")
		os.Exit(1)
	}

	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(os.Args[3], "is not a integer type")
		os.Exit(1)
	}

	var uop string
	if _, ok := ops[os.Args[2]]; ok {
		uop = os.Args[2]

	} else {
		fmt.Println("Invalid operation:", os.Args[2])
		os.Exit(1)
	}

	fmt.Println(ops[uop](x, y))
}
