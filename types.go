package restful

import (
	"encoding/json"
	"fmt"
)

type NullString struct {
	Value string `json:"value,omitempty"`
	Valid bool `json:"-"`
}

func (s *NullString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Value)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		fmt.Println("UnmarshalJSON 1",data,s.Value,s.Valid)
		return nil
	}

	err := json.Unmarshal(data, &s.Value)
	if err != nil {
		fmt.Println("UnmarshalJSON 3",data,s.Value,s.Valid)
		return err
	}

	s.Valid = true

	fmt.Println("UnmarshalJSON 2",data,s.Value,s.Valid)

	return nil
}

type NullBool struct {
	Value bool
	Valid bool
}


type NullInt32 struct {
	Value int32
	Valid bool
}

type NullInt64 struct {
	Value int64
	Valid bool
}

type NullFloat32 struct {
	Value float32
	Valid bool
}

type NullFloat64 struct {
	Value float64
	Valid bool
}