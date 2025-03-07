// Code generated by github.com/global-soft-ba/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/global-soft-ba/gogen-avro/v10/compiler"
	"github.com/global-soft-ba/gogen-avro/v10/vm"
	"github.com/global-soft-ba/gogen-avro/v10/vm/types"
)

type UnionStringIntTypeEnum int

const (
	UnionStringIntTypeEnumString UnionStringIntTypeEnum = 0

	UnionStringIntTypeEnumInt UnionStringIntTypeEnum = 1
)

type UnionStringInt struct {
	String    string
	Int       int32
	UnionType UnionStringIntTypeEnum
}

func writeUnionStringInt(r UnionStringInt, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionStringIntTypeEnumString:
		return vm.WriteString(r.String, w)
	case UnionStringIntTypeEnumInt:
		return vm.WriteInt(r.Int, w)
	}
	return fmt.Errorf("invalid value for UnionStringInt")
}

func NewUnionStringInt() UnionStringInt {
	return UnionStringInt{}
}

func (r UnionStringInt) Serialize(w io.Writer) error {
	return writeUnionStringInt(r, w)
}

func DeserializeUnionStringInt(r io.Reader) (UnionStringInt, error) {
	t := NewUnionStringInt()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionStringIntFromSchema(r io.Reader, schema string) (UnionStringInt, error) {
	t := NewUnionStringInt()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r UnionStringInt) Schema() string {
	return "[\"string\",\"int\"]"
}

func (_ UnionStringInt) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ UnionStringInt) SetInt(v int32)      { panic("Unsupported operation") }
func (_ UnionStringInt) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ UnionStringInt) SetDouble(v float64) { panic("Unsupported operation") }
func (_ UnionStringInt) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ UnionStringInt) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionStringInt) SetLong(v int64) {

	r.UnionType = (UnionStringIntTypeEnum)(v)
}

func (r *UnionStringInt) Get(i int) types.Field {

	switch i {
	case 0:
		return &types.String{Target: (&r.String)}
	case 1:
		return &types.Int{Target: (&r.Int)}
	}
	panic("Unknown field index")
}
func (_ UnionStringInt) NullField(i int)                  { panic("Unsupported operation") }
func (_ UnionStringInt) HintSize(i int)                   { panic("Unsupported operation") }
func (_ UnionStringInt) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ UnionStringInt) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ UnionStringInt) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ UnionStringInt) Finalize()                        {}

func (r UnionStringInt) MarshalJSON() ([]byte, error) {

	switch r.UnionType {
	case UnionStringIntTypeEnumString:
		return json.Marshal(map[string]interface{}{"string": r.String})
	case UnionStringIntTypeEnumInt:
		return json.Marshal(map[string]interface{}{"int": r.Int})
	}
	return nil, fmt.Errorf("invalid value for UnionStringInt")
}

func (r *UnionStringInt) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["string"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.String)
	}
	if value, ok := fields["int"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Int)
	}
	return fmt.Errorf("invalid value for UnionStringInt")
}
