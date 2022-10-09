package main

import "fmt"

func main(){
	var newslice = []int{1,2,3,4}

	var sum int = 0

	for _, value:= range newslice{
		sum = sum + value
	}


	fmt.Println(sum)
	
}