package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_OMathParaPrConstructor(t *testing.T) {
	v := math.NewCT_OMathParaPr()
	if v == nil {
		t.Errorf("math.NewCT_OMathParaPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_OMathParaPr should validate: %s", err)
	}
}

func TestCT_OMathParaPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_OMathParaPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_OMathParaPr()
	xml.Unmarshal(buf, v2)
}
