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

type RecordSchema struct {
	Data []UnionRecordFooRecordBar `json:"data"`
}

const RecordSchemaAvroCRC64Fingerprint = "\xb8\xce$\xde\v\xd5\x00<"

func NewRecordSchema() RecordSchema {
	r := RecordSchema{}
	r.Data = make([]UnionRecordFooRecordBar, 0)

	return r
}

func DeserializeRecordSchema(r io.Reader) (RecordSchema, error) {
	t := NewRecordSchema()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRecordSchemaFromSchema(r io.Reader, schema string) (RecordSchema, error) {
	t := NewRecordSchema()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRecordSchema(r RecordSchema, w io.Writer) error {
	var err error
	err = writeArrayUnionRecordFooRecordBar(r.Data, w)
	if err != nil {
		return err
	}
	return err
}

func (r RecordSchema) Serialize(w io.Writer) error {
	return writeRecordSchema(r, w)
}

func (r RecordSchema) Schema() string {
	return "{\"fields\":[{\"name\":\"data\",\"type\":{\"items\":[{\"fields\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"RecordFoo\",\"type\":\"record\"},{\"fields\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"RecordBar\",\"type\":\"record\"}],\"type\":\"array\"}}],\"name\":\"RecordSchema\",\"type\":\"record\"}"
}

func (r RecordSchema) SchemaName() string {
	return "RecordSchema"
}

func (_ RecordSchema) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RecordSchema) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RecordSchema) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RecordSchema) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RecordSchema) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RecordSchema) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RecordSchema) SetString(v string)   { panic("Unsupported operation") }
func (_ RecordSchema) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RecordSchema) Get(i int) types.Field {
	switch i {
	case 0:
		r.Data = make([]UnionRecordFooRecordBar, 0)

		w := ArrayUnionRecordFooRecordBarWrapper{Target: &r.Data}

		return w

	}
	panic("Unknown field index")
}

func (r *RecordSchema) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RecordSchema) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RecordSchema) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RecordSchema) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RecordSchema) HintSize(int)                     { panic("Unsupported operation") }
func (_ RecordSchema) Finalize()                        {}

func (_ RecordSchema) AvroCRC64Fingerprint() []byte {
	return []byte(RecordSchemaAvroCRC64Fingerprint)
}

func (r RecordSchema) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["data"], err = json.Marshal(r.Data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RecordSchema) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["data"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Data); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for data")
	}
	return nil
}
