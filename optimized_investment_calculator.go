package main

import (
	"fmt"
	"math"
)

func main2() {
	
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)


	fmt.Println(futureValue)
	

}
