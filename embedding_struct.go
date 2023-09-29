package main
import "fmt"
type ContactInfo struct{
	email string
	zipcode int

}
type person struct{
	firstName string
	lastName string
	contact ContactInfo
}
func main(){
	jim := person{
		firstName:"jim",
		lastName:"Carry",
		contact:ContactInfo{
			email:"jim@yahoo.com",
			zipcode:560030,
		},
	}
	fmt.Printf("%+v",jim)
}

