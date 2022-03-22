package main

import "fmt"

const Pi = 3.1416

func circle(radius float64) (area func() float64, perimeter func() float64) {
	area = func() float64 {
		return Pi * radius * radius
	}
	perimeter = func() float64 {
		return 2 * Pi * radius
	}
	return
}

func main() {
	var radius float64
	fmt.Println("would you like to let us know the circle's radius : \n ")
	fmt.Scanf("%f", &radius)

	area, perimeter := circle(radius)
	fmt.Println("area value is : ", area())
	fmt.Println("perimeter value is : ", perimeter())

}
