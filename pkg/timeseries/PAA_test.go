package timeseries

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"timeseries/pkg/models"
)

func TestNewPAA(t *testing.T) {

	ts := models.TimeSeries{}

	start := time.Now()
	ts.SetDimensions([]string{"one"})

	ts.Add(start.Add(1 * time.Minute).Unix(), []float64{1.0})
	ts.Add(start.Add(2 * time.Minute).Unix(), []float64{2.0})
	ts.Add(start.Add(3 * time.Minute).Unix(), []float64{3.0})
	ts.Add(start.Add(4 * time.Minute).Unix(), []float64{4.0})

	paa, err := NewPAA(&ts, 2)

	assert.Nil(t, err)
	assert.Equal(t, 2, paa.Size())
	assert.Equal(t, 4, paa.GetOriginalSize())
	assert.Equal(t, 2, paa.GetAggregatePtSize(1))

}
