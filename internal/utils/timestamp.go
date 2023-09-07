package utils

import (
	"fmt"
	"time"
)

type Timestamp struct {
	time.Time
}

func (ct *Timestamp) Scan(value interface{}) error {
	str_val, ok := value.(string)
	if !ok {
		return fmt.Errorf("Expected string value for Timestamp, got %T\n", value)
	}

	parsed_time, err := time.Parse("2006-01-02 15:04:05", str_val)
	if err != nil {
		return err
	}

	ct.Time = parsed_time
	return nil
}
