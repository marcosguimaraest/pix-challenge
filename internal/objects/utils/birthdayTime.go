package utils

import (
	"fmt"
	"strings"
	"time"
)

// Type created to unmarshal the time format from the ASAASApi
type BirthdayTime struct {
	time.Time
}

const birthdayLayout = "2006-01-02"

func (at *BirthdayTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		at = &BirthdayTime{}
		return
	}
	at.Time, err = time.Parse(birthdayLayout, s)
	if err != nil {
		return err
	}
	return
}

func (at BirthdayTime) MarshalJSON() ([]byte, error) {
	if at.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", at.Time.Format(birthdayLayout))), nil
}
