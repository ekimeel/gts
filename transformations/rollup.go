package transformations

import (
	"errors"
	"github.com/ekimeel/timeseries/model"
	"time"
)

//Piecewise Aggregate Approximation (RollUp)
type RollUp struct {
	Duration time.Duration
	Agg model.ValueFunction
}

func (t *RollUp) Transform(ts *model.TimeSeries) (model.TimeSeries, error) {

	start := ts.FirstTime().Round(t.Duration)
	end := ts.LastTime().Round(t.Duration)
	dur := end.UnixMilli() - start.UnixMilli()
	intervalMs := t.Duration.Milliseconds()
	size := int(dur / intervalMs)

	var rollup model.TimeSeries
	rollup.SetDimensions(ts.GetDimensions())


	if size >= ts.Size() {
		return rollup, errors.New("cannot preform a Piecewise Aggregate Approximation (RollUp) to a size larger than " +
			"the source time series")
	}

	var from = start
	for ok := true; ok; ok = from.Before(end) {
		to := from.Add(t.Duration)
		chunk, err := ts.Filter(func(time int64, values []float64) bool {
			return time >= from.Unix() && time < to.Unix()
		})

		if err != nil {
			return rollup, err
		}

		row := make([]float64, ts.GetDimensionCount())
		for i, d := range ts.GetDimensions() {
			values := chunk.GetDimension(d)
			row[i], _ = t.Agg.Eval(values)
		}
		rollup.Add(from.Unix(), row)

		from = from.Add(t.Duration)
	}

	return rollup, nil



}
