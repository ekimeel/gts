package funcs

import (
	"github.com/ekimeel/timeseries/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSum_Compute(t *testing.T) {
	ts := model.TimeSeries{}

	start := time.Now()
	ts.SetDimensions([]string{"one", "two", "three"})
	ts.Add(start.Add(1*time.Minute).Unix(), []float64{1.0, 2.0, 3.0})
	ts.Add(start.Add(2*time.Minute).Unix(), []float64{1.1, 2.1, 3.1})
	ts.Add(start.Add(3*time.Minute).Unix(), []float64{1.2, 2.2, 3.2})
	ts.Add(start.Add(4*time.Minute).Unix(), []float64{1.3, 2.3, 3.3})
	ts.Add(start.Add(5*time.Minute).Unix(), []float64{1.4, 2.4, 3.4})

	sum, err := ts.ComputeValue(Sum{Dimension: "two"})
	assert.Nil(t, err)
	assert.Equal(t, 11.0, sum)
}
