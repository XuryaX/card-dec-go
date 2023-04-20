package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type Cards []Card

func (c Cards) Value() (driver.Value, error) {
	if len(c) == 0 {
		return "[]", nil
	}
	return json.Marshal(c)
}

func (c *Cards) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var data []byte

	switch v := value.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return fmt.Errorf("unsupported type for Scan: %T", value)
	}

	return json.Unmarshal(data, c)
}
