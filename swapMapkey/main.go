package main

import "fmt"

func main(){

	newmap := map[int]string {1:"one", 2:"two", 3:"three"}

	swappedMap  := make(map[string] int)

	for index, value := range newmap{
		swappedMap[value] = index
	}
	fmt.Println(swappedMap)

}