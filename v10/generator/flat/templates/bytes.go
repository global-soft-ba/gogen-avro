package templates

const BytesTemplate = `
import (
	"encoding/json"

	"github.com/global-soft-ba/gogen-avro/v10/vm/types"
	"github.com/global-soft-ba/gogen-avro/v10/util"
)

type Bytes []byte

func (b *Bytes) UnmarshalJSON(data []byte) (error) {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*b = util.DecodeByteString(s)
	return nil
}

func (b Bytes) MarshalJSON() ([]byte, error) {
	return []byte(util.EncodeByteString(b)), nil
}

type BytesWrapper struct {
	Target *Bytes
}

func (b BytesWrapper) SetBoolean(v bool) {
	panic("Unable to assign bytes to bytes field")
}

func (b BytesWrapper) SetInt(v int32) {
	panic("Unable to assign int to bytes field")
}

func (b BytesWrapper) SetLong(v int64) {
	panic("Unable to assign long to bytes field")
}

func (b BytesWrapper) SetFloat(v float32) {
	panic("Unable to assign float to bytes field")
}

func (b BytesWrapper) SetDouble(v float64) {
	panic("Unable to assign double to bytes field")
}

func (b BytesWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to bytes field")
}

func (b BytesWrapper) SetBytes(v []byte) {
	*(b.Target) = v
}

func (b BytesWrapper) SetString(v string) {
	*(b.Target) = []byte(v)
}

func (b BytesWrapper) Get(i int) types.Field {
	panic("Unable to get field from bytes field")
}

func (b BytesWrapper) SetDefault(i int) {
	panic("Unable to set default on bytes field")
}

func (b BytesWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from bytes field")
}

func (b BytesWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from bytes field")
}

func (b BytesWrapper) NullField(int) {
	panic("Unable to null field in bytes field")
}

func (b BytesWrapper) HintSize(int) {
	panic("Unable to hint size in bytes field")
}

func (b BytesWrapper) Finalize() {}

`
