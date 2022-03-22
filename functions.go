package main

import "fmt"

const Pi = 3.1416

func circle(radius float64) (area float64, perimeter float64) {
	area = Pi * radius * radius
	perimeter = 2 * Pi * radius
	return
}

func main() {
	var radius float64

	fmt.Println("could you share the circle's radius ?")
	fmt.Scanf("%f", &radius)

	area, perimeter := circle(radius)

	fmt.Println(" the circle's area is : ", area)
	fmt.Println(" the circle's perimeter is : ", perimeter)

}
