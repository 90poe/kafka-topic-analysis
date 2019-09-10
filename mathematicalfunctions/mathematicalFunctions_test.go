package mathematicalfunctions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const lambda = 0.0083333333

func TestFindFloat64MinAndMax(t *testing.T) {
	//arrange
	var values = []float64{0.122262, 0.0124758, 0.1562322, 1.2273322, 26.7233265, 1.6332245, 2.8733365, 0.9822765, 0.72226282, 55.763354}
	var expectedMin, expectedMax = 0.0124758, 55.763354

	//act
	min, max := FindFloat64MinAndMax(values)

	//assert
	assert.Equal(t, expectedMin, min)
	assert.Equal(t, expectedMax, max)
}

func TestFindIntMinAndMax(t *testing.T) {
	//arrange
	var values = []int{12, 65, 13, 3, 23, 25, 11, 55, 6, 1}
	var expectedMin, expectedMax = 1, 65

	//act
	min, max := FindIntMinAndMax(values)

	//assert
	assert.Equal(t, expectedMin, min)
	assert.Equal(t, expectedMax, max)
}

func TestMode(t *testing.T) {
	//arrange
	var values = []float64{0.122262, 0.0124758, 0.1562322, 1.2273322, 26.7233265, 1.6332245, 2.8733365, 0.9822765, 0.72226282, 55.763354, 1.6332245}
	var expectedMode = 1.6332245
	//act
	mode := Mode(values)

	//assert
	assert.Equal(t, expectedMode, mode)

}

func TestMedian(t *testing.T) {
	//arrange
	var values = []float64{0.122262, 0.0124758, 0.1562322, 1.2273322, 26.7233265, 1.6332245, 2.8733365, 0.9822765, 0.72226282, 55.763354, 1.6332245}
	var expectedMedian = 1.2273322

	//act
	median := Median(values)

	//assert
	assert.Equal(t, expectedMedian, median)

}

func TestMean(t *testing.T) {
	//arrange
	var values = []float64{0.122262, 0.0124758, 0.1562322, 1.2273322, 26.7233265, 1.6332245, 2.8733365, 0.9822765, 0.72226282, 55.763354, 1.6332245}
	var expectedMean = 8.349937047272725

	//act
	mean := Mean(values)

	//assert
	assert.Equal(t, expectedMean, mean)
}

func TestStandardDeviation(t *testing.T) {
	//arrange
	var values = []float64{0.122262, 0.0124758, 0.1562322, 1.2273322, 26.7233265, 1.6332245, 2.8733365, 0.9822765, 0.72226282, 55.763354, 1.6332245}
	var expectedStdDev = 16.715195012268076
	//act
	stdDev := StandardDeviation(values)

	//assert
	assert.Equal(t, expectedStdDev, stdDev)
}

func TestProbabilityLessThan(t *testing.T) {
	// Arrange
	var bigX = float64(120) //120 seconds
	var expectedProbability = 0.63212055735704

	//Act
	probability := ProbabilityLessThan(lambda, bigX) // Probability values are Less than or Equal to 2 mins

	assert.Equal(t, expectedProbability, probability)
}

func TestProbabilityGreaterThan(t *testing.T) {
	// Arrange
	var bigX = float64(120) //120 seconds
	var expectedProbability = 0.36787944264296

	//Act
	probability := ProbabilityGreaterThan(lambda, bigX) // Probability values are Less than or Equal to 2 mins

	assert.Equal(t, expectedProbability, probability)
}

func TestProbabilityBetweenTwoValues(t *testing.T) {
	// Arrange
	var bigX1 = float64(120) //120 seconds
	var bigX2 = float64(180) //120 seconds
	var expectedProbability = 0.14474928115574925

	//Act
	probability := ProbabilityBetweenTwoValues(lambda, bigX1, bigX2) // Probability values are Less than or Equal to 2 mins

	assert.Equal(t, expectedProbability, probability)
}

// unhappy path tests
func TestMedian2(t *testing.T) {
	//arrange
	var values []float64
	var expectedMedian = 0.0

	//act
	median := Median(values)

	//assert
	assert.Equal(t, expectedMedian, median)
}

func TestMode2(t *testing.T) {
	//arrange
	var values []float64
	var expectedMode = 0.0

	//act
	mean := Mode(values)

	//assert
	assert.Equal(t, expectedMode, mean)
}

func TestMean2(t *testing.T) {
	//arrange
	var values []float64
	var expectedMean = 0.0

	//act
	mean := Mean(values)

	//assert
	assert.Equal(t, expectedMean, mean)
}

func TestStandardDeviation2(t *testing.T) {
	//arrange
	var values []float64
	var expectedStdDev = 0.0
	//act
	stdDev := StandardDeviation(values)

	//assert
	assert.Equal(t, expectedStdDev, stdDev)
}
