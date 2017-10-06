//10-ReverseString.go
//Author: Keith Daly

//Adapted from https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
//https://github.com/golang/example/blob/master/stringutil/reverse.go

package main

import "fmt"


func revStr(s string) string {
	var reversedStr  string
	
    // Loop through the string backwards and store in a string
    for i := len(s)-1; i >= 0; i-- {
        reversedStr  += string(s[i])
    }
    return reversedStr 
}
 
func main() {
	var str string
	
    //start
    fmt.Println("Sentence Reverser!")
    //user input
	fmt.Println("\nPlease enter a sentence to be reversed: ")
	//scan string
	fmt.Scanf("%s ", &str)

	
	//output reversed sentence:
	fmt.Println("Reversed Sentence: ", revStr(str))
}