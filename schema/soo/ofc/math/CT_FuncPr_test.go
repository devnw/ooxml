package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_FuncPrConstructor(t *testing.T) {
	v := math.NewCT_FuncPr()
	if v == nil {
		t.Errorf("math.NewCT_FuncPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_FuncPr should validate: %s", err)
	}
}

func TestCT_FuncPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_FuncPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_FuncPr()
	xml.Unmarshal(buf, v2)
}
