package model

import (
	"database/sql/driver"
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
