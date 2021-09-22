package transformations

import (
	"github.com/ekimeel/timeseries/model"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestBoundaryScale_Scale_1(t *testing.T) {

	a := model.TimeSeries{}
	a.SetDimensions([]string{"v0"})
	a.Add(1, []float64{1.0})
	a.Add(2, []float64{2.0})
	a.Add(3, []float64{1.0})
	a.Add(4, []float64{4.0})
	a.Add(5, []float64{1.0})
	a.Add(6, []float64{8.0})
	a.Add(7, []float64{1.0})
	a.Add(8, []float64{16.0})
	a.Add(9, []float64{1.0})
	a.Add(10, []float64{32.0})

	/**/
	b := model.TimeSeries{}
	b.SetDimensions([]string{"v0"})
	b.Add(1, []float64{10.0})
	b.Add(2, []float64{20.0})
	b.Add(3, []float64{10.0})
	b.Add(4, []float64{40.0})
	b.Add(5, []float64{10.0})
	b.Add(6, []float64{80.0})
	b.Add(7, []float64{10.0})
	b.Add(8, []float64{160.0})
	b.Add(9, []float64{10.0})
	b.Add(10, []float64{320.0})

	c, err := a.Transform(&Scale{FitTo: &b})

	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, a.Size(), c.Size())
	assert.Equal(t, 10.0, *c.At(0, 0))
	assert.Equal(t, 320.0, *c.At(c.Size()-1, 0))
}
