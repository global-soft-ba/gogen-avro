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

type Event struct {
	// Unique ID for this event.
	Id string `json:"id"`
}

const EventAvroCRC64Fingerprint = "\xe8\x8cH,\xc2W\x8a\xe0"

func NewEvent() *Event {
	return &Event{}
}

func DeserializeEvent(r io.Reader) (*Event, error) {
	t := NewEvent()
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

func DeserializeEventFromSchema(r io.Reader, schema string) (*Event, error) {
	t := NewEvent()

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

func writeEvent(r *Event, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Event) Serialize(w io.Writer) error {
	return writeEvent(r, w)
}

func (r *Event) Schema() string {
	return "{\"fields\":[{\"doc\":\"Unique ID for this event.\",\"name\":\"id\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}}],\"name\":\"event\",\"subject\":\"event\",\"type\":\"record\"}"
}

func (r *Event) SchemaName() string {
	return "event"
}

func (_ *Event) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Event) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Event) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Event) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Event) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Event) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Event) SetString(v string)   { panic("Unsupported operation") }
func (_ *Event) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Event) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Id}
	}
	panic("Unknown field index")
}

func (r *Event) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Event) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Event) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Event) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Event) Finalize()                        {}

func (_ *Event) AvroCRC64Fingerprint() []byte {
	return []byte(EventAvroCRC64Fingerprint)
}

func (r *Event) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Event) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	if val, ok := fields["id"]; ok {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	return nil
}
