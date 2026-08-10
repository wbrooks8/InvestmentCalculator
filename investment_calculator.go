package main

import (
	"fmt"
	"math"
)

const InflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate = 5.5
	var years float64

	// "&" creates a pointer to the variable so it can be reassigned.
	fmt.Println("Please give your investment amount.")
	// fmt.Scan can struggle with multi word inputs
	fmt.Scan(&investmentAmount)

	fmt.Println("Please give the number of years")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// Calculating with inflation
	futureRealValue := futureValue / math.Pow(1+InflationRate/100, years)

	// fmt.Sprintf allows the creation of formatted string that can be stored as var
	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.2f\n", futureRealValue)

	// Printf allows the use of string formatters like %4
	// fmt.printf("Future Value: %.2f\n", futureValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)

}

func outputText(text string, text2 string) {
	fmt.Print(text, text2)
}

// function definition for single return
// func sampleFunc(sampleValue int, sampleString string) string {
	// return string
// }

// function definition for multiple return values
// func sampleFunc(sampleValue int, sampleString string) (string, string) {
	// return string, string
// }


// function definition for multiple return values and instantiating them in decleration
// func sampleFunc(sampleValue int, sampleString string) (st1 string, st2 string) {
// 		return can look like either of the following:
//		return
//		return st1, st2
// }