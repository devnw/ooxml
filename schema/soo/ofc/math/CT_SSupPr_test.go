package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_SSupPrConstructor(t *testing.T) {
	v := math.NewCT_SSupPr()
	if v == nil {
		t.Errorf("math.NewCT_SSupPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_SSupPr should validate: %s", err)
	}
}

func TestCT_SSupPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_SSupPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_SSupPr()
	xml.Unmarshal(buf, v2)
}
