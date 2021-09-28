package funcs

import (
	"github.com/ekimeel/timeseries/model"
	"github.com/ekimeel/timeseries/readers"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestMax_Eval(t *testing.T) {
	reader := readers.CsvReader{Path: "../testdata/1x100.csv"}
	ts, err := reader.Read()
	assert.Nil(t, err)

	max, err := ts.Eval(Max{Dimension: "v0" })

	assert.Nil(t, err)
	assert.Equal(t, 1.99, max)
}

func TestMax_Eval_WithEmptyTimeSeries(t *testing.T) {
	ts := model.TimeSeries{}
	ts.SetDimensions([]string{"v0"})

	max, err := ts.Eval(Max{Dimension: "v0" })

	assert.Nil(t, err)
	assert.True(t, math.IsNaN(max))
}

func TestMax_Eval_WithDimensionThatDoestExist(t *testing.T) {
	ts := model.TimeSeries{}
	ts.SetDimensions([]string{"v0"})
	ts.Add(time.Now().Unix(), []float64{1.0})

	max, err := ts.Eval(Max{Dimension: "v1" })

	assert.NotNil(t, err)
	assert.True(t, math.IsNaN(max))
}