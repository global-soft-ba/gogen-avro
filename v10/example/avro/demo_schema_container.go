// Code generated by github.com/global-soft-ba/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     example.avsc
 */
package avro

import (
	"io"

	"github.com/global-soft-ba/gogen-avro/v10/compiler"
	"github.com/global-soft-ba/gogen-avro/v10/container"
	"github.com/global-soft-ba/gogen-avro/v10/vm"
)

func NewDemoSchemaWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewDemoSchema()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type DemoSchemaReader struct {
	r io.Reader
	p *vm.Program
}

func NewDemoSchemaReader(r io.Reader) (*DemoSchemaReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewDemoSchema()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &DemoSchemaReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r DemoSchemaReader) Read() (DemoSchema, error) {
	t := NewDemoSchema()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
