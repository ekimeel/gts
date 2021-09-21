package funcs

import (
	"fmt"
	"github.com/ekimeel/timeseries/model"
)

type Sum struct {
	Dimension string
}

func (function Sum) Compute(series *model.TimeSeries) (float64, error) {
	index := series.GetDimensionIndex(function.Dimension)
	if index < 0 {
		return -1, fmt.Errorf("dimension not found [%s]", function.Dimension)
	}

	sum := 0.0
	for i := 0; i < series.Size(); i++ {
		sum += *series.At(i, index)
	}

	return sum, nil
}

