package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_SSupConstructor(t *testing.T) {
	v := math.NewCT_SSup()
	if v == nil {
		t.Errorf("math.NewCT_SSup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_SSup should validate: %s", err)
	}
}

func TestCT_SSupMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_SSup()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_SSup()
	xml.Unmarshal(buf, v2)
}
