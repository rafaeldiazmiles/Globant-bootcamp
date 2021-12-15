package main

import (
	"fmt"
	"os"
	"strconv"

	calculor "github.com/Ksualboy/Globant-bootcamp/2ndPart/calculator/calculor"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Syntax error: expected only 3 arguments")
		fmt.Println("Example: 3 + 5")
		return
	}

	var result float64
	a, err := strconv.ParseFloat(os.Args[1], 64)
	b, err2 := strconv.ParseFloat(os.Args[3], 64)

	// Handling parsing errors
	if err != nil || err2 != nil {
		if err != nil {
			fmt.Println("error 1:", err)
		}
		if err2 != nil {
			fmt.Println("error 2:", err2)
		}
		return
	}

	switch os.Args[2] {
	case "+":
		result = calculor.Add(a, b)
	case "-":
		result = calculor.Subs(a, b)
	case "*":
		result = calculor.Mul(a, b)
	case "/":
		result, err = calculor.Div(a, b)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("Syntax error: Invalid operator")
		fmt.Println("Available operators: +, -, * and /")
		return
	}

	fmt.Printf("Result: %.3f\n", result)
}
