//09-NewtonsMethodForSquareRoots.go
//Author: Keith Daly

//https://tour.golang.org/flowcontrol/8
//https://en.wikipedia.org/wiki/Methods_of_computing_square_roots
//https://gist.github.com/abesto/3476594
//https://golang.org/pkg/math/#Sqrt

//Newtons algorithm for guessing the square root of a number.
//Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating: z - (z*z - x) / (2 * z)

package main 

import "fmt"
import "math"

func main() {

    //X number (that will be found to be square root of)
	x := 67.0 
    //Z = first guess
	z := 1.0  

    fmt.Printf("Newtons Square Root of a number calculator")

	//Run code while z doesn't equal the guess
	for z = 1.0; z != squareRoot(z, x); z = squareRoot(z, x) 
    {
        //output guess:
		fmt.Printf("\nGuess: %2f\n", z)
	}

	//output approximation according to Netwon's Method for sqaure root:
	fmt.Printf("\n%2f is an approximation for the square root of %2f\n", z, x)

	//output square root:
    fmt.Printf("\nThe math.Sqrt gives the value of %2f", math.Sqrt(x))
}

func squareRoot(z float64, x float64)float64
    {
	if z == 0{
		return 0
	}

    //newtons formula for square root
	return z - (((z*z) - x) / (2*z))	
}

