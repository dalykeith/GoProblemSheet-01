//06-LargestAndSmallestInList.go
//Author: Keith Daly

package main

import "fmt"

func main(){

    //10 ints in array
    list:=[10] int{41, 7, 2, 6, 14, 9, 3, 26, 10, 64}   
    
    //declare variables
    var largest int =0 
	var smallest int =0
	
	//init variables 
	largest = list[0]
	smallest = list[0]

    //for loop
    for i := range list{

        //nested for loop 
		if array[i] < smallest
        {
			smallest = list[i]
		} 
		if array[i] > largest
        {
			largest = list[i]
		}

	}
    //output smallest int and the largest int
	fmt.Println("The smallest number is: ",smallest,"\nThe largest number is: ",largest)
	


}