package main

import "fmt"

func main(){
	slice1 := []int{1,2,3}
	slice2 := []string{"How","are","you"}



	if len(slice1) == len(slice2){

		newMap := make(map[int]string) //new map var declaration

		for i:=0; i<len(slice1);i++{

			newMap[slice1[i]] = slice2[i]
		}
		fmt.Println(newMap)

	}

	
}