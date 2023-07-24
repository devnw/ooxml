package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtRunConstructor(t *testing.T) {
	v := wml.NewCT_SdtRun()
	if v == nil {
		t.Errorf("wml.NewCT_SdtRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtRun should validate: %s", err)
	}
}

func TestCT_SdtRunMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtRun()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtRun()
	xml.Unmarshal(buf, v2)
}
