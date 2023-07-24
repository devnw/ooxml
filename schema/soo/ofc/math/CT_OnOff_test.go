package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_OnOffConstructor(t *testing.T) {
	v := math.NewCT_OnOff()
	if v == nil {
		t.Errorf("math.NewCT_OnOff must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_OnOff should validate: %s", err)
	}
}

func TestCT_OnOffMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_OnOff()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_OnOff()
	xml.Unmarshal(buf, v2)
}
