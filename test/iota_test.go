package test

import (
	"fmt"
	"testing"
)

// https://yourbasic.org/golang/iota/
// Test iota  [aɪˈoʊtə]
const (
	A0 = iota
	A1 = iota
	A2 = iota
)
// This can be simplified to
//const (
//	A0 = iota
//	A1
//	A2
//)

// Start from one
const (
	B1 = iota + 1
	B2
	B3
)

// Skip value
const (
	C1 = iota + 1
	_
	C3
	C4
)

// Complete enum type with strings [best practice]
type Direction int
const (
	North Direction = iota
	East
	South
	West
)

func TestIota(t *testing.T)  {
	// iota basic example
	fmt.Println(A1) // "1"

	// Start from one
	fmt.Println(B1, B2, B3) // "1 2 3"

	// Skip value
	fmt.Println(C1, C3, C4) // "1 3 4"

	// Complete enum type with strings [best practice]
	var d Direction = North
	fmt.Println("North Value:",d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}
