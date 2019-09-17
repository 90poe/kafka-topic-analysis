package topics

import (
	"encoding/json"
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

type Values struct {
	EventTimes    EventTimes
	Magnetometer  DeviceReadings
	Compass       DeviceReadings
	Accelerometer DeviceReadings
	Gyro          DeviceReadings
	TiltX         DeviceReadings
	TiltY         DeviceReadings
	Intervals     Intervals
}

func NewIOTVesselSensorOktopusYoctopuceGyroscope() *IOTVesselSensorOktopusYoctopuceGyroscope {
	return &IOTVesselSensorOktopusYoctopuceGyroscope{}
}

func NewData() Dataset {
	return Dataset{}
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

func RemoveDuplicates(s Dataset) Dataset {
	var newDataset Dataset
	oldDataSize := len(s)
	seen := make(map[IOTVesselSensorOktopusYoctopuceGyroscope]struct{}, oldDataSize)
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	newDataset = s[:j]
	newDataSize := len(newDataset)
	log.Printf("Function removed %v messages from the Dataset\n", oldDataSize-newDataSize)
	return newDataset

}

func GetTopicValues(dataset *Dataset) Values {
	accelerometerValues := getAccelerometerValues(dataset)
	compassValues := getCompassValues(dataset)
	gyroValues := getGyroValues(dataset)
	magnetometerValues := getMagnetometerValues(dataset)
	tiltXValues := getTiltXValues(dataset)
	tiltYValues := getTiltYValues(dataset)
	eventTimes := getEventTimes(dataset)
	intervals := calculateEventTimeIntervals(eventTimes)

	return Values{
		EventTimes:    eventTimes,
		Magnetometer:  magnetometerValues,
		Compass:       compassValues,
		Accelerometer: accelerometerValues,
		Gyro:          gyroValues,
		TiltX:         tiltXValues,
		TiltY:         tiltYValues,
		Intervals:     intervals,
	}
}

func CreateDataTable(results Values) [][]string {
	var dataTable [][]string
	for i := 0; i <= len(results.EventTimes)-1; i++ {
		eventTimeString := time.Unix(results.EventTimes[i]/1000, 0).UTC().Format(time.RFC822)
		magnetometerValue := strconv.FormatFloat(results.Magnetometer[i], 'f', -1, 64)
		compassValue := strconv.FormatFloat(results.Compass[i], 'f', -1, 64)
		accelerometerValue := strconv.FormatFloat(results.Accelerometer[i], 'f', -1, 64)
		tiltXValue := strconv.FormatFloat(results.TiltX[i], 'f', -1, 64)
		tiltYValue := strconv.FormatFloat(results.TiltY[i], 'f', -1, 64)
		gyroValue := strconv.FormatFloat(results.Gyro[i], 'f', -1, 64)

		dataTable = append(dataTable, []string{eventTimeString, gyroValue, magnetometerValue, compassValue, accelerometerValue, tiltXValue, tiltYValue})
	}

	return dataTable
}

func getAccelerometerValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.Accelerometer)) //nolint
	}
	return Readings
}

func getCompassValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.Compass)) //nolint
	}
	return Readings
}

func getGyroValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.Gyro)) //nolint
	}
	return Readings
}

func getMagnetometerValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.Magnetometer)) //nolint
	}
	return Readings
}

func getTiltXValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.TiltX)) //nolint
	}
	return Readings
}

func getTiltYValues(dataset *Dataset) DeviceReadings {
	var Readings DeviceReadings
	for _, reading := range *dataset {
		Readings = append(Readings, float64(reading.Payload.TiltY)) //nolint
	}

	return Readings
}

func getEventTimes(dataset *Dataset) EventTimes {
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

func calculateEventTimeIntervals(eventTimes EventTimes) Intervals {
	var intervals Intervals

	for i := 0; i < len(eventTimes)-1; i++ {
		interval := eventTimes[i+1] - eventTimes[i]
		intervals = append(intervals, int(interval))
	}
	return intervals
}
