package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_FuncConstructor(t *testing.T) {
	v := math.NewCT_Func()
	if v == nil {
		t.Errorf("math.NewCT_Func must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Func should validate: %s", err)
	}
}

func TestCT_FuncMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Func()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Func()
	xml.Unmarshal(buf, v2)
}
