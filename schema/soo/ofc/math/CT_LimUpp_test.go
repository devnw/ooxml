package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_LimUppConstructor(t *testing.T) {
	v := math.NewCT_LimUpp()
	if v == nil {
		t.Errorf("math.NewCT_LimUpp must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_LimUpp should validate: %s", err)
	}
}

func TestCT_LimUppMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_LimUpp()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_LimUpp()
	xml.Unmarshal(buf, v2)
}
