//07-PalindromeTest.go
//Author: Keith Daly

//Adapted from: http://www.golangpro.com/2016/01/check-string-palindrome-go.html

package main

import "fmt"
import "strings"

function main(){
  
    var input string

    fmt.Println("Palindrome Checker")
    fmt.Println("\nA palindrome is a word, phrase, number, or other sequence 
    of characters which reads the same backward or forward.")
    
	fmt.Println("Please enter a word: ")
    //string input
	fmt.Scanf("%s\n", &input)
	input = strings.ToLower(input)
	fmt.Println(isP(input))
}

//Function reads in word, from the last, then counts back to find midpoint
 func isP(s string) string 
 {
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
            //if word matches count both sides output:
			return "This word is not a Palindrome!"
		}
	}//else
    //if word does not matches count both sides output:
	return "This word is a Palindrome!"
}