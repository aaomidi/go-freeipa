package freeipa

import (
	"encoding/json"
	"fmt"
	"testing"
)

type PrimJsonStruct struct {
	Test IPAPrim[string]
}

func TestPrimJson(t *testing.T) {
	data := `{test: ['a', 'b']}`
	resp := PrimJsonStruct{}
	err := json.Unmarshal([]byte(data), &PrimJsonStruct{})
	if err != nil {
		t.Fatal("error PrimJson: %v", err)
	}
	fmt.Println(resp)
}
