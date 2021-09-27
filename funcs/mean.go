package funcs

import (
	"fmt"
	"github.com/ekimeel/timeseries/model"
	"math"
)

//Mean of a TimeSeries, NaN if empty
type Mean struct {
	Dimension string
}

func (function Mean) Eval(series *model.TimeSeries) (float64, error) {
	index := series.GetDimensionIndex(function.Dimension)
	if index < 0 {
		return -1, fmt.Errorf("dimension not found [%s]", function.Dimension)
	}

	if series.Size() == 0 {
		return math.NaN(), nil
	}

	sum, err := series.Eval(&Sum{Dimension: function.Dimension})
	if err != nil {
		return 0.0, err
	}

	return sum / float64(series.Size()), nil
}

