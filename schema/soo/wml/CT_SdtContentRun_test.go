package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtContentRunConstructor(t *testing.T) {
	v := wml.NewCT_SdtContentRun()
	if v == nil {
		t.Errorf("wml.NewCT_SdtContentRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtContentRun should validate: %s", err)
	}
}

func TestCT_SdtContentRunMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtContentRun()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtContentRun()
	xml.Unmarshal(buf, v2)
}
