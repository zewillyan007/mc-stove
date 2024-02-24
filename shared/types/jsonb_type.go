package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JsonbType map[string]interface{}

func (o JsonbType) Value() (driver.Value, error) {
	bytes, err := json.Marshal(o)
	return string(bytes), err
}

func (o *JsonbType) Scan(src interface{}) error {

	var source []byte
	mapt := make(map[string]interface{})
	switch src.(type) {
	case []uint8:
		source = []byte(src.([]uint8))
	case string:
		source = []byte(src.(string))
	case nil:
		return nil
	default:
		return errors.New("incompatible type for string interface map")
	}
	err := json.Unmarshal(source, &mapt)
	if err != nil {
		return err
	}
	*o = JsonbType(mapt)
	return nil
}
