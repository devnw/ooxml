package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_SSubSupPrConstructor(t *testing.T) {
	v := math.NewCT_SSubSupPr()
	if v == nil {
		t.Errorf("math.NewCT_SSubSupPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_SSubSupPr should validate: %s", err)
	}
}

func TestCT_SSubSupPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_SSubSupPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_SSubSupPr()
	xml.Unmarshal(buf, v2)
}
