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

var _ = fmt.Printf

type Child struct {
	Name string `json:"name"`
}

const ChildAvroCRC64Fingerprint = "\xad$C\xb1bU\xc0\x12"

func NewChild() *Child {
	return &Child{}
}

func DeserializeChild(r io.Reader) (*Child, error) {
	t := NewChild()
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

func DeserializeChildFromSchema(r io.Reader, schema string) (*Child, error) {
	t := NewChild()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeChild(r *Child, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Child) Serialize(w io.Writer) error {
	return writeChild(r, w)
}

func (r *Child) Schema() string {
	return "{\"fields\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Child\",\"type\":\"record\"}"
}

func (r *Child) SchemaName() string {
	return "Child"
}

func (_ *Child) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Child) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Child) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Child) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Child) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Child) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Child) SetString(v string)   { panic("Unsupported operation") }
func (_ *Child) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Child) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Name}
	}
	panic("Unknown field index")
}

func (r *Child) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Child) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Child) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Child) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Child) Finalize()                        {}

func (_ *Child) AvroCRC64Fingerprint() []byte {
	return []byte(ChildAvroCRC64Fingerprint)
}

func (r *Child) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Child) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	if val, ok := fields["name"]; ok {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	return nil
}
