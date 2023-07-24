package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_OMathJcConstructor(t *testing.T) {
	v := math.NewCT_OMathJc()
	if v == nil {
		t.Errorf("math.NewCT_OMathJc must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_OMathJc should validate: %s", err)
	}
}

func TestCT_OMathJcMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_OMathJc()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_OMathJc()
	xml.Unmarshal(buf, v2)
}
