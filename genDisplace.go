package main

import "fmt"

func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&acceleration)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&initialVelocity)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&initialDisplacement)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Print("Enter time: ")
	fmt.Scan(&time)
	fmt.Println("Displacement after time t:", fn(time))
}
