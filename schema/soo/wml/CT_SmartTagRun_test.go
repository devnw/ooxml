package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SmartTagRunConstructor(t *testing.T) {
	v := wml.NewCT_SmartTagRun()
	if v == nil {
		t.Errorf("wml.NewCT_SmartTagRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SmartTagRun should validate: %s", err)
	}
}

func TestCT_SmartTagRunMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SmartTagRun()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SmartTagRun()
	xml.Unmarshal(buf, v2)
}
