package data

import (
	"fmt"
	"github.com/hermanschaaf/stats"
	"kafka-topic-analysis/mathematicalfunctions"
	"kafka-topic-analysis/topics"
)

const lambda = 0.0083333333

func AnalyseIntervals(intervals topics.Intervals) {
	// Interval analysis
	mean := stats.Mean(intervals)
	mode := stats.Mode(intervals)
	median := stats.Median(intervals)
	stdDev := stats.StandardDeviation(intervals)
	min, max := mathematicalfunctions.FindIntMinAndMax(intervals)
	//Calculate Probability based on exponential analysis
	probability1 := mathematicalfunctions.ProbabilityGreaterThan(lambda, 120)
	probability2 := mathematicalfunctions.ProbabilityLessThan(lambda, 120)
	probability3 := mathematicalfunctions.ProbabilityBetweenTwoValues(lambda, 120, 270)

	fmt.Printf(`
	----------------------------------- Basic Stats Analysis: Time Intervals -----------------------------------
	Time interval Mean: %v Seconds
	Time interval Mode: %v Seconds
	Time interval Median: %v Seconds
	Time interval Max Value: %v Seconds
	Time interval Min Value: %v Seconds
	Time interval Standard Deviation: %v Seconds
	
	`,
		float32(mean)/1000,
		mode/1000,
		median/1000,
		max/1000,
		min/1000,
		float32(stdDev)/1000,
	)
	fmt.Printf(`
	Poisson Distribution: Time Intervals
	The probability of a time interval being Greater than 2 minutes is %v 
	The probability of a time interval being Less than 2 minutes is %v
	The probability of a data interval being Between 2 & 4.5 minutes is %v
	`,
		probability1,
		probability2,
		probability3,
	)

}

func AnalyseReadings(name string, deviceReading topics.DeviceReadings, lowerBound, upperBound float64) {

	mean := mathematicalfunctions.Mean(deviceReading)
	mode := mathematicalfunctions.Mode(deviceReading)
	median := mathematicalfunctions.Median(deviceReading)
	stdDev := mathematicalfunctions.StandardDeviation(deviceReading)
	min, max := mathematicalfunctions.FindFloat64MinAndMax(deviceReading)
	//Calculate Probability based on exponential analysis
	probability1 := mathematicalfunctions.ProbabilityGreaterThan(lambda, upperBound)
	probability2 := mathematicalfunctions.ProbabilityLessThan(lambda, lowerBound)
	probability3 := mathematicalfunctions.ProbabilityBetweenTwoValues(lambda, lowerBound, upperBound)

	fmt.Printf(`
	----------------------------------- Basic Stats Analysis: %v Values -----------------------------------
	%v Mean: %v 
	%v Mode: %v 
	%v Median: %v 
	%v Max Value: %v 
	%v Min Value: %v 
	%v Standard Deviation: %v 
	
	`,
		name,
		name, float32(mean),
		name, mode,
		name, median,
		name, max,
		name, min,
		name, float32(stdDev),
	)
	fmt.Printf(`
	Poisson Distribution: %v Values
	The probability of a %v being Greater than %v is %v 
	The probability of a %v being Less than %v is %v
	The probability of a %v being Between %v & %v is %v
	`,
		name,
		name, upperBound, probability1,
		name, lowerBound, probability2,
		name, lowerBound, upperBound, probability3,
	)

}

func AnalyseReadingsNoProb(name string, deviceReading topics.DeviceReadings) {

	mean := mathematicalfunctions.Mean(deviceReading)
	mode := mathematicalfunctions.Mode(deviceReading)
	median := mathematicalfunctions.Median(deviceReading)
	stdDev := mathematicalfunctions.StandardDeviation(deviceReading)
	min, max := mathematicalfunctions.FindFloat64MinAndMax(deviceReading)

	fmt.Printf(`
	----------------------------------- Basic Stats Analysis: %v Values -----------------------------------
	%v Mean: %v 
	%v Mode: %v 
	%v Median: %v 
	%v Max Value: %v
	%v Min Value: %v 
	%v Standard Deviation: %v 
	
	`,
		name,
		name, float32(mean),
		name, mode,
		name, median,
		name, max,
		name, min,
		name, float32(stdDev),
	)

}
