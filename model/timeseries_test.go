package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMultivariateTimeSeries_Add(t *testing.T) {

	timeseries := TimeSeries{}

	start := time.Now()
	timeseries.SetDimensions([]string{"one", "two", "three"})
	timeseries.Add(start.Add(1*time.Minute).Unix(), []float64{1.0, 2.0, 3.0})
	timeseries.Add(start.Add(2*time.Minute).Unix(), []float64{1.1, 2.1, 3.1})
	timeseries.Add(start.Add(3*time.Minute).Unix(), []float64{1.2, 2.2, 3.2})
	timeseries.Add(start.Add(4*time.Minute).Unix(), []float64{1.3, 2.3, 3.3})
	timeseries.Add(start.Add(5*time.Minute).Unix(), []float64{1.4, 2.4, 3.4})

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
	ts.Add(start.Add(1*time.Minute).Unix(), []float64{1.0, 2.0, 3.0})
	ts.Add(start.Add(2*time.Minute).Unix(), []float64{1.1, 2.1, 3.1})
	ts.Add(start.Add(3*time.Minute).Unix(), []float64{1.2, 2.2, 3.2})
	ts.Add(start.Add(4*time.Minute).Unix(), []float64{1.3, 2.3, 3.3})
	ts.Add(start.Add(5*time.Minute).Unix(), []float64{1.4, 2.4, 3.4})

	i, err := ts.AppendDimension("four", []float64{4.0, 4.1, 4.2, 4.4, 4.3})
	assert.Nil(t, err)
	assert.Equal(t, 3, i)
	assert.Equal(t, 4.0, ts.GetDimension("four")[0])

}

func TestTimeSeries_Add_With1MillionPoints(t *testing.T) {

	ts := TimeSeries{}

	clock := time.Now()
	ts.SetDimensions([]string{"v0"})
	for i := 0; i < 1000000; i++ {
		clock = clock.Add(1*time.Second)
		ts.Add(clock.Unix(), []float64{float64(i)})
	}

	assert.Equal(t, 1000000, ts.Size())
}

func TestTimeSeries_Filter(t *testing.T) {
	ts := TimeSeries{}

	start := time.Now()
	ts.SetDimensions([]string{"one", "two", "three"})
	ts.Add(start.Add(1*time.Minute).Unix(), []float64{1.0, 2.0, 3.0})
	ts.Add(start.Add(2*time.Minute).Unix(), []float64{1.1, 2.1, 3.1})
	ts.Add(start.Add(3*time.Minute).Unix(), []float64{1.2, 2.2, 3.2})
	ts.Add(start.Add(4*time.Minute).Unix(), []float64{1.3, 2.3, 3.3})
	ts.Add(start.Add(5*time.Minute).Unix(), []float64{1.4, 2.4, 3.4})

	v3d1 := ts.At(3, 1)

	filtered := ts.Filter(func(time int64, values []float64) bool {
		return values[1] > 2.2
	})


	assert.Equal(t, 2, filtered.Size())
	assert.Equal(t, 2.3, *filtered.At(0,1))
	*v3d1 = 100
	assert.Equal(t, 100.0, *filtered.At(0,1))
}


func TestTimeSeries_LastTime(t *testing.T) {
	reader := CsvReader{Path: "../testdata/3x1000.csv"}
	ts, err := reader.Read()
	assert.Nil(t, err)

	latest := ts.LastTime()
	assert.Equal(t, int64(1609537200), latest.Unix())
}

func TestTimeSeries_Last(t *testing.T) {
	reader := CsvReader{Path: "../testdata/3x1000.csv"}
	ts, err := reader.Read()
	assert.Nil(t, err)

	values := ts.Last()
	assert.Equal(t, 3, len(*values))
	assert.Equal(t, 1.999000, (*values)[0])
	assert.Equal(t, 2.999000, (*values)[1])
	assert.Equal(t, 3.999000, (*values)[2])
}

func TestTimeSeries_FirstTime(t *testing.T) {
	reader := CsvReader{Path: "../testdata/3x1000.csv"}
	ts, err := reader.Read()
	assert.Nil(t, err)

	latest := ts.FirstTime()
	assert.Equal(t, int64(1609477260), latest.Unix())
}

func TestTimeSeries_First(t *testing.T) {
	reader := CsvReader{Path: "../testdata/3x1000.csv"}
	ts, err := reader.Read()
	assert.Nil(t, err)

	values := ts.First()
	assert.Equal(t, 3, len(*values))
	assert.Equal(t, 1.000000, (*values)[0])
	assert.Equal(t, 2.000000, (*values)[1])
	assert.Equal(t, 3.000000, (*values)[2])
}