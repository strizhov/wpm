// Copyright (C) 2017 Space Monkey, Inc.

package main

import (
	"fmt"
	"testing"
)

func TestWPMUnequalSlices(t *testing.T) {
	res1 := []float64{1, 2, 3}
	res2 := []float64{4, 5}
	weights := []float64{}

	_, err := GetWPM(ctx, res1, res2, weights)
	if !WPMResourcesNotEqualError.Contains(err) {
		fmt.Println(err)
	}
}

func TestWPMDevisionByZero(t *testing.T) {
	res1 := []float64{1, 2, 3}
	res2 := []float64{4, 5, 0}
	weights := []float64{7, 7, 7}
	_, err := GetWPM(ctx, res1, res2, weights)
	if !WPMDivisionByZeroError.Contains(err) {
		fmt.Println(err)
	}

}

func TestWPMWeightsError(t *testing.T) {

	res1 := []float64{1, 2, 3}
	res2 := []float64{4, 5, 6}
	weights := []float64{7, 7, 7}
	_, err := GetWPM(ctx, res1, res2, weights)
	if !WPMWeightsError.Contains(err) {
		fmt.Println(err)
	}
}

func TestWPMProductionValues(t *testing.T) {
	// test good values from example in
	// https://en.wikipedia.org/wiki/Weighted_product_model

	// 1.
	res1 := []float64{25, 20, 15, 30}
	res2 := []float64{10, 30, 20, 30}
	weights := []float64{0.20, 0.15, 0.40, 0.25}
	result, err := GetWPM(ctx, res1, res2, weights)
	if err != nil {
		fmt.Println(err)
	}
	if result < 1.0 {
		fmt.Println("result is smaller that 1.0")
	}
	truncated_result := fmt.Sprintf("%.3f", result)
	if truncated_result != "1.007" {
		fmt.Println("truncated result is not 1.007")
	}

	// 2.
	res3 := []float64{30, 10, 30, 10}
	result, err = GetWPM(ctx, res1, res3, weights)
	if err != nil {
		fmt.Println(err)
	}
	if result < 1.0 {
		fmt.Println("result is smaller that 1.0")
	}
	truncated_result = fmt.Sprintf("%.3f", result)
	if truncated_result != "1.067" {
		fmt.Println("truncated result is not 1.067")
	}

	// 3.
	result, err = GetWPM(ctx, res2, res3, weights)
	if err != nil {
		fmt.Println(err)
	}
	if result < 1.0 {
		fmt.Println("result is smaller that 1.0")
	}
	truncated_result = fmt.Sprintf("%.3f", result)
	if truncated_result != "1.059" {
		fmt.Println("truncated result is not 1.059")
	}

	// 4.
	result, err = GetWPM(ctx, res3, res2, weights)
	if err != nil {
		fmt.Println(err)
	}
	if result > 1.0 {
		fmt.Println("result is larger that 1.0")
	}
	truncated_result = fmt.Sprintf("%.3f", result)
	if truncated_result != "0.944" {
		fmt.Println("truncated result is not 0.944")
	}
}
