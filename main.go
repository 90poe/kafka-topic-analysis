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

		CorrectedDataset := topics.RemoveDuplicates(dataset)
		results := topics.GetTopicValues(&CorrectedDataset)

		dataTable = topics.CreateDataTable(results)
		if env.Settings.CreateTable {
			data.RenderTable(dataTable)
		}
		if env.Settings.ToCSV {
			data.ToCSVFile(dataTable, env.Settings.CSVFilename)
		}

		data.AnalyseIntervals(results.Intervals)
		data.AnalyseReadings("Gyro", results.Gyro, 0, 0.3)
		data.AnalyseReadings("Magnetometer", results.Magnetometer, 0.5, 1.0)
		data.AnalyseReadings("Compass", results.Compass, 150, 200)
		data.AnalyseReadings("Accelerometer", results.Accelerometer, 0.2, 5.0)
		data.AnalyseReadingsNoProb("TiltX", results.TiltX)
		data.AnalyseReadingsNoProb("TiltY", results.TiltY)

	}
}
