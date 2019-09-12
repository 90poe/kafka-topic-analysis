package env

import (
	"fmt"
	"github.com/spf13/pflag"
	"log"
	"os"
)

//VERSION of the Program
const (
	VERSION     = "v0.1.0"
	DESCRIPTION = "THIS TOOL HAS BEEN DESIGNED TO TAKE NEWLINE DELIMITED JSON OUTPUT FROM KAFKACAT, AND \nPROVIDE AN ANALYSIS OF THE EVENT TIMES, ACCELEROMETER, GYRO, COMPASS, MAGNETOMETER, TILTX \nAND Y READINGS AND SOME BASIC PROBABILITY ANALYSIS OF WHERE THE DATA IS PREDICTED TIO LIE BASED ON POISSON DISTRIBUTION." //nolint
)

type Config struct {
	JSONFilePath string
	CreateTable  bool
	ToCSV        bool
	Operation    string
}

var Settings *Config

//List of operations understood by a program
const (
	OperationAnalyse  = "analyse"
	OperationDescribe = "describe"
)

func (c *Config) verifyOperations() { //nolint
	var errorMsg string
	switch c.Operation {
	case OperationAnalyse:
		if len(c.JSONFilePath) == 0 {
			errorMsg = "The path to the JSON output file from kafkacat must be provided."
		}
	case OperationDescribe:
		//nothing is needed for this case
	default:
		errorMsg = fmt.Sprintf("Unknown operation: %v", c.Operation)
	}
	if len(errorMsg) != 0 {
		log.Println(errorMsg)
		pflag.PrintDefaults()
		os.Exit(1)
	}
}

func Init() {
	version := pflag.BoolP("version", "v", false, "Prints the current version")
	operation := pflag.String("operation", "", fmt.Sprintf(`Operation to perform:
  %s - Analyse a set of results and the time gaps between the data
  %s - Describe the application`,
		OperationAnalyse,
		OperationDescribe))

	// Compulsory arg flags
	jsonFilepath := pflag.String("json-filepath", "", "The path to the JSON output file from kafkacat")
	createTable := pflag.Bool("create-table", false, "Creates a table")
	toCSV := pflag.Bool("toCSV", false, "Output to a CSV File: output.csv")

	// Optional arg flags
	pflag.Parse()

	if *version {
		fmt.Printf("Version: %s\n", VERSION)
		os.Exit(0)
	}

	Settings = &Config{
		JSONFilePath: *jsonFilepath,
		CreateTable:  *createTable,
		ToCSV:        *toCSV,
		Operation:    *operation,
	}

	Settings.verifyOperations()
}
