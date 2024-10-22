package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Type created to unmarshal the time format from the ASAASApi
type AsaasTime struct {
	time.Time
}

var _ json.Unmarshaler = &AsaasTime{}
var _ json.Marshaler = &AsaasTime{}

const asaasLayout = "2006-01-02 15:04:05"

func (at *AsaasTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		at = &AsaasTime{}
		return
	}

	at.Time, err = time.Parse(asaasLayout, s)
	fmt.Println(at.Time)
	if err != nil {
		return err
	}
	return
}

func (at *AsaasTime) MarshalJSON() ([]byte, error) {
	if at.Time.IsZero() {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", at.Time.Format(asaasLayout))), nil
}
