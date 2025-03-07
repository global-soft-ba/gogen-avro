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

func writeArrayUnionRecordFooRecordBar(r []UnionRecordFooRecordBar, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUnionRecordFooRecordBar(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUnionRecordFooRecordBarWrapper struct {
	Target *[]UnionRecordFooRecordBar
}

func (_ ArrayUnionRecordFooRecordBarWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayUnionRecordFooRecordBarWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayUnionRecordFooRecordBarWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayUnionRecordFooRecordBarWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayUnionRecordFooRecordBarWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayUnionRecordFooRecordBarWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayUnionRecordFooRecordBarWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayUnionRecordFooRecordBarWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayUnionRecordFooRecordBarWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayUnionRecordFooRecordBarWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayUnionRecordFooRecordBarWrapper) Finalize()        {}
func (_ ArrayUnionRecordFooRecordBarWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayUnionRecordFooRecordBarWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]UnionRecordFooRecordBar, 0, s)
	}
}
func (r ArrayUnionRecordFooRecordBarWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayUnionRecordFooRecordBarWrapper) AppendArray() types.Field {
	var v UnionRecordFooRecordBar
	v = NewUnionRecordFooRecordBar()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
