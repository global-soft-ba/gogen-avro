// Code generated by github.com/global-soft-ba/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
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

var _ = fmt.Printf

type StringRec struct {
	ProductName string `json:"productName"`
}

const StringRecAvroCRC64Fingerprint = "w\x836\xab\x9d\xe9\x00\x15"

func NewStringRec() StringRec {
	r := StringRec{}
	return r
}

func DeserializeStringRec(r io.Reader) (StringRec, error) {
	t := NewStringRec()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeStringRecFromSchema(r io.Reader, schema string) (StringRec, error) {
	t := NewStringRec()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeStringRec(r StringRec, w io.Writer) error {
	var err error
	err = vm.WriteString(r.ProductName, w)
	if err != nil {
		return err
	}
	return err
}

func (r StringRec) Serialize(w io.Writer) error {
	return writeStringRec(r, w)
}

func (r StringRec) Schema() string {
	return "{\"fields\":[{\"name\":\"productName\",\"type\":\"string\"}],\"name\":\"StringRec\",\"type\":\"record\"}"
}

func (r StringRec) SchemaName() string {
	return "StringRec"
}

func (_ StringRec) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ StringRec) SetInt(v int32)       { panic("Unsupported operation") }
func (_ StringRec) SetLong(v int64)      { panic("Unsupported operation") }
func (_ StringRec) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ StringRec) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ StringRec) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ StringRec) SetString(v string)   { panic("Unsupported operation") }
func (_ StringRec) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *StringRec) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.ProductName}

		return w

	}
	panic("Unknown field index")
}

func (r *StringRec) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *StringRec) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ StringRec) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ StringRec) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ StringRec) HintSize(int)                     { panic("Unsupported operation") }
func (_ StringRec) Finalize()                        {}

func (_ StringRec) AvroCRC64Fingerprint() []byte {
	return []byte(StringRecAvroCRC64Fingerprint)
}

func (r StringRec) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["productName"], err = json.Marshal(r.ProductName)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *StringRec) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["productName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductName); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for productName")
	}
	return nil
}
