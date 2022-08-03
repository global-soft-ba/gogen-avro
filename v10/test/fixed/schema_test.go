package avro

import (
	"io"
	"testing"

	"github.com/global-soft-ba/gogen-avro/v10/container"
	"github.com/global-soft-ba/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTripExactBytes(t,
		func() container.AvroRecord { return &FixedTestRecord{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeFixedTestRecord(r)
			return &record, err
		})
}
