package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_FConstructor(t *testing.T) {
	v := math.NewCT_F()
	if v == nil {
		t.Errorf("math.NewCT_F must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_F should validate: %s", err)
	}
}

func TestCT_FMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_F()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_F()
	xml.Unmarshal(buf, v2)
}
