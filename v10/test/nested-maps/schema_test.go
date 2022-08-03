package avro

import (
	"io"
	"testing"

	"github.com/global-soft-ba/gogen-avro/v10/container"
	"github.com/global-soft-ba/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTrip(t,
		func() container.AvroRecord { return &NestedMap{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeNestedMap(r)
			return &record, err
		})
}
