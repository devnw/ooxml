package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SignedHpsMeasureConstructor(t *testing.T) {
	v := wml.NewCT_SignedHpsMeasure()
	if v == nil {
		t.Errorf("wml.NewCT_SignedHpsMeasure must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SignedHpsMeasure should validate: %s", err)
	}
}

func TestCT_SignedHpsMeasureMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SignedHpsMeasure()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SignedHpsMeasure()
	xml.Unmarshal(buf, v2)
}
