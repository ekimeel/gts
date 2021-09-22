package transformations

import (
	"github.com/ekimeel/timeseries/model"
	"math"
)


// fits
type Scale struct {
	FitTo *model.TimeSeries
}

func (t *Scale) Transform(ts *model.TimeSeries) (model.TimeSeries, error) {

	c := model.TimeSeries{}
	for _, time := range ts.Times() {
		c.Add(time, []float64{})
	}

	for i, dimension := range ts.GetDimensions() {
		aD := ts.GetDimension(dimension)
		bD := t.FitTo.GetDimensionAt(i)
		cD := scale(&aD, &bD)

		c.AppendDimension(dimension, cD)

	}

	return c, nil
}

func scale(nums, model *[]float64) []float64  {

	aMin := math.MaxFloat64
	aMax := math.SmallestNonzeroFloat64
	for i := 0; i < len(*nums); i++ {
		aMin = math.Min(aMin, (*nums)[i])
		aMax = math.Max(aMax, (*nums)[i])
	}

	bMin := math.MaxFloat64
	bMax := math.SmallestNonzeroFloat64
	for i := 0; i < len(*model); i++ {
		bMin = math.Min(bMin, (*model)[i])
		bMax = math.Max(bMax, (*model)[i])
	}

	scaled := make([]float64, len(*nums))

	if aMin == aMax && aMin != 0 {
		// in the case that A is nums non-zero line, scale the line to the mid-point of B
		bMid := (bMin + bMax) / 2
		for i, _ := range *nums {
			scaled[i] = bMid
		}
	} else {
		factor := bMax / aMax

		for i, num := range *nums {
			if num == aMin {
				scaled[i] = bMin
			} else {
				scaled[i] = num * factor
			}
		}
	}

	return scaled
}
