package main
import "fmt"
func main(){
	slice_of_int := []int{1,2,3,4,5,6,7,8,9,10}

// Check whether each number is even or odd and print accordingly
	for _, num := range slice_of_int{
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
	}
}
}