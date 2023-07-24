package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_BorderBoxPrConstructor(t *testing.T) {
	v := math.NewCT_BorderBoxPr()
	if v == nil {
		t.Errorf("math.NewCT_BorderBoxPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_BorderBoxPr should validate: %s", err)
	}
}

func TestCT_BorderBoxPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_BorderBoxPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_BorderBoxPr()
	xml.Unmarshal(buf, v2)
}
