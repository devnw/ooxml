package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SignedTwipsMeasureConstructor(t *testing.T) {
	v := wml.NewCT_SignedTwipsMeasure()
	if v == nil {
		t.Errorf("wml.NewCT_SignedTwipsMeasure must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SignedTwipsMeasure should validate: %s", err)
	}
}

func TestCT_SignedTwipsMeasureMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SignedTwipsMeasure()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SignedTwipsMeasure()
	xml.Unmarshal(buf, v2)
}
