package main

import (
	"fmt"
)

type contactInfo struct{
	email string
	zipCode string
}

type person struct{
	firstName string
	lastName string
	contact contactInfo
}

func main(){
	//Amit := person{"Amit","Goswami"}
	//var Amit person
	//also it can be created by this to remove order problem:
	// Amit := person{firstName : "Amit",lastName : "Goswami"}

	//And also like this:
	//var Amit person
	//Amit.firstName = "Amit"
	//Amit.lastName = "Goswami"

	//If we dont assign any value in var, it will assign deafault value as:
	// [Type]		[Zero value]
	// String		""
	// int			0
	// float		0
	// boolean		false
	// struct		{}

	//fmt.Println(reflect.TypeOf(Amit))
	//fmt.Println(Amit)
	//fmt.Printf("%+v",Amit)i := 0; i < count; i++ {

	Ram := person{
		firstName: "Shree Ram",
		lastName: "Raghuvanshi",
		contact: contactInfo{
			email: "ayodhyakeraja@gamil.com",
			zipCode: "108",
		},	//when using new line sepration use comma even in end.
	}
	RamPointer := &Ram
	RamPointer.updateName("Maryada purshottam")
	fmt.Println(Ram)
	fmt.Println(Ram.firstName)
	fmt.Printf("%+v",Ram)

	//code to upadate a slice without pointers:
	mySlice := [] string{"Hi","How","are","you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func(pointertoperson *person) updateName(newFirstName string) {
	(*pointertoperson).firstName = newFirstName
}

func updateSlice(s []string){
	s[0] = "Bye"
}