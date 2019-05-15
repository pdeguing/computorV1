package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	str "strings"
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

func reduce(terms []term) ([]term, error) {

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

	fmt.Printf("Reduced form: %v\n", polynome)
	fmt.Printf("Polynomial degree: %v\n", degree)
	if degree > 2 {
		return nil, fmt.Errorf("The polynomial degree is strictly greater than 2, I can't solve.")
	}
	return terms, nil
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
	terms, err = reduce(terms)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	// Find discriminant
	// Find solutions
}
