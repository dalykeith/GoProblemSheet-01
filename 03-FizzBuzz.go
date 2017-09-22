//03-FizzBuzz.go
//Author: Keith Daly

//Wiki for FizzBuzz http://wiki.c2.com/?FizzBuzzTest

package main
import "fmt"

func main() {

    //print the numbers from 1 to 100
	for i := 1; i <= 100; i++{
		
		var result string = ""
        //if the number is multiples of three print Fizz
		if i % 3 == 0 {
			result += "Fizz"
		}
        //if the number is multiples of 5 print Buzz
		if i % 5 == 0 {
			result += "Buzz"
		}
		//if the number is multiples of 3 & 5 print Fizz Buzz
		if i % 3 == 0 && i % 5 == 0 {
		fmt.Println(i, Fizz + Buzz)

		if result != "" {
			fmt.println(result)
		}
        else{
			fmt.println(i)
			continue
		}	
	}
}