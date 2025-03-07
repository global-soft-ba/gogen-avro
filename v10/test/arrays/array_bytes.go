// Code generated by github.com/global-soft-ba/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/global-soft-ba/gogen-avro/v10/vm"
	"github.com/global-soft-ba/gogen-avro/v10/vm/types"
)

func writeArrayBytes(r []Bytes, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteBytes(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayBytesWrapper struct {
	Target *[]Bytes
}

func (_ ArrayBytesWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayBytesWrapper) Finalize()                        {}
func (_ ArrayBytesWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayBytesWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Bytes, 0, s)
	}
}
func (r ArrayBytesWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayBytesWrapper) AppendArray() types.Field {
	var v Bytes

	*r.Target = append(*r.Target, v)
	return &BytesWrapper{Target: &(*r.Target)[len(*r.Target)-1]}
}
