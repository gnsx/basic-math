package main

import (
	"fmt"
	"sort"
)

//FindMean Finds mean/average of []floats
func FindMean(Input []float64) float64 {
	Value := 0.0
	for _, Element := range Input {
		Value += float64(Element)
	}
	return Value / float64(len(Input))
}

//FindMedian - find median of floats
func FindMedian(Input []float64) float64 {
	sort.Float64s(Input)

	if len(Input) == 0 {
		return 0
	}
	if len(Input) == 1 {
		return Input[0]
	}

	if len(Input)%2 == 0 {
		EvenPoint := (len(Input) / 2) - 1
		//fmt.Print("\nUtil-01: EvenNumber for Median EvenPoint:", EvenPoint, "a:", Input[EvenPoint], " b:", Input[1+EvenPoint])

		return (Input[EvenPoint] + Input[1+EvenPoint]) / 2
	}

	//fmt.Print("\nUtil-02: Odd for Median")
	return Input[(len(Input) / 2)]

}

func main() {
	Input := []float64{12.0, 32.10, 10}
	fmt.Print("\nMean:", FindMean(Input))
	fmt.Print("\nMedian:", FindMedian(Input))
}
