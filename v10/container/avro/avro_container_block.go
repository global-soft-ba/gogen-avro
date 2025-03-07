// Code generated by github.com/global-soft-ba/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     block.avsc
 *     header.avsc
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

type AvroContainerBlock struct {
	NumRecords int64 `json:"numRecords"`

	RecordBytes Bytes `json:"recordBytes"`

	Sync Sync `json:"sync"`
}

const AvroContainerBlockAvroCRC64Fingerprint = "\x0e\xecj@ٔ\xe14"

func NewAvroContainerBlock() AvroContainerBlock {
	r := AvroContainerBlock{}
	return r
}

func DeserializeAvroContainerBlock(r io.Reader) (AvroContainerBlock, error) {
	t := NewAvroContainerBlock()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAvroContainerBlockFromSchema(r io.Reader, schema string) (AvroContainerBlock, error) {
	t := NewAvroContainerBlock()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAvroContainerBlock(r AvroContainerBlock, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.NumRecords, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.RecordBytes, w)
	if err != nil {
		return err
	}
	err = writeSync(r.Sync, w)
	if err != nil {
		return err
	}
	return err
}

func (r AvroContainerBlock) Serialize(w io.Writer) error {
	return writeAvroContainerBlock(r, w)
}

func (r AvroContainerBlock) Schema() string {
	return "{\"fields\":[{\"name\":\"numRecords\",\"type\":\"long\"},{\"name\":\"recordBytes\",\"type\":\"bytes\"},{\"name\":\"sync\",\"type\":{\"name\":\"sync\",\"size\":16,\"type\":\"fixed\"}}],\"name\":\"AvroContainerBlock\",\"type\":\"record\"}"
}

func (r AvroContainerBlock) SchemaName() string {
	return "AvroContainerBlock"
}

func (_ AvroContainerBlock) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AvroContainerBlock) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AvroContainerBlock) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AvroContainerBlock) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AvroContainerBlock) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AvroContainerBlock) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AvroContainerBlock) SetString(v string)   { panic("Unsupported operation") }
func (_ AvroContainerBlock) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AvroContainerBlock) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.NumRecords}

		return w

	case 1:
		w := BytesWrapper{Target: &r.RecordBytes}

		return w

	case 2:
		w := SyncWrapper{Target: &r.Sync}

		return w

	}
	panic("Unknown field index")
}

func (r *AvroContainerBlock) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AvroContainerBlock) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AvroContainerBlock) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AvroContainerBlock) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AvroContainerBlock) HintSize(int)                     { panic("Unsupported operation") }
func (_ AvroContainerBlock) Finalize()                        {}

func (_ AvroContainerBlock) AvroCRC64Fingerprint() []byte {
	return []byte(AvroContainerBlockAvroCRC64Fingerprint)
}

func (r AvroContainerBlock) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["numRecords"], err = json.Marshal(r.NumRecords)
	if err != nil {
		return nil, err
	}
	output["recordBytes"], err = json.Marshal(r.RecordBytes)
	if err != nil {
		return nil, err
	}
	output["sync"], err = json.Marshal(r.Sync)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AvroContainerBlock) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["numRecords"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumRecords); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numRecords")
	}
	val = func() json.RawMessage {
		if v, ok := fields["recordBytes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RecordBytes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for recordBytes")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sync"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sync); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sync")
	}
	return nil
}
