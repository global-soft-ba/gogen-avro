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

func NewNestedMapWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewNestedMap()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type NestedMapReader struct {
	r io.Reader
	p *vm.Program
}

func NewNestedMapReader(r io.Reader) (*NestedMapReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewNestedMap()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &NestedMapReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r NestedMapReader) Read() (NestedMap, error) {
	t := NewNestedMap()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
