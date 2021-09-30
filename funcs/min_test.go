package funcs

import (
	"github.com/ekimeel/timeseries/model"
	"github.com/ekimeel/timeseries/readers"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestMin_Eval(t *testing.T) {
	reader := readers.CsvReader{Path: "../testdata/1x100.csv"}
	ts, err := reader.Read()
	assert.Nil(t, err)

	min, err := ts.Eval(Min{Dimension: "v0" })

	assert.Nil(t, err)
	assert.Equal(t, 1.00, min)
}

func TestMin_Eval_WithEmptyTimeSeries(t *testing.T) {
	ts := model.TimeSeries{}
	ts.SetDimensions([]string{"v0"})

	min, err := ts.Eval(Min{Dimension: "v0" })

	assert.Nil(t, err)
	assert.True(t, math.IsNaN(min))
}

func TestMin_Eval_WithDimensionThatDoestExist(t *testing.T) {
	ts := model.TimeSeries{}
	ts.SetDimensions([]string{"v0"})
	ts.Add(time.Now().Unix(), []float64{1.0})

	min, err := ts.Eval(Min{Dimension: "v1" })

	assert.NotNil(t, err)
	assert.True(t, math.IsNaN(min))
}