package main

import (
	"testing"
	"fmt"
)

func TestComputerV1(t *testing.T) {

	sep := "---"

	fmt.Println("Solving: 5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0")
	ComputorV1("5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0")
	fmt.Println(sep)

	fmt.Println("Solving: 5 * X^0 + 4 * X^1 = 4 * X^0")
	ComputorV1("5 * X^0 + 4 * X^1 = 4 * X^0")
	fmt.Println(sep)

	fmt.Println("Solving: 5 * X^0 = 5 * X^0")
	ComputorV1("5 * X^0 = 5 * X^0")
	fmt.Println(sep)

	fmt.Println("Solving: 4 * X^0 = 8 * X^0")
	ComputorV1("4 * X^0 = 8 * X^0")
	fmt.Println(sep)

	fmt.Println("Solving: 5 * X^0 = 4 * X^0 + 7 * X^1")
	ComputorV1("5 * X^0 = 4 * X^0 + 7 * X^1")
	fmt.Println(sep)

	fmt.Println("Solving: 5 * X^0 + 13 * X^1 + 3 * X^2 = 1 * X^0 + 1 * X^1")
	ComputorV1("5 * X^0 + 13 * X^1 + 3 * X^2 = 1 * X^0 + 1 * X^1")
	fmt.Println(sep)

	fmt.Println("Solving: 6 * X^0 + 11 * X^1 + 5 * X^2 = 1 * X^0 + 1 * X^1")
	ComputorV1("6 * X^0 + 11 * X^1 + 5 * X^2 = 1 * X^0 + 1 * X^1")
	fmt.Println(sep)

	fmt.Println("Solving: 5 * X^0 + 3 * X^1 + 3 * X^2 = 1 * X^0 + 0 * X^1")
	ComputorV1("5 * X^0 + 3 * X^1 + 3 * X^2 = 1 * X^0 + 0 * X^1")
	fmt.Println(sep)

	fmt.Println("Solving: 8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0")
	ComputorV1("8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0")
	fmt.Println(sep)

	fmt.Println("Solving: 8*X^0-6*X^1+0 * X^2 - 5.6 * X^3 = 3 * X^0")
	ComputorV1("8*X^0-6*X^1+0 * X^2 - 5.6 * X^3 = 3 * X^0")
	fmt.Println(sep)

	fmt.Println("Solving: 5 * X + 4 * X^1 = 4 * X^0")
	ComputorV1("5 * X + 4 * X^1 = 4 * X^0")
	fmt.Println(sep)

	fmt.Println("Solving: empty string")
	ComputorV1("")
	fmt.Println(sep)
}
