package jsonutil

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MustMarshal(t *testing.T, v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	return b
}

func MustUnmarshal(t *testing.T, data []byte, v any) {
	if err := json.Unmarshal(data, v); err != nil {
		assert.FailNow(t, err.Error())
	}
}
