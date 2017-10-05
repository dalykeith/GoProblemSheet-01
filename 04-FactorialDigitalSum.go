//04-FactorialDigitalSum.go
//Author: Keith Daly

package main
import "fmt"
func main{
	
	
	var num int = 10
	var factor int = 1


	//fmt.Println("Welcome to Factorial Digital Sum")

	for i:=0; i < num; i++ {

		factorial += factorial * i
	}

	//output factorial 
	fmt.Println(factorial)
	var digit, sum int = 0,0;

	for factorial > 0 {
		digit = factorial  % 10
		sum += digit
		factorial /=10
	}

	fmt.Println(sum)
}


//Adapted from:
//https://play.golang.org/p/yW7sAyJpPe
//https://gist.github.com/shockalotti/7fa310e915ee66766039
//https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion