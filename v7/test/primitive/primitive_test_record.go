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

type PrimitiveTestRecord struct {
	IntField int32 `json:"IntField"`

	LongField int64 `json:"LongField"`

	FloatField float32 `json:"FloatField"`

	DoubleField float64 `json:"DoubleField"`

	StringField string `json:"StringField"`

	BoolField bool `json:"BoolField"`

	BytesField Bytes `json:"BytesField"`
}

const PrimitiveTestRecordAvroCRC64Fingerprint = "۲\x16\xe9\xfet@\x13"

func NewPrimitiveTestRecord() *PrimitiveTestRecord {
	return &PrimitiveTestRecord{}
}

func DeserializePrimitiveTestRecord(r io.Reader) (*PrimitiveTestRecord, error) {
	t := NewPrimitiveTestRecord()
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

func DeserializePrimitiveTestRecordFromSchema(r io.Reader, schema string) (*PrimitiveTestRecord, error) {
	t := NewPrimitiveTestRecord()

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

func writePrimitiveTestRecord(r *PrimitiveTestRecord, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IntField, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.LongField, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.FloatField, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.DoubleField, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.StringField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.BoolField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.BytesField, w)
	if err != nil {
		return err
	}
	return err
}

func (r *PrimitiveTestRecord) Serialize(w io.Writer) error {
	return writePrimitiveTestRecord(r, w)
}

func (r *PrimitiveTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"LongField\",\"type\":\"long\"},{\"name\":\"FloatField\",\"type\":\"float\"},{\"name\":\"DoubleField\",\"type\":\"double\"},{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"PrimitiveTestRecord\",\"type\":\"record\"}"
}

func (r *PrimitiveTestRecord) SchemaName() string {
	return "PrimitiveTestRecord"
}

func (_ *PrimitiveTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PrimitiveTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Int{Target: &r.IntField}
	case 1:
		return &types.Long{Target: &r.LongField}
	case 2:
		return &types.Float{Target: &r.FloatField}
	case 3:
		return &types.Double{Target: &r.DoubleField}
	case 4:
		return &types.String{Target: &r.StringField}
	case 5:
		return &types.Boolean{Target: &r.BoolField}
	case 6:
		return &BytesWrapper{Target: &r.BytesField}
	}
	panic("Unknown field index")
}

func (r *PrimitiveTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PrimitiveTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *PrimitiveTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) Finalize()                        {}

func (_ *PrimitiveTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(PrimitiveTestRecordAvroCRC64Fingerprint)
}

func (r *PrimitiveTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IntField"], err = json.Marshal(r.IntField)
	if err != nil {
		return nil, err
	}
	output["LongField"], err = json.Marshal(r.LongField)
	if err != nil {
		return nil, err
	}
	output["FloatField"], err = json.Marshal(r.FloatField)
	if err != nil {
		return nil, err
	}
	output["DoubleField"], err = json.Marshal(r.DoubleField)
	if err != nil {
		return nil, err
	}
	output["StringField"], err = json.Marshal(r.StringField)
	if err != nil {
		return nil, err
	}
	output["BoolField"], err = json.Marshal(r.BoolField)
	if err != nil {
		return nil, err
	}
	output["BytesField"], err = json.Marshal(r.BytesField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PrimitiveTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	if val, ok := fields["IntField"]; ok {
		if err := json.Unmarshal([]byte(val), &r.IntField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IntField")
	}
	if val, ok := fields["LongField"]; ok {
		if err := json.Unmarshal([]byte(val), &r.LongField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LongField")
	}
	if val, ok := fields["FloatField"]; ok {
		if err := json.Unmarshal([]byte(val), &r.FloatField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FloatField")
	}
	if val, ok := fields["DoubleField"]; ok {
		if err := json.Unmarshal([]byte(val), &r.DoubleField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DoubleField")
	}
	if val, ok := fields["StringField"]; ok {
		if err := json.Unmarshal([]byte(val), &r.StringField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StringField")
	}
	if val, ok := fields["BoolField"]; ok {
		if err := json.Unmarshal([]byte(val), &r.BoolField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BoolField")
	}
	if val, ok := fields["BytesField"]; ok {
		if err := json.Unmarshal([]byte(val), &r.BytesField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BytesField")
	}
	return nil
}
