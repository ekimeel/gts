package transformations

import (
	"github.com/ekimeel/timeseries/funcs"
	"github.com/ekimeel/timeseries/model"
	"github.com/ekimeel/timeseries/writers"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)


func TestRollUp_Transform_WithMax(t *testing.T) {
	reader := model.CsvReader{Path: "../testdata/1x100.csv"}
	ts, _ := reader.Read()

	rollup, err := ts.Transform(&RollUp{ Duration: 15*time.Minute, Agg: funcs.Max{} })



	writer := writers.CsvWriter{Writer: os.Stdout}
	writer.Write(&rollup)

	assert.Nil(t, err)
	assert.Equal(t, 100, ts.Size())
	assert.Equal(t, 7, rollup.Size())
	assert.Equal(t, 1.9, *rollup.At(0,0))
	assert.Equal(t, 1.28, *rollup.At(1,0))
	assert.Equal(t, 1.43, *rollup.At(2,0))
	assert.Equal(t, 1.58, *rollup.At(3,0))
	assert.Equal(t, 1.73, *rollup.At(4,0))
	assert.Equal(t, 1.88, *rollup.At(5,0))
	assert.Equal(t, 1.99, *rollup.At(6,0))
}
