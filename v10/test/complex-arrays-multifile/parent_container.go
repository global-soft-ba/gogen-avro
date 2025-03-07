// Code generated by github.com/global-soft-ba/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     parent.avsc
 *     child.avsc
 */
package avro

import (
	"io"

	"github.com/global-soft-ba/gogen-avro/v10/compiler"
	"github.com/global-soft-ba/gogen-avro/v10/container"
	"github.com/global-soft-ba/gogen-avro/v10/vm"
)

func NewParentWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewParent()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type ParentReader struct {
	r io.Reader
	p *vm.Program
}

func NewParentReader(r io.Reader) (*ParentReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewParent()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ParentReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ParentReader) Read() (Parent, error) {
	t := NewParent()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
