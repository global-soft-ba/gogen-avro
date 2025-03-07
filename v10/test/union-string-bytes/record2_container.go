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

func NewRecord2Writer(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewRecord2()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type Record2Reader struct {
	r io.Reader
	p *vm.Program
}

func NewRecord2Reader(r io.Reader) (*Record2Reader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewRecord2()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &Record2Reader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r Record2Reader) Read() (Record2, error) {
	t := NewRecord2()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
