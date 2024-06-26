package model

import (
	"encoding/json"
)

type UserID string

func (u UserID) String() string {
	return string(u)
}

type DialogID string

func (d DialogID) String() string {
	return string(d)
}

type Counter struct {
	DialogID DialogID
	Value    CounterValue
}

type CounterValue int

func (c *CounterValue) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *CounterValue) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}

	return nil
}
