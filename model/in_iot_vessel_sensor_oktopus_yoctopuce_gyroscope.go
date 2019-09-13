package model

import (
	"encoding/csv"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type (
	Dataset        []IOTVesselSensorOktopusYoctopuceGyroscope
	DeviceReadings []float64 //nolint
	EventTimes     []int64   //nolint
	Intervals      []int     //nolint
)
type IOTVesselSensorOktopusYoctopuceGyroscope struct {
	Topic     string  `json:"topic"`
	Partition int     `json:"partition"`
	Offset    int     `json:"offset"`
	TsType    string  `json:"tstype"`
	Ts        int64   `json:"ts"`
	Key       string  `json:"key"`
	Payload   Payload `json:"Payload"`
}

type Payload struct {
	Imo            string  `json:"imo"`
	System         string  `json:"system"`
	SubSystem      string  `json:"subSystem"`
	OriginatorType string  `json:"originatorType"`
	OriginatorID   string  `json:"originatorId"`
	EventTime      int64   `json:"eventTime"`
	AccountID      string  `json:"accountId"`
	Accelerometer  float64 `json:"accelerometer"`
	Compass        float64 `json:"compass"`
	Gyro           float64 `json:"gyro"`
	Magnetometer   float64 `json:"magnetometer"`
	TiltX          float64 `json:"tiltX"`
	TiltY          float64 `json:"tiltY"`
}

func NewIOTVesselSensorOktopusYoctopuceGyroscope() *IOTVesselSensorOktopusYoctopuceGyroscope {
	return &IOTVesselSensorOktopusYoctopuceGyroscope{}
}

func NewData() Dataset {
	var dataset Dataset
	return dataset
}

func (dataset *Dataset) JSONFileToStruct(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	d := json.NewDecoder(f)
	for {
		var v IOTVesselSensorOktopusYoctopuceGyroscope
		if err := d.Decode(&v); err == io.EOF {
			break // done decoding file
		} else if err != nil {
			log.Println(err.Error())
			return err
		}
		*dataset = append(*dataset, v)
	}
	return nil
}

func ExtractAccelerometerValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.Accelerometer)) //nolint
	}
	return Readings
}

func ExtractCompassValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.Compass)) //nolint
	}
	return Readings
}

func ExtractGyroValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.Gyro)) //nolint
	}
	return Readings
}

func ExtractMagnetometerValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.Magnetometer)) //nolint
	}
	return Readings
}

func ExtractTiltXValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.TiltX)) //nolint
	}
	return Readings
}

func ExtractTiltYValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.TiltY)) //nolint
	}

	return Readings
}

func ExtractEventTimes(dataset *Dataset) EventTimes {
	var eventTimes EventTimes
	int64AsIntValues := make([]int, len(*dataset))

	for _, eventTime := range *dataset {
		eventTimes = append(eventTimes, eventTime.Payload.EventTime)
	}

	// sort the event times
	for i, val := range eventTimes {
		int64AsIntValues[i] = int(val)
	}
	sort.Ints(int64AsIntValues)

	for i, val := range int64AsIntValues {
		eventTimes[i] = int64(val)
	}
	return eventTimes
}

func CalculateEventTimeIntervals(eventTimes EventTimes) Intervals {
	var intervals Intervals

	for i := 0; i < len(eventTimes)-1; i++ {
		interval := eventTimes[i+1] - eventTimes[i]
		intervals = append(intervals, int(interval))
	}
	return intervals
}

func CreateTable(eventTimes EventTimes, magnetometer, compass, accelerometer, gyro, tiltX, tiltY DeviceReadings) [][]string {
	var dataTable [][]string
	for i := 0; i <= len(eventTimes)-1; i++ {
		eventTimeString := time.Unix(eventTimes[i]/1000, 0).UTC().Format(time.RFC822)
		magnetometerValue := strconv.FormatFloat(magnetometer[i], 'f', -1, 64)
		compassValue := strconv.FormatFloat(compass[i], 'f', -1, 64)
		accelerometerValue := strconv.FormatFloat(accelerometer[i], 'f', -1, 64)
		tiltXValue := strconv.FormatFloat(tiltX[i], 'f', -1, 64)
		tiltYValue := strconv.FormatFloat(tiltY[i], 'f', -1, 64)
		gyroValue := strconv.FormatFloat(gyro[i], 'f', -1, 64)

		dataTable = append(dataTable, []string{eventTimeString, magnetometerValue, compassValue, accelerometerValue, gyroValue, tiltXValue, tiltYValue})
	}

	return dataTable
}

func ToCSVFile(dataTable [][]string, filename string) {
	// to csv
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range dataTable {
		err := writer.Write(value)
		if err != nil {
			log.Println(err.Error())
		}
	}
	log.Println("The file has bee successfully created: " + filename)
}

func RenderTable(dataTable [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"_______Event_Time_______", "Magnetometer", "Compass", "Accelerometer", "Gyro", "TiltX", "TiltY"})
	for _, v := range dataTable {
		table.Append(v)
	}
	table.SetColMinWidth(0, 75)
	table.SetRowLine(true)

	table.Render()

}
