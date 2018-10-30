// Copyright (C) 2018

// Implements Weighted Product Model (WPM) defined in
// https://en.wikipedia.org/wiki/Weighted_product_model

package main

import (
	"context"
	"fmt"
	"math"

	"github.com/spacemonkeygo/errors"
)

var (
	ctx = context.Background()

	WPMError                  = errors.NewClass("wpm")
	WPMResourcesNotEqualError = WPMError.NewClass("resouce length is not equal")
	WPMDivisionByZeroError    = WPMError.NewClass("division by zero")
	WPMWeightsError           = WPMError.NewClass("weight sum is not 1")
)

func GetWPM(ctx context.Context, a_resources []float64,
	b_resources []float64, weights []float64) (result float64, err error) {

	// check for length of given slices
	if len(a_resources) != len(b_resources) ||
		len(a_resources) != len(weights) {
		return 0, WPMResourcesNotEqualError.New("")
	}

	// division by zero check
	for _, value := range b_resources {
		if value == 0 {
			return 0, WPMDivisionByZeroError.New("")
		}
	}

	// check if sum of weights is equal to 1.0
	var weight_sum float64 = 0
	for _, value := range weights {
		weight_sum = weight_sum + value
	}

	// check if weight sum is in 0.00001 range of 1.0
	if math.Abs(weight_sum-1.0) > 0.00001 {
		return 0, WPMWeightsError.New("")
	}

	// init wpm result
	var wpm_result float64 = 1.0

	// https://en.wikipedia.org/wiki/Weighted_product_model
	// loop over, calcualte ratio and update wpm result
	for i, _ := range a_resources {
		// calcualte ratio
		ratio := float64(a_resources[i] / b_resources[i])

		// update wpm result
		wpm_result = wpm_result * math.Pow(ratio, weights[i])
	}

	return wpm_result, nil
}

func main() {
	// setup some values for main function
	// tests include more values
	res1 := []float64{25, 20, 15, 30}
	res2 := []float64{10, 30, 20, 30}
	weights := []float64{0.20, 0.15, 0.40, 0.25}
	result, err := GetWPM(ctx, res1, res2, weights)
	if err == nil {
		fmt.Println("WPM is: ", result)
	} else {
		fmt.Println(err)
	}
}
