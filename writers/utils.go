package writers

import (
	"fmt"
	"github.com/ekimeel/timeseries/model"
	"os"
	"path/filepath"
	"time"
)

func collectHeaders(series *model.TimeSeries) []string {
	headers := make([]string, series.GetDimensionCount()+1)
	headers[0] = timeStampCol
	for i, dimension := range series.GetDimensions() {
		headers[i+1] = dimension
	}
	return headers
}

func collectValuesAt(series *model.TimeSeries, i int, timeLayout string) []string {
	epoch := series.GetTimeAtNthPoint(i)
	measures := series.GetMeasurementVector(i)
	values := make([]string, len(measures)+1)
	values[0] = time.Unix(epoch, 0).Format(timeLayout)
	for j := 0; j < len(measures); j++ {
		values[j+1] = fmt.Sprintf("%f", measures[j])
	}

	return values
}

//ensures the provided path's directory structure exists if not, it creates it
func ensureDir(path string) error {
	dirName := filepath.Dir(path)
	if _, serr := os.Stat(dirName); serr != nil {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

