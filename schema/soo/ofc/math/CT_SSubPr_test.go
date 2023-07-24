package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_SSubPrConstructor(t *testing.T) {
	v := math.NewCT_SSubPr()
	if v == nil {
		t.Errorf("math.NewCT_SSubPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_SSubPr should validate: %s", err)
	}
}

func TestCT_SSubPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_SSubPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_SSubPr()
	xml.Unmarshal(buf, v2)
}
