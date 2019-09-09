package mathematicalfunctions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFloat64MinAndMax(t *testing.T) {
	//arrange
	var values = []float64{0.122262, 0.0124758, 0.1562322, 1.2273322, 26.7233265, 1.6332245, 2.8733365, 0.9822765, 0.72226282, 55.763354}

	//act
	min, max := FindFloat64MinAndMax(values)

	//assert
	assert.Equal(t, 0.0124758, min)
	assert.Equal(t, 55.763354, max)
}
