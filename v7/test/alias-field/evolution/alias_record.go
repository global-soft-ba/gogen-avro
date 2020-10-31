// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
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

type AliasRecord struct {
	B string `json:"b"`

	D string `json:"d"`
}

const AliasRecordAvroCRC64Fingerprint = "wT\xf0\x8a+ɨ\xce"

func NewAliasRecord() *AliasRecord {
	return &AliasRecord{}
}

func DeserializeAliasRecord(r io.Reader) (*AliasRecord, error) {
	t := NewAliasRecord()
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

func DeserializeAliasRecordFromSchema(r io.Reader, schema string) (*AliasRecord, error) {
	t := NewAliasRecord()

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

func writeAliasRecord(r *AliasRecord, w io.Writer) error {
	var err error
	err = vm.WriteString(r.B, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.D, w)
	if err != nil {
		return err
	}
	return err
}

func (r *AliasRecord) Serialize(w io.Writer) error {
	return writeAliasRecord(r, w)
}

func (r *AliasRecord) Schema() string {
	return "{\"fields\":[{\"aliases\":[\"a\"],\"name\":\"b\",\"type\":\"string\"},{\"name\":\"d\",\"type\":\"string\"}],\"name\":\"AliasRecord\",\"type\":\"record\"}"
}

func (r *AliasRecord) SchemaName() string {
	return "AliasRecord"
}

func (_ *AliasRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *AliasRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *AliasRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *AliasRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *AliasRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *AliasRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *AliasRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *AliasRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AliasRecord) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.B}
	case 1:
		return &types.String{Target: &r.D}
	}
	panic("Unknown field index")
}

func (r *AliasRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AliasRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *AliasRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *AliasRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *AliasRecord) Finalize()                        {}

func (_ *AliasRecord) AvroCRC64Fingerprint() []byte {
	return []byte(AliasRecordAvroCRC64Fingerprint)
}

func (r *AliasRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["b"], err = json.Marshal(r.B)
	if err != nil {
		return nil, err
	}
	output["d"], err = json.Marshal(r.D)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AliasRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	if val, ok := fields["b"]; ok {
		if err := json.Unmarshal([]byte(val), &r.B); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for b")
	}
	if val, ok := fields["d"]; ok {
		if err := json.Unmarshal([]byte(val), &r.D); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for d")
	}
	return nil
}
