package model

import (
	"fmt"
	"time"
)

type JsonTime struct {
	time.Time
}

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", this.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (this *JsonTime) Scan(src interface{}) error {
	if value, ok := src.(time.Time); ok {
		*this = JsonTime{value}
	}
	return nil
}
