package models

import (
	"database/sql/driver"
	"encoding/json"
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
	return json.Unmarshal(value.([]byte), c)
}
