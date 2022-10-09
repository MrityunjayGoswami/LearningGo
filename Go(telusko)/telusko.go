package main

import (
	"fmt"
	"math"
	"time"
)
func main(){

	var num1 = 2
	var num2 = 3
	result1, result2 := calc(num1, num2)

	// we can replace var with :=
	if num1<5 {
		print("addition is : ",result1)
	}else if num1==5{
		fmt.Println("substration is : ",result2)
	}else{
		fmt.Println("none")
	}
	print("\n")
	i := 1
	//also we can declare loop at once like this: for i:=1;i<=2;i++
	for i<=2{

		fmt.Println("hello",i)
		i++
	}
	var num float64 = 9
	var result = math.Sqrt(num)
	fmt.Printf("square root is %g, Thank you\n",result)

	swith()
	fmt.Println(time.Now().Weekday())

}
//when we create a new fuction we have to declare the type of parameters and return value

func calc(x, y int) (int,int){
	add := x+y
	sub := x-y
	return add,sub
}

func swith(){
	num := 2
//In go you don't need to use breaks for switch cases
	switch num{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("None")
	}
}