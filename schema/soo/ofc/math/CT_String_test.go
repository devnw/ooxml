package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_StringConstructor(t *testing.T) {
	v := math.NewCT_String()
	if v == nil {
		t.Errorf("math.NewCT_String must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_String should validate: %s", err)
	}
}

func TestCT_StringMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_String()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_String()
	xml.Unmarshal(buf, v2)
}
