package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_XAlignConstructor(t *testing.T) {
	v := math.NewCT_XAlign()
	if v == nil {
		t.Errorf("math.NewCT_XAlign must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_XAlign should validate: %s", err)
	}
}

func TestCT_XAlignMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_XAlign()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_XAlign()
	xml.Unmarshal(buf, v2)
}
