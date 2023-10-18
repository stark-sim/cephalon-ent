package common

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"entgo.io/ent/schema/field"
	"fmt"
)

type Bytes2StringSliceValueScanner struct{}

func (c Bytes2StringSliceValueScanner) Value(strings []string) (driver.Value, error) {
	if strings == nil {
		return nil, nil
	}
	b, err := json.Marshal(strings)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (c Bytes2StringSliceValueScanner) ScanValue() field.ValueScanner {
	return &sql.NullString{}
}

func (c Bytes2StringSliceValueScanner) FromValue(v driver.Value) ([]string, error) {
	s, ok := v.(*sql.NullString)
	if !ok {
		return nil, fmt.Errorf("unexpected input for FromValue: %T", v)
	}
	if !s.Valid {
		return nil, nil
	}
	var res []string
	err := json.Unmarshal([]byte(s.String), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
