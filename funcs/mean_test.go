package funcs

import (
	"github.com/ekimeel/timeseries/model"
	"github.com/ekimeel/timeseries/readers"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMean_Eval(t *testing.T) {
	reader := readers.CsvReader{Path: "../testdata/1x100.csv"}
	ts, err := reader.Read()
	assert.Nil(t, err)

	mean, err := ts.Eval(Mean{Dimension: "v0" })

	assert.Nil(t, err)
	assert.Equal(t, 1.5354999999999999, mean)
}

func TestMean_Eval_WithEmptyTimeSeries(t *testing.T) {
	ts := model.TimeSeries{}
	ts.SetDimensions([]string{"v0"})

	mean, err := ts.Eval(Mean{Dimension: "v0"})

	assert.Nil(t, err)
	assert.True(t, math.IsNaN(mean))
}
