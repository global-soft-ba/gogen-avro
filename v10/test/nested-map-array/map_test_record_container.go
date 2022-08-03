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

func NewMapTestRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewMapTestRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type MapTestRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewMapTestRecordReader(r io.Reader) (*MapTestRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewMapTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &MapTestRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r MapTestRecordReader) Read() (MapTestRecord, error) {
	t := NewMapTestRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
