package timeseries

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPAA(t *testing.T) {

	ts := TimeSeries{}

	start := time.Now()
	ts.SetDimensions([]string{"one"})

	ts.Add(start.Add(1*time.Minute).Unix(), []float64{1.0})
	ts.Add(start.Add(2*time.Minute).Unix(), []float64{2.0})
	ts.Add(start.Add(3*time.Minute).Unix(), []float64{3.0})
	ts.Add(start.Add(4*time.Minute).Unix(), []float64{4.0})

	nts, err := ts.Transform(PAA{ShrinkSize: 2})

	assert.Nil(t, err)
	assert.Equal(t, 2, nts.Size())
	assert.Equal(t, 4, ts.Size())
	assert.Equal(t, 1.5, nts.GetDimensionAt(0)[0])
	assert.Equal(t, 3.5, nts.GetDimensionAt(0)[1])

}
