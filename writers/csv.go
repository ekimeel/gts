package writers

import (
	"encoding/csv"
	"fmt"
	"github.com/ekimeel/timeseries/model"
	"io"
	"os"
)
const (
	timeStampCol      = "timestamp"
)


/* csv */
type CsvWriter struct {
	//Path       	string
	TimeLayout	string
	Writer		io.Writer
}

//Writes the provided TimeSeries to a CSV formatted file at the provided path
func (csvWriter *CsvWriter) writeCsv(series *model.TimeSeries, w io.Writer) error {
	timeLayout := csvWriter.TimeLayout
	if len(timeLayout) == 0 {
		timeLayout = model.DefaultTimeLayout
	}

	/* header */
	writer := csv.NewWriter(w)
	defer writer.Flush()
	header := collectHeaders(series)

	err := writer.Write(header)
	if err != nil {
		return fmt.Errorf("failed to write header: %s", err)
	}

	for i := 0; i < series.Size(); i++ {
		values := collectValuesAt(series, i, timeLayout)

		err := writer.Write(values)
		if err != nil {
			return fmt.Errorf("error writing values at row [%d]: %s", i, err)
		}
	}

	return nil
}


//Writes the provided TimeSeries to a CSV formatted file at the provided path
func (csvWriter *CsvWriter) WriteToFile(series *model.TimeSeries, path string) error {
	err := ensureDir(path)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("cannot write to file caused by: %s", err)
	}
	defer file.Close()
	return csvWriter.writeCsv(series, file)
}



//Writes a new CSV file to the current writer's path
func (csvWriter *CsvWriter) Write(series *model.TimeSeries) error {
	return csvWriter.writeCsv(series, csvWriter.Writer)
}

