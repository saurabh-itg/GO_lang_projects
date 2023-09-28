package main

import "fmt"

// Define a custom type
type Rectangle struct {
    Width  float64
    Height float64
}

// Define a method on the Rectangle type
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // Create an instance of the Rectangle type
    rect := Rectangle{Width: 5.0, Height: 3.0}

    // Call the Area method on the Rectangle instance
    area := rect.Area()

    fmt.Printf("Area of the rectangle: %f\n", area)
}
