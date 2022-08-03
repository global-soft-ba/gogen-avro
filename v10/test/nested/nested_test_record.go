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

type NestedTestRecord struct {
	NumberField NumberRecord `json:"NumberField"`

	OtherField NestedRecord `json:"OtherField"`
}

const NestedTestRecordAvroCRC64Fingerprint = "b{m\\D\xbe\xaa\x96"

func NewNestedTestRecord() NestedTestRecord {
	r := NestedTestRecord{}
	r.NumberField = NewNumberRecord()

	r.OtherField = NewNestedRecord()

	return r
}

func DeserializeNestedTestRecord(r io.Reader) (NestedTestRecord, error) {
	t := NewNestedTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNestedTestRecordFromSchema(r io.Reader, schema string) (NestedTestRecord, error) {
	t := NewNestedTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNestedTestRecord(r NestedTestRecord, w io.Writer) error {
	var err error
	err = writeNumberRecord(r.NumberField, w)
	if err != nil {
		return err
	}
	err = writeNestedRecord(r.OtherField, w)
	if err != nil {
		return err
	}
	return err
}

func (r NestedTestRecord) Serialize(w io.Writer) error {
	return writeNestedTestRecord(r, w)
}

func (r NestedTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"NumberField\",\"type\":{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"LongField\",\"type\":\"long\"},{\"name\":\"FloatField\",\"type\":\"float\"},{\"name\":\"DoubleField\",\"type\":\"double\"}],\"name\":\"NumberRecord\",\"type\":\"record\"}},{\"name\":\"OtherField\",\"type\":{\"fields\":[{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"NestedRecord\",\"type\":\"record\"}}],\"name\":\"NestedTestRecord\",\"type\":\"record\"}"
}

func (r NestedTestRecord) SchemaName() string {
	return "NestedTestRecord"
}

func (_ NestedTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NestedTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NestedTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NestedTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NestedTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NestedTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NestedTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ NestedTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NestedTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.NumberField = NewNumberRecord()

		w := types.Record{Target: &r.NumberField}

		return w

	case 1:
		r.OtherField = NewNestedRecord()

		w := types.Record{Target: &r.OtherField}

		return w

	}
	panic("Unknown field index")
}

func (r *NestedTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NestedTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NestedTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NestedTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NestedTestRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ NestedTestRecord) Finalize()                        {}

func (_ NestedTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(NestedTestRecordAvroCRC64Fingerprint)
}

func (r NestedTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["NumberField"], err = json.Marshal(r.NumberField)
	if err != nil {
		return nil, err
	}
	output["OtherField"], err = json.Marshal(r.OtherField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NestedTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["NumberField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumberField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumberField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OtherField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OtherField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OtherField")
	}
	return nil
}
