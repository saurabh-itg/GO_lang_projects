package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main(){
	//Saurabh := person{firstName:"Saurabh", lastName:"Joshi"}
	//fmt.Println(Saurabh)
	var sonu person
	sonu.firstName ="Saurabh"
	sonu.lastName="Joshi"
	fmt.Println(sonu)
	fmt.Println("%+v",sonu)
}
