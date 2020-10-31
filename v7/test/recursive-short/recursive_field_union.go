// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

type RecursiveFieldUnionTypeEnum int

const (
	RecursiveFieldUnionTypeEnumRecursiveUnionTestRecord RecursiveFieldUnionTypeEnum = 1
)

type RecursiveFieldUnion struct {
	Null                     *types.NullVal
	RecursiveUnionTestRecord *RecursiveUnionTestRecord
	UnionType                RecursiveFieldUnionTypeEnum
}

func writeRecursiveFieldUnion(r *RecursiveFieldUnion, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case RecursiveFieldUnionTypeEnumRecursiveUnionTestRecord:
		return writeRecursiveUnionTestRecord(r.RecursiveUnionTestRecord, w)
	}
	return fmt.Errorf("invalid value for *RecursiveFieldUnion")
}

func NewRecursiveFieldUnion() *RecursiveFieldUnion {
	return &RecursiveFieldUnion{}
}

func (r *RecursiveFieldUnion) Serialize(w io.Writer) error {
	return writeRecursiveFieldUnion(r, w)
}

func DeserializeRecursiveFieldUnion(r io.Reader) (*RecursiveFieldUnion, error) {
	t := NewRecursiveFieldUnion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func (r *RecursiveFieldUnion) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"RecursiveField\",\"type\":[\"null\",\"RecursiveUnionTestRecord\"]}],\"name\":\"RecursiveUnionTestRecord\",\"type\":\"record\"}]"
}

func (_ *RecursiveFieldUnion) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *RecursiveFieldUnion) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *RecursiveFieldUnion) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *RecursiveFieldUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *RecursiveFieldUnion) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *RecursiveFieldUnion) SetString(v string)  { panic("Unsupported operation") }
func (r *RecursiveFieldUnion) SetLong(v int64) {
	r.UnionType = (RecursiveFieldUnionTypeEnum)(v)
}
func (r *RecursiveFieldUnion) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		r.RecursiveUnionTestRecord = NewRecursiveUnionTestRecord()
		return r.RecursiveUnionTestRecord
	}
	panic("Unknown field index")
}
func (_ *RecursiveFieldUnion) NullField(i int)                  { panic("Unsupported operation") }
func (_ *RecursiveFieldUnion) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *RecursiveFieldUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *RecursiveFieldUnion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *RecursiveFieldUnion) Finalize()                        {}

func (r *RecursiveFieldUnion) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case RecursiveFieldUnionTypeEnumRecursiveUnionTestRecord:
		return json.Marshal(map[string]interface{}{"RecursiveUnionTestRecord": r.RecursiveUnionTestRecord})
	}
	return nil, fmt.Errorf("invalid value for *RecursiveFieldUnion")
}

func (r *RecursiveFieldUnion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["RecursiveUnionTestRecord"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.RecursiveUnionTestRecord)
	}
	return fmt.Errorf("invalid value for *RecursiveFieldUnion")
}
