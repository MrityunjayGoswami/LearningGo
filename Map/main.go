package main

import (
	"fmt"
)

func main(){
	colors := map[string] string{
		"Red": "#ff0000",
		"Blue": "#0000ff",
		"White": "#ffffff",
		
	//also map can be delared by 1: var colors map[string]string
	}
	delete(colors,"White")
	//map can be declared by this way also:
	number := make(map[int]string)
	number[1] = "One"
	number[2] = "two"

	printMap(colors)
	//fmt.Println(colors)
	fmt.Println(number)
	//fmt.Println(reflect.TypeOf(color))
}

func printMap(c map[string]string){
	for color, hex := range c{
		fmt.Println("Hex code for ",color,"is ",hex)
	}	
}

//Map is different in many ways with struct:

//1 : Map keys should be same data type and keys work as index
//so we can iterate through it.

//2 : struct valuse can be of diff data types.

//3 : Map is reffrence type whereas struct is value type.

//4 : (map) used to represent a collection of related properties.

//5 : (struct) used to represent a thing with a lot of diff properties.
