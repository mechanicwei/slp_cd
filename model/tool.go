package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type JsonTime struct {
	time.Time
}

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	result := []byte("null")

	if !this.IsZero() {
		var stamp = fmt.Sprintf("\"%s\"", this.Format("2006-01-02 15:04:05"))
		result = []byte(stamp)
	}

	return result, nil
}

func (this *JsonTime) Scan(src interface{}) error {
	if value, ok := src.(time.Time); ok {
		*this = JsonTime{value}
	}
	return nil
}

func (this JsonTime) Value() (driver.Value, error) {
	return this.Format("2006-01-02 15:04:05"), nil
}

type NullString struct {
	sql.NullString
}

func (v *NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.String = *s
	} else {
		v.Valid = false
	}
	return nil
}
