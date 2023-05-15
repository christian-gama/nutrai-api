package httputil

import (
	"encoding/json"
)

// Stringify converts any data to a string. It panics if it fails.
func Stringify(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(b)
}

// Json converts any data to a byte array. It panics if it fails.
func Json(fromData interface{}) []byte {
	b, err := json.Marshal(fromData)
	if err != nil {
		panic(err)
	}

	return b
}

// UnmarshalJson converts a byte array to any data. It panics if it fails.
func UnmarshalJson(data []byte, toData interface{}) {
	err := json.Unmarshal(data, toData)
	if err != nil {
		panic(err)
	}
}
