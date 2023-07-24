package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_BoxConstructor(t *testing.T) {
	v := math.NewCT_Box()
	if v == nil {
		t.Errorf("math.NewCT_Box must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Box should validate: %s", err)
	}
}

func TestCT_BoxMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Box()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Box()
	xml.Unmarshal(buf, v2)
}
