package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	str "strings"
	"math"
)

type term struct {
	coef   float64
	degree int64
}

func parse(equation string) ([]term, error) {

	words := str.Split(equation, " ")

	terms := []term{}
	var sign, inv float64 = 1, 1
	var curr term

	for _, val := range words {

		if n, err := strconv.ParseFloat(val, 64); err == nil {
			curr.coef = n * sign * inv
			sign = 1
		} else if val == "*" {
			// Do nothing
		} else if val == "+" {
			// Do nothing
		} else if val == "-" {
			sign = -1
		} else if val == "=" {
			inv = -1
		} else {
			tmp := str.Split(val, "^")
			if len(tmp) != 2 || tmp[0] != "X" {
				return nil, fmt.Errorf("too many '^' operators in term expression")
			}
			e, err := strconv.ParseInt(tmp[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("incorrect exponent in term expression")
			}
			curr.degree = e
			terms = append(terms, curr)
			curr = term{0, 0}
		}
	}

	return terms, nil
}

func printPolynome(polynome []term) {
	printed := false
	for _, term := range polynome {
		if printed {
			fmt.Printf(" ")
			if term.coef < 0 {
				fmt.Printf("- ")
			} else {
				fmt.Printf("+ ")
			}
		}
		fmt.Printf("%v * X^%v", math.Abs(term.coef), term.degree)
		printed = true
	}
	fmt.Println(" = 0")
}

func reduce(terms []term) []term {

	sort.Slice(terms, func(i, j int) bool {
		return terms[i].degree > terms[j].degree
	})

	degree := terms[0].degree

	polynome := []term{}
	for i := 0; int64(i) < degree+1; i++ {
		t := term{0, int64(i)}
		polynome = append(polynome, t)
	}

	for _, term := range terms {
		polynome[term.degree].coef += term.coef
	}

	fmt.Print("Reduced form: ")
	printPolynome(polynome)
	fmt.Printf("Polynomial degree: %v\n", degree)
	return polynome
}

func solveDegreeOne(polynome []term) {

	fmt.Println("The solution is:")
	root := -1 * polynome[0].coef / polynome[1].coef
	fmt.Println(root)
}

func solveDegreeTwo(polynome []term) {

	// Find discriminant
	a, b, c := polynome[2].coef, polynome[1].coef, polynome[0].coef
	discriminant := b * b - 4 * a * c

	// Solve
	if discriminant == 0 {
		fmt.Println("Discriminant is equal to 0, the solution is:")
		s := ((-1 * b) / (2 * a))
		fmt.Println(s)
	} else {
		if discriminant > 0 {
			fmt.Println("Discriminant is strictly positive, the two solutions are:")
		} else {
			fmt.Println("Discriminant is strictly negative, the two solutions are:")
		}
		s1 := ((-1 * b - math.Sqrt(discriminant)) / (2 * a))
		s2 := ((-1 * b + math.Sqrt(discriminant)) / (2 * a))
		fmt.Println(s1)
		fmt.Println(s2)
	}
}

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: ./computor equation")
		return
	}
	// Parse equation into a list of terms
	terms, err := parse(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	// Reduce list of terms
	polynome := reduce(terms)
	// Solve equation
	lenPoly := len(polynome)
	if lenPoly == 3 {
		solveDegreeTwo(polynome)
	} else if lenPoly == 2 {
		solveDegreeOne(polynome)
	} else if lenPoly == 1 {
		fmt.Println("The solution is:")
		fmt.Println(polynome[0].coef)
	} else if lenPoly > 2 {
		fmt.Println("The polynomial degree is strictly greater than 2, I can't solve.")
	} else {
		fmt.Println("All real numbers are solutions.")
	}
}
