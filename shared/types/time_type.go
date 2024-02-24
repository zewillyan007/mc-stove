package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type TimeType time.Time

func NewTimeType(hour, min, sec int) TimeType {
	t := time.Date(0, time.January, 1, hour, min, sec, 0, time.UTC)
	return TimeType(t)
}

func (t *TimeType) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(string(v))
	case string:
		return t.UnmarshalText(v)
	case time.Time:
		*t = TimeType(v)
	case nil:
		*t = TimeType{}
	default:
		return fmt.Errorf("cannot sql.Scan() TimeType from: %#v", v)
	}
	return nil
}

func (t TimeType) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format("15:04:05")), nil
}

func (t *TimeType) UnmarshalText(value string) error {
	dd, err := time.Parse("15:04:05", value)
	if err != nil {
		return err
	}
	*t = TimeType(dd)
	return nil
}

func (TimeType) GormDataType() string {
	return "TIME"
}
