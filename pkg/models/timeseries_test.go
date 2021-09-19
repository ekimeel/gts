package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMultivariateTimeSeries_Add(t *testing.T) {

	timeseries := TimeSeries{}

	start := time.Now()
	timeseries.SetDimensions([]string{"one", "two", "three"})

	timeseries.Add(start.Add(1 * time.Minute).Unix(), []float64{1.0, 2.0, 3.0})
	timeseries.Add(start.Add(2 * time.Minute).Unix(), []float64{1.1, 2.1, 3.1})
	timeseries.Add(start.Add(3 * time.Minute).Unix(), []float64{1.2, 2.2, 3.2})
	timeseries.Add(start.Add(4 * time.Minute).Unix(), []float64{1.3, 2.3, 3.3})
	timeseries.Add(start.Add(5 * time.Minute).Unix(), []float64{1.4, 2.4, 3.4})


	assert.Equal(t, 3, timeseries.CountOfDimensions())
	assert.Equal(t, 5, timeseries.Size())

	one := timeseries.GetDimension("one")
	assert.Equal(t, 5, len(one))
	assert.Equal(t, 1.0, one[0])
	assert.Equal(t, 1.4, one[4])

	two := timeseries.GetDimension("two")
	assert.Equal(t, 2.0, two[0])
	assert.Equal(t, 2.4, two[4])

	three := timeseries.GetDimension("three")
	assert.Equal(t, 5, len(three))
	assert.Equal(t, 3.0, three[0])
	assert.Equal(t, 3.4, three[4])
}

func TestTimeSeries_AppendDimension(t *testing.T) {
	ts := TimeSeries{}

	start := time.Now()
	ts.SetDimensions([]string{"one", "two", "three"})
	ts.Add(start.Add(1 * time.Minute).Unix(), []float64{1.0, 2.0, 3.0})
	ts.Add(start.Add(2 * time.Minute).Unix(), []float64{1.1, 2.1, 3.1})
	ts.Add(start.Add(3 * time.Minute).Unix(), []float64{1.2, 2.2, 3.2})
	ts.Add(start.Add(4 * time.Minute).Unix(), []float64{1.3, 2.3, 3.3})
	ts.Add(start.Add(5 * time.Minute).Unix(), []float64{1.4, 2.4, 3.4})

	i, err := ts.AppendDimension("four", []float64{4.0, 4.1, 4.2, 4.4, 4.3})
	assert.Nil(t, err)
	assert.Equal(t, 3, i)
	assert.Equal(t, 4.0, ts.GetDimension("four")[0])

}