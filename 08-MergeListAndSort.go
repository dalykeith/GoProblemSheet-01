//08-MergeListAndSort.go
//Author: Keith Daly


//Function that merges two sorted lists into a new sorted list
//https://golang.org/pkg/builtin/#append
//https://golang.org/pkg/sort/
//http://austingwalters.com/merge-sort-in-go-golang/


package main

import "fmt" 
import "time"

function main(){

    list1:= [] int {7,4,3,2}
	list2 := [] int {9,5,6}

	//merging list1 & 2 
	list3 := append(list1, list2)

	//sorting arrays
	 sort.Ints(list3)

	//outputing lists 
	fmt.Println("List1: ",list1)
	fmt.Println("\nList 2: ",list2)
	fmt.Println("\nList 3 (list1 and list2 sorted and merged): ",list3)

}
