package mathematicalfunctions

import (
	"math"
	"sort"
)

// THIS PACKAGE CONTAINS USEFUL MATHEMATICAL FUNCTIONS FOR ANALYSING DATA COMING FROM THE TOPICS

// Find The Max and Min of an array of FLOAT64
func FindFloat64MinAndMax(a []float64) (min float64, max float64) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// Find The Max and Min of an array of Int
func FindIntMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// Functions to calculate the Exponential Distribution for Predictive Analysis
func ProbabilityGreaterThan(lambda, bigX float64) float64 {
	probability := math.Exp((-1) * lambda * bigX)

	return probability
}

func ProbabilityLessThan(lambda, bigX float64) float64 {
	probability := 1 - math.Exp((-1)*lambda*bigX)

	return probability
}

func ProbabilityBetweenTwoValues(lambda, bigX1, bigX2 float64) float64 {
	probability1 := 1 - math.Exp((-1)*lambda*bigX1)
	probability2 := 1 - math.Exp((-1)*lambda*bigX2)

	probability := probability2 - probability1
	return probability
}

// Mean returns the mean of an integer array as a float
func Mean(nums []float64) (mean float64) {
	if len(nums) == 0 {
		return 0.0
	}
	for _, n := range nums {
		mean += n
	}
	return mean / float64(len(nums))
}

// Median returns the median of an integer array as a float
func Median(nums []float64) (median float64) {
	if len(nums) == 0 {
		return 0.0
	}

	t := make([]float64, len(nums))
	copy(t, nums)

	sort.Float64s(t)
	l := len(t)
	if l == 0 { //nolint
		return 0.0
	} else if l%2 == 0 {
		median = Mean(t[l/2-1 : l/2+1])
	} else {
		median = t[l/2]
	}
	return median
}

// Mode returns the value of the most common element, or the smallest
// value if multiple elements satisfy this criteria.
func Mode(nums []float64) (mode float64) {
	if len(nums) == 0 {
		return 0.0
	}

	m := map[float64]float64{}
	var lowest, count float64
	lowest, count = 0, 0
	for _, n := range nums {
		m[n]++
		if m[n] > count || (m[n] == count && n < lowest) {
			count = m[n]
			lowest = n
		}
	}
	return lowest
}

// StandardDeviation returns the standard deviation of the slice
// as a float
func StandardDeviation(nums []float64) (dev float64) {
	if len(nums) == 0 {
		return 0.0
	}

	m := Mean(nums)
	for _, n := range nums {
		dev += (n - m) * (n - m)
	}
	dev = math.Pow(dev/float64(len(nums)), 0.5)
	return dev
}
