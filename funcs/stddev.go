package funcs

import (
	"fmt"
	"github.com/ekimeel/timeseries/model"
	"math"
)

type StdDev struct {
	Dimension string
}

func (function StdDev) EvalTimeSeries(series *model.TimeSeries) (float64, error) {

	dim := series.GetDimensionIndex(function.Dimension)
	if dim < 0 {
		return math.NaN(), fmt.Errorf("dimension not found [%s]", function.Dimension)
	}

	mean, err := series.Eval(Mean{Dimension: function.Dimension})
	if err != nil {
		return math.NaN(), err
	}

	sd := 0.0
	for i := 0; i < series.Size(); i++ {
		sd += math.Pow(*series.At(i, dim) - mean, 2)
	}

	variance := sd / float64(series.Size())
	return math.Sqrt(variance), nil

}

func (function StdDev) Eval(values []float64) (float64, error) {

	mean, err := Mean{}.Eval(values)
	if err != nil {
		return math.NaN(), err
	}
	sd := 0.0
	for i := 0; i < len(values); i++ {
		sd += math.Pow(values[i] - mean, 2)
	}

	variance := sd / float64(len(values))
	return math.Sqrt(variance), nil
}
