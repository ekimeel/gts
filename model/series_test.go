package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeries_GetDimensions(t *testing.T) {

	var series Series
	series.Set("v0", 1.0)
	series.Set("v1", 2.0)
	series.Set("v2", 3.0)

	assert.Equal(t, 3, len(series.GetDimensions()))
	assert.Equal(t, "v0", series.GetDimensions()[0])
	assert.Equal(t, "v1", series.GetDimensions()[1])
	assert.Equal(t, "v2", series.GetDimensions()[2])

}

func TestTimeSeries_GetDimension_WhenEmpty(t *testing.T) {

	var series Series
	assert.Equal(t, 0, len(series.GetDimensions()))
}
