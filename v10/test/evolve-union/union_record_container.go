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

func NewUnionRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewUnionRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type UnionRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewUnionRecordReader(r io.Reader) (*UnionRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewUnionRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &UnionRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r UnionRecordReader) Read() (UnionRecord, error) {
	t := NewUnionRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
