// Code generated by github.com/global-soft-ba/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/global-soft-ba/gogen-avro/v10/compiler"
	"github.com/global-soft-ba/gogen-avro/v10/container"
	"github.com/global-soft-ba/gogen-avro/v10/vm"
)

func NewFixedTestRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewFixedTestRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type FixedTestRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewFixedTestRecordReader(r io.Reader) (*FixedTestRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewFixedTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &FixedTestRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r FixedTestRecordReader) Read() (FixedTestRecord, error) {
	t := NewFixedTestRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
