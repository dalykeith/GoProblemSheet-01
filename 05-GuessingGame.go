//05-GuessingGame.go
//Author: Keith Daly

package main 

import "fmt" 
import "time"
import "math/rand" 



func main(){

	var input int
	ranNum := ranNumGen(1,100)
	var prevousNum int
	var found bool = false
	var count int = 0

    //start the game
	fmt.Println("Welcome to Guessing Game")
	
	for found != true{

        //message to player
		fmt.Println("Please Enter a Number between 1-100: ")
		fmt.Scan(&input)
		count++

		if input == prevousNum
        {
            //if users input is repeated, output:
			count--
			fmt.Println("Guess Again")
            fmt.Println("\nThat number has already been entered")
		}
        //if users input is greater that X random number, output:
		if input > ranNum
        {
			fmt.Println("Your guess is too high!")
		}else if input < ranNum{
    //if users input is less that X random number, output:
			fmt.Println("Your guess is too low"!)
		}else
        {
            //if users input that the same as X random number, output:
			fmt.Println("You guessed the correct number")
			fmt.Println("\nIn",count,"tries")
			found = true 
		}

		prevousNum = input
	}
}
