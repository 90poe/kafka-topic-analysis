package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"kafka-topic-analysis/data"
	"kafka-topic-analysis/env"
	"kafka-topic-analysis/topics"
	"log"
	"os"
)

var dataTable [][]string
var values topics.Values

func main() {
	env.Init()

	switch env.Settings.Operation {
	case env.OperationDescribe:
		fmt.Println(env.DESCRIPTION)
		pflag.PrintDefaults()
		os.Exit(0)

	case env.OperationAnalyse:
		dataset := topics.NewData()
		err := dataset.JSONFileToStruct(env.Settings.JSONFilePath)
		if err != nil {
			log.Printf("Failed to parse file from location: %v", env.Settings.JSONFilePath)
		}

		if env.Settings.RemoveDuplicates {
			log.Println("Removing Duplicates...")
			CorrectedDataset := topics.RemoveDuplicates(dataset)
			values = topics.GetTopicValues(&CorrectedDataset)
		} else {
			log.Println("Leaving the Duplicates in the set...")
			values = topics.GetTopicValues(&dataset)
		}

		dataTable = topics.CreateDataTable(values)
		if env.Settings.CreateTable {
			data.RenderTable(dataTable)
		}
		if env.Settings.ToCSV {
			data.ToCSVFile(dataTable, env.Settings.CSVFilename)
		}

		data.AnalyseIntervals(values.Intervals)
		data.AnalyseReadings("Gyro", values.Gyro, 0, 0.3)
		data.AnalyseReadings("Magnetometer", values.Magnetometer, 0.5, 1.0)
		data.AnalyseReadings("Compass", values.Compass, 150, 200)
		data.AnalyseReadings("Accelerometer", values.Accelerometer, 0.2, 5.0)
		data.AnalyseReadingsNoProb("TiltX", values.TiltX)
		data.AnalyseReadingsNoProb("TiltY", values.TiltY)

	}
}
