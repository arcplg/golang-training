package model

import (
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

const customTimeLayout = "2006-01-02 15:04:05"

func (ct CustomTime) TimeValue() time.Time {
	return ct.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	// Remove quotes from the JSON string
	str = str[1 : len(str)-1]
	parsedTime, err := time.Parse(customTimeLayout, str)
	if err != nil {
		return fmt.Errorf("error parsing time: %v", err)
	}
	ct.Time = parsedTime
	return nil
}
