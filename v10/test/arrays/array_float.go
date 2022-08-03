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

func writeArrayFloat(r []float32, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteFloat(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayFloatWrapper struct {
	Target *[]float32
}

func (_ ArrayFloatWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayFloatWrapper) Finalize()                        {}
func (_ ArrayFloatWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayFloatWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]float32, 0, s)
	}
}
func (r ArrayFloatWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayFloatWrapper) AppendArray() types.Field {
	var v float32

	*r.Target = append(*r.Target, v)
	return &types.Float{Target: &(*r.Target)[len(*r.Target)-1]}
}
