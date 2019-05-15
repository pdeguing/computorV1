package main

import (
	"fmt"
	"os"
	str "strings"
	"strconv"
)

type term struct {
	coef	float64
	degree	int
}

func parse(equation string) error {

	fmt.Printf("parse: %v\n", equation);
	words := str.Split(equation, " ")
	fmt.Println(words)

	//terms := []term{}
	//sign := 1
	var curr term
	fmt.Println(curr)
	for  _, val := range words {
		fmt.Println(val)

		if _, err := strconv.ParseFloat(val, 64); err == nil {
			fmt.Println("val is number")
		} else if val == "*" {
			fmt.Println("val is *")
		} else if val == "+" {
			fmt.Println("val is +")
		} else if val == "-" {
			fmt.Println("val is -")
		} else if val == "=" {
			fmt.Println("val is =")
		// add condition to validate X^y
		} else {
			return fmt.Errorf("unrecognized token: %v", val)
		}
	}

	return nil
}

func main() {

	args := os.Args[1:]

	if  len(args) != 1 {
		fmt.Println("Usage: ./computor equation")
		return
	}
	// Parse equation into a list of terms
	if err := parse(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	// Reduce list of terms
	// Find discriminant
	// Find solutions
}
