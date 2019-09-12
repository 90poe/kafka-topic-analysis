package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	exampleJSON = `{
  "topic": "in_iot_vessel_sensor_oktopus_yoctopuce_gyroscope",
  "partition": 14,
  "offset": 213646,
  "tstype": "create",
  "ts": 1567687080000,
  "key": "9505998",
  "Payload": {
    "imo": "9505998",
    "system": "YOCTOPUCE",
    "subSystem": "NONE",
    "originatorType": "VESSEL_SYSTEM",
    "originatorId": "",
    "eventTime": 1567687080000,
    "accountId": "8a396ee1-a16b-4cc4-8b54-d3de465b8fc8",
    "accelerometer": 0.984000027179718,
    "compass": 284.79998779296875,
    "gyro": 0.10000000149011612,
    "magnetometer": 0.3019999861717224,
    "tiltX": -57.29999923706055,
    "tiltY": -2.5
  }
}`

	exampleJSON2 = `{
  "topic": "in_iot_vessel_sensor_oktopus_yoctopuce_gyroscope",
  "partition": 14,
  "offset": 213648,
  "tstype": "create",
  "ts": 1567687320000,
  "key": "9505998",
  "payload": {
    "imo": "9505998",
    "system": "YOCTOPUCE",
    "subSystem": "NONE",
    "originatorType": "VESSEL_SYSTEM",
    "originatorId": "",
    "eventTime": 1567687320000,
    "accountId": "8a396ee1-a16b-4cc4-8b54-d3de465b8fc8",
    "accelerometer": 0.9879999756813049,
    "compass": 284.70001220703125,
    "gyro": 0.10000000149011612,
    "magnetometer": 0.27300000190734863,
    "tiltX": -57.099998474121094,
    "tiltY": -2.5
  }
}
{
  "topic": "in_iot_vessel_sensor_oktopus_yoctopuce_gyroscope",
  "partition": 14,
  "offset": 213649,
  "tstype": "create",
  "ts": 1567687440000,
  "key": "9505998",
  "payload": {
    "imo": "9505998",
    "system": "YOCTOPUCE",
    "subSystem": "NONE",
    "originatorType": "VESSEL_SYSTEM",
    "originatorId": "",
    "eventTime": 1567687440000,
    "accountId": "8a396ee1-a16b-4cc4-8b54-d3de465b8fc8",
    "accelerometer": 0.9819999933242798,
    "compass": 284.70001220703125,
    "gyro": 0.10000000149011612,
    "magnetometer": 0.28299999237060547,
    "tiltX": -57.20000076293945,
    "tiltY": -2.4000000953674316
  }
}
`

	emptyJSON = `{"topic":"","partition":0,"offset":0,"tstype":"","ts":0,"key":"","Payload":{"imo":"","system":"","subSystem":"","originatorType":"","originatorId":"","eventTime":0,"accountId":"","accelerometer":0,"compass":0,"gyro":0,"magnetometer":0,"tiltX":0,"tiltY":0}}`
)

func TestNewIOTVesselSensorOktopusYoctopuceGyroscope(t *testing.T) {
	//act
	j := NewIOTVesselSensorOktopusYoctopuceGyroscope()

	bytes, err := json.Marshal(&j)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	//assert
	assert.Equal(t, emptyJSON, string(bytes))
}

func TestNewData(t *testing.T) {
	// Arrange
	var expectedDataset Dataset

	// Act
	Dt := NewData()

	// Assert
	assert.Equal(t, expectedDataset, Dt)
}

func TestIOTVesselSensorOktopusYoctopuceGyroscope_JsonToStruct(t *testing.T) {
	// Arrange
	Dt := NewData()

	// Act
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")

	// Assert
	// test first dataset
	assert.Equal(t, 0.9750000238418579, Dt[0].Payload.Accelerometer)
	assert.Equal(t, 197.89999389648438, Dt[0].Payload.Compass)
	assert.Equal(t, 0.10000000149011612, Dt[0].Payload.Gyro)
	assert.Equal(t, 0.7789999842643738, Dt[0].Payload.Magnetometer)
	assert.Equal(t, -0.20000000298023224, Dt[0].Payload.TiltX)
	assert.Equal(t, 2.299999952316284, Dt[0].Payload.TiltY)

	// test second dataset
	assert.Equal(t, 0.9639999866485596, Dt[1].Payload.Accelerometer)
	assert.Equal(t, 197.1999969482422, Dt[1].Payload.Compass)
	assert.Equal(t, float64(0), Dt[1].Payload.Gyro)
	assert.Equal(t, 0.7630000114440918, Dt[1].Payload.Magnetometer)
	assert.Equal(t, -0.20000000298023224, Dt[1].Payload.TiltX)
	assert.Equal(t, -0.10000000149011612, Dt[1].Payload.TiltY)

	// test third dataset
	assert.Equal(t, 0.9670000076293945, Dt[2].Payload.Accelerometer)
	assert.Equal(t, float64(196), Dt[2].Payload.Compass)
	assert.Equal(t, 0.10000000149011612, Dt[2].Payload.Gyro)
	assert.Equal(t, 0.7839999794960022, Dt[2].Payload.Magnetometer)
	assert.Equal(t, -0.10000000149011612, Dt[2].Payload.TiltX)
	assert.Equal(t, 0.10000000149011612, Dt[2].Payload.TiltY)
	assert.Nil(t, err)

}

func TestExtractAccelerometerValues(t *testing.T) {
	// Arrange
	Dt := NewData()
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")
	var expectedValues = DeviceReadings{0.9750000238418579, 0.9639999866485596, 0.9670000076293945, 0.9760000109672546, 0.9629999995231628, 0.9539999961853027, 0.9369999766349792, 0.9700000286102295, 0.9200000166893005, 0.9390000104904175}

	// Act
	values := ExtractAccelerometerValues(&Dt)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedValues, values)
}

func TestExtractGyroValues(t *testing.T) {
	// Arrange
	Dt := NewData()
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")
	var expectedValues = DeviceReadings{0.10000000149011612, 0, 0.10000000149011612, 0, 0.10000000149011612, 0, 0, 0.20000000298023224, 0.10000000149011612, 0}

	// Act
	values := ExtractGyroValues(&Dt)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedValues, values)
}

func TestExtractMagnetometerValues(t *testing.T) {
	// Arrange
	Dt := NewData()
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")
	var expectedValues = DeviceReadings{0.7789999842643738, 0.7630000114440918, 0.7839999794960022, 0.7710000276565552, 0.7799999713897705, 0.7910000085830688, 0.7839999794960022, 0.7820000052452087, 0.7789999842643738, 0.7749999761581421}

	// Act
	values := ExtractMagnetometerValues(&Dt)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedValues, values)
}

func TestExtractCompassValues(t *testing.T) {
	// Arrange
	Dt := NewData()
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")
	var expectedValues = DeviceReadings{197.89999389648438, 197.1999969482422, 196, 194.1999969482422, 193.60000610351562, 192.8000030517578, 186.8000030517578, 179.89999389648438, 194.60000610351562, 194.3000030517578}

	// Act
	values := ExtractCompassValues(&Dt)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedValues, values)
}

func TestExtractTiltXValues(t *testing.T) {
	// Arrange
	Dt := NewData()
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")
	var expectedValues = DeviceReadings{-0.20000000298023224, -0.20000000298023224, -0.10000000149011612, -0.20000000298023224, -0.10000000149011612, -0.20000000298023224, -0.10000000149011612, -0.10000000149011612, 0.20000000298023224, 0.10000000149011612}

	// Act
	values := ExtractTiltXValues(&Dt)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedValues, values)
}

func TestExtractTiltYValues(t *testing.T) {
	// Arrange
	Dt := NewData()
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")
	var expectedValues = DeviceReadings{2.299999952316284, -0.10000000149011612, 0.10000000149011612, 0, -0.699999988079071, -0.6000000238418579, -2.0999999046325684, -4.900000095367432, -0.6000000238418579, -1.2999999523162842}

	// Act
	values := ExtractTiltYValues(&Dt)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedValues, values)
}

func TestExtractEventTimes(t *testing.T) {
	// Arrange
	Dt := NewData()
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")
	var expectedValues = EventTimes{1567784040025, 1567784160073, 1567784280113, 1567784400162, 1567784520205, 1567784640244, 1567784760292, 1567784880332, 1567785000377, 1567785120413}

	// Act
	values := ExtractEventTimes(&Dt)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedValues, values)
}

func TestCalculateEventTimeIntervals(t *testing.T) {
	// Arrange
	Dt := NewData()
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")
	var expectedValues = Intervals{120048, 120040, 120049, 120043, 120039, 120048, 120040, 120045, 120036}
	eventTimes := ExtractEventTimes(&Dt)

	// Act
	values := CalculateEventTimeIntervals(eventTimes)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedValues, values)
}

func TestCreateTable(t *testing.T) {
	// Arrange
	Dt := NewData()
	err := Dt.JSONFileToStruct("../testNewlineJSONOutput")
	accelerometerValues := ExtractAccelerometerValues(&Dt)
	compassValues := ExtractCompassValues(&Dt)
	gyroValues := ExtractGyroValues(&Dt)
	magnetometerValues := ExtractMagnetometerValues(&Dt)
	tiltXValues := ExtractTiltXValues(&Dt)
	tiltYValues := ExtractTiltYValues(&Dt)
	eventTimes := ExtractEventTimes(&Dt)

	// Act
	CreateTable(eventTimes, magnetometerValues, compassValues, accelerometerValues, gyroValues, tiltXValues, tiltYValues)

	// Assert
	assert.Nil(t, err)

}
