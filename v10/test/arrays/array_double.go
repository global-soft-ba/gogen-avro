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

func writeArrayDouble(r []float64, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteDouble(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayDoubleWrapper struct {
	Target *[]float64
}

func (_ ArrayDoubleWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayDoubleWrapper) Finalize()                        {}
func (_ ArrayDoubleWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayDoubleWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]float64, 0, s)
	}
}
func (r ArrayDoubleWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayDoubleWrapper) AppendArray() types.Field {
	var v float64

	*r.Target = append(*r.Target, v)
	return &types.Double{Target: &(*r.Target)[len(*r.Target)-1]}
}
