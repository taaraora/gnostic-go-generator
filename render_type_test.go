package main


import surface "github.com/googleapis/gnostic/surface"
import "testing"


func Test_JsonTag_positive(t *testing.T) {
	var field = &surface.Field{
		Name:          "DpaStatus",
		Serialize:     true,
	}

	res := JsonTag(field)

	if res != `json:"dpaStatus,omitempty"` {
		t.Fatalf("expected %v, got %v",`json:"dpaStatus,omitempty"`, res)
	}
}
