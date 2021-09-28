package funcs

import (
	"github.com/ekimeel/timeseries/model"
	"github.com/ekimeel/timeseries/readers"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestStdDev_Eval(t *testing.T) {
	reader := readers.CsvReader{Path: "../testdata/1x100.csv"}
	ts, err := reader.Read()
	assert.Nil(t, err)

	mean, err := ts.Eval(StdDev{Dimension: "v0" })

	assert.Nil(t, err)
	assert.Equal(t, 0.2642058856270995, mean)
}

func TestStdDev_Eval_WithEmptyTimeSeries(t *testing.T) {
	ts := model.TimeSeries{}
	ts.SetDimensions([]string{"v0"})

	mean, err := ts.Eval(StdDev{Dimension: "v0"})

	assert.Nil(t, err)
	assert.True(t, math.IsNaN(mean))
}

func TestStdDev_Eval_WithDimensionThatDoestExist(t *testing.T) {
	ts := model.TimeSeries{}
	ts.SetDimensions([]string{"v0"})

	sd, err := ts.Eval(StdDev{Dimension: "v1"})

	assert.NotNil(t, err)
	assert.True(t, math.IsNaN(sd))
}
