package transformations

import (
	"github.com/ekimeel/timeseries/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRollUp_Transform(t *testing.T) {

	reader := model.CsvReader{Path: "../testdata/1x1000.csv"}
	ts, err := reader.Read()


	//ts.RollUp(15*time.Minute, &funcs.Mean{Dimension: "v0"})


	nts, err := ts.Transform(RollUp{ShrinkSize: 2})

	assert.Nil(t, err)
	assert.Equal(t, 2, nts.Size())
	assert.Equal(t, 4, ts.Size())
	assert.Equal(t, 1.5, nts.GetDimensionAt(0)[0])
	assert.Equal(t, 3.5, nts.GetDimensionAt(0)[1])

}
