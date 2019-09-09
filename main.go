package main

import (
	"fmt"
	"github.com/hermanschaaf/stats"
	"kafka-topic-analysis/env"
	"kafka-topic-analysis/log"
	"kafka-topic-analysis/mathematicalfunctions"
	"kafka-topic-analysis/model"
)

const lambda = 0.0083333333

func main() {
	dataset := model.NewData()
	err := dataset.JSONFileToStruct("/Users/moh/go/src/kafka-topic-analysis/testNewlineJSONOutput")
	//err := dataset.JSONFileToStruct(env.Settings.JSONFilePath)
	if err != nil {
		log.ParseFileError().Writef("Failed to parse file from location: %v", env.Settings.JSONFilePath)
	}

	accelerometerValues := model.ExtractAccelerometerValues(&dataset)
	compassValues := model.ExtractCompassValues(&dataset)
	gyroValues := model.ExtractGyroValues(&dataset)
	magnetometerValues := model.ExtractMagnetometerValues(&dataset)
	tiltXValues := model.ExtractTiltXValues(&dataset)
	tiltYValues := model.ExtractTiltYValues(&dataset)
	eventTimes := model.ExtractEventTimes(&dataset)
	intervals := model.CalculateEventTimeIntervals(eventTimes)

	model.CreateTable(eventTimes, magnetometerValues, compassValues, accelerometerValues, gyroValues, tiltXValues, tiltYValues)

	// Interval analysis
	intervalMean := stats.Mean(intervals)
	intervalMode := stats.Mode(intervals)
	intervalMedian := stats.Median(intervals)
	intervalStdDev := stats.StandardDeviation(intervals)
	intervalMin, intervalMax := mathematicalfunctions.FindIntMinAndMax(intervals)
	//Calculate Probability based on exponential analysis
	intervalProbability1 := mathematicalfunctions.ProbabilityGreaterThan(lambda, 120)
	intervalProbability2 := mathematicalfunctions.ProbabilityLessThan(lambda, 120)
	intervalProbability3 := mathematicalfunctions.ProbabilityBetweenTwoValues(lambda, 120, 180)
	fmt.Println("\n----------------------------------- Basic Stats Analysis: Time Intervals -----------------------------------")
	fmt.Printf("Time interval Mean: %v Seconds \n", float32(intervalMean)/1000)
	fmt.Printf("Time interval Mode: %v Seconds\n", intervalMode/1000)
	fmt.Printf("Time interval Median: %v Seconds\n", intervalMedian/1000)
	fmt.Printf("Time interval Max Value: %v Seconds\n", intervalMax/1000)
	fmt.Printf("Time interval Min Value: %v Seconds\n", intervalMin/1000)
	fmt.Printf("Time interval Standard Deviation: %v\n", float32(intervalStdDev)/1000)
	fmt.Println("------ Poisson Distribution ------")
	fmt.Printf("\nThe probability of a time interval being Greater than 2 minutes is %v ( %v percent)", intervalProbability1, int(intervalProbability1*100))
	fmt.Printf("\nThe probability of a time interval being Equal to 2 minutes is %v ( %v percent)", intervalProbability2, int(intervalProbability2*100))
	fmt.Printf("\nThe probability of a data interval being Between 2 & 3 minutes is %v ( %v percent)", intervalProbability3, int(intervalProbability3*100))

	// Gyro analysis
	gyroMean := mathematicalfunctions.Mean(gyroValues)
	gyroMode := mathematicalfunctions.Mode(gyroValues)
	gyroMedian := mathematicalfunctions.Median(gyroValues)
	gyroStdDev := mathematicalfunctions.StandardDeviation(gyroValues)
	gyroMin, gyroMax := mathematicalfunctions.FindFloat64MinAndMax(gyroValues)
	fmt.Println("\n----------------------------------- Basic Stats Analysis: Gyro Readings -----------------------------------")
	fmt.Printf("Gyro Mean: %v °/s \n", gyroMean)
	fmt.Printf("Gyro Mode: %v °/s\n", gyroMode)
	fmt.Printf("Gyro Median: %v °/s\n", gyroMedian)
	fmt.Printf("Gyro Max Value: %v °/s\n", gyroMax)
	fmt.Printf("Gyro Min Value: %v °/s\n", gyroMin)
	fmt.Printf("Gyro Standard Deviation: %v\n", gyroStdDev)

	// Accelerometer analysis
	accelerometerMean := mathematicalfunctions.Mean(accelerometerValues)
	accelerometerMode := mathematicalfunctions.Mode(accelerometerValues)
	accelerometerMedian := mathematicalfunctions.Median(accelerometerValues)
	accelerometerStdDev := mathematicalfunctions.StandardDeviation(accelerometerValues)
	accelerometerMin, accelerometerMax := mathematicalfunctions.FindFloat64MinAndMax(accelerometerValues)
	fmt.Println("\n----------------------------------- Basic Stats Analysis: Accelerometer Readings -----------------------------------")
	fmt.Printf("Accelerometer Mean: %v g \n", accelerometerMean)
	fmt.Printf("Accelerometer Mode: %v g\n", accelerometerMode)
	fmt.Printf("Accelerometer Median: %v g\n", accelerometerMedian)
	fmt.Printf("Accelerometer Max Value: %v g\n", accelerometerMax)
	fmt.Printf("Accelerometer Min Value: %v g\n", accelerometerMin)
	fmt.Printf("Accelerometer Standard Deviation: %v\n", accelerometerStdDev)

	// Compass analysis
	compassMean := mathematicalfunctions.Mean(compassValues)
	compassMode := mathematicalfunctions.Mode(compassValues)
	compassMedian := mathematicalfunctions.Median(compassValues)
	compassStdDev := mathematicalfunctions.StandardDeviation(compassValues)
	compassMin, compassMax := mathematicalfunctions.FindFloat64MinAndMax(compassValues)
	fmt.Println("\n----------------------------------- Basic Stats Analysis: Compass Readings -----------------------------------")
	fmt.Printf("Compass Mean: %v°\n", compassMean)
	fmt.Printf("Compass Mode: %v°\n", compassMode)
	fmt.Printf("Compass Median: %v°\n", compassMedian)
	fmt.Printf("Compass Max Value: %v°\n", compassMax)
	fmt.Printf("Compass Min Value: %v°\n", compassMin)
	fmt.Printf("Compass Standard Deviation: %v\n", compassStdDev)

	// Magnetometer analysis
	magnetometerMean := mathematicalfunctions.Mean(magnetometerValues)
	magnetometerMode := mathematicalfunctions.Mode(magnetometerValues)
	magnetometerMedian := mathematicalfunctions.Median(magnetometerValues)
	magnetometerStdDev := mathematicalfunctions.StandardDeviation(magnetometerValues)
	magnetometerMin, magnetometerMax := mathematicalfunctions.FindFloat64MinAndMax(magnetometerValues)
	fmt.Println("\n----------------------------------- Basic Stats Analysis: Magnetometer Readings -----------------------------------")
	fmt.Printf("Magnetometer Mean: %v gauss \n", magnetometerMean)
	fmt.Printf("Magnetometer Mode: %v gauss\n", magnetometerMode)
	fmt.Printf("Magnetometer Median: %v gauss\n", magnetometerMedian)
	fmt.Printf("Magnetometer Max Value: %v gauss\n", magnetometerMax)
	fmt.Printf("Magnetometer Min Value: %v gauss\n", magnetometerMin)
	fmt.Printf("Magnetometer Standard Deviation: %v\n", magnetometerStdDev)

	// TiltX analysis
	tiltXMean := mathematicalfunctions.Mean(tiltXValues)
	tiltXMode := mathematicalfunctions.Mode(tiltXValues)
	tiltXMedian := mathematicalfunctions.Median(tiltXValues)
	tiltXStdDev := mathematicalfunctions.StandardDeviation(tiltXValues)
	tiltXMin, tiltXMax := mathematicalfunctions.FindFloat64MinAndMax(tiltXValues)
	fmt.Println("\n----------------------------------- Basic Stats Analysis: TiltX Readings -----------------------------------")
	fmt.Printf("TiltX Mean: %v° \n", tiltXMean)
	fmt.Printf("TiltX Mode: %v°\n", tiltXMode)
	fmt.Printf("TiltX Median: %v°\n", tiltXMedian)
	fmt.Printf("TiltX Max Value: %v°\n", tiltXMax)
	fmt.Printf("TiltX Min Value: %v°\n", tiltXMin)
	fmt.Printf("TiltX Standard Deviation: %v\n", tiltXStdDev)

	// TiltY analysis
	tiltYMean := mathematicalfunctions.Mean(tiltYValues)
	tiltYMode := mathematicalfunctions.Mode(tiltYValues)
	tiltYMedian := mathematicalfunctions.Median(tiltYValues)
	tiltYStdDev := mathematicalfunctions.StandardDeviation(tiltYValues)
	tiltYMin, tiltYMax := mathematicalfunctions.FindFloat64MinAndMax(tiltYValues)
	fmt.Println("\n----------------------------------- Basic Stats Analysis: TiltY Readings -----------------------------------")
	fmt.Printf("TiltY Mean: %v° \n", tiltYMean)
	fmt.Printf("TiltY Mode: %v°\n", tiltYMode)
	fmt.Printf("TiltY Median: %v°\n", tiltYMedian)
	fmt.Printf("TiltY Max Value: %v°\n", tiltYMax)
	fmt.Printf("TiltY Min Value: %v°\n", tiltYMin)
	fmt.Printf("TiltY Standard Deviation: %v\n", tiltYStdDev)
}
