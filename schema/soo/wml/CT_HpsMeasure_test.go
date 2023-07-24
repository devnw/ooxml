package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_HpsMeasureConstructor(t *testing.T) {
	v := wml.NewCT_HpsMeasure()
	if v == nil {
		t.Errorf("wml.NewCT_HpsMeasure must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_HpsMeasure should validate: %s", err)
	}
}

func TestCT_HpsMeasureMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_HpsMeasure()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_HpsMeasure()
	xml.Unmarshal(buf, v2)
}
