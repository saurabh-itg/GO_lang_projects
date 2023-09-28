package main

import (
	"fmt"
	
)

func main(){
	var num1 float64
	var sign string
	var num2 float64
	
	fmt.Print("Number1 :")
	fmt.Scan(&num1)
	fmt.Print("sign :")
	fmt.Scan(&sign)
	fmt.Print("Number2 :")
	fmt.Scan(&num2)

	if sign=="+"{
		fmt.Print("Sum is ", num1+num2)
	}else if sign=="/"{
		fmt.Print("divison result is :\n",num1/num2 )
	}else if sign=="*"{
		fmt.Print("multiplication is :",num1*num2 )
	}

}
