//04-FactorialDigitalSum.go
//Author: Keith Daly

package main

import "fmt"

	func main
	{
	
	var num int = 8
	var factorNum int = 1
	//fmt.Println("Welcome to Factorial Digital Sum")

	for i:=0; i < num; i++ 
	{
		factorNum += factorNum * i
	}
	//output factorial 
	fmt.Println(factorNum)
	var digit, sum int = 0,0;

	for factorial > 0 
	{
		digit = factorNum  % 10
		sum += digit
		factorNum /=10
	}
	fmt.Println(sum)
}


//Adapted from:
//https://play.golang.org/p/yW7sAyJpPe
//https://gist.github.com/shockalotti/7fa310e915ee66766039
//https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion
