package main

import (
	"fmt"
	"os"
)

type term struct {
	coef	float64
	degree	int
}

func parse(equation string) error {

	fmt.Printf("parse: %v\n", equation);

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
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
	// Reduce list of terms
	// Find discriminant
	// Find solutions
}
