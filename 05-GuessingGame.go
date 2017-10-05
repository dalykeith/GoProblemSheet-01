//05-GuessingGame.go
//Author: Keith Daly

package main 
import "fmt" 
import "math/rand" 
import "time"


func main(){

	var input int
	randomNum := randomNumGen(1,100)
	var prevousNum int
	var found bool = false
	var count int = 0

	fmt.Println("Welcome to Guessing Game")
	
	for found != true{

		
		fmt.Println("Please Enter a Number between 1-100: ")
		fmt.Scan(&input)
		count++

		if input == prevousNum{
			count--
			fmt.Println("Guess Again")
            fmt.Println("\nThat number has already been entered")
		}

		//if else if code block to sort number 
		if input > randomNum{
			fmt.Println("Your guess is too high!")
		}else if input < randomNum{
			fmt.Println("Your guess is too low"!)
		}else{
			fmt.Println("You guessed the correct number")
			fmt.Println("\nin",count,"tries")
			found = true // ends loop
		}

		prevousNum = input



	}

}
