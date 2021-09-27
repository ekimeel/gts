package funcs

import (
	"github.com/ekimeel/timeseries/model"
	"github.com/ekimeel/timeseries/readers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum_Eval(t *testing.T) {
	reader := readers.CsvReader{Path: "../testdata/1x100.csv"}
	ts, err := reader.Read()
	assert.Nil(t, err)

	sum, err := ts.Eval(Sum{Dimension: "v0" })

	assert.Nil(t, err)
	assert.Equal(t, 153.54999999999998, sum)
}

func TestSum_Eval_WithEmptyTimeSeries(t *testing.T) {
	ts := model.TimeSeries{}
	ts.SetDimensions([]string{"v0"})

	sum, err := ts.Eval(Sum{Dimension: "v0"})

	assert.Nil(t, err)
	assert.Equal(t, 0.0, sum)
}