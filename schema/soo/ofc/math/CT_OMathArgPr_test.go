package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_OMathArgPrConstructor(t *testing.T) {
	v := math.NewCT_OMathArgPr()
	if v == nil {
		t.Errorf("math.NewCT_OMathArgPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_OMathArgPr should validate: %s", err)
	}
}

func TestCT_OMathArgPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_OMathArgPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_OMathArgPr()
	xml.Unmarshal(buf, v2)
}
