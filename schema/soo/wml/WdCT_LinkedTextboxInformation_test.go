package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_LinkedTextboxInformationConstructor(t *testing.T) {
	v := wml.NewWdCT_LinkedTextboxInformation()
	if v == nil {
		t.Errorf("wml.NewWdCT_LinkedTextboxInformation must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_LinkedTextboxInformation should validate: %s", err)
	}
}

func TestWdCT_LinkedTextboxInformationMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_LinkedTextboxInformation()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_LinkedTextboxInformation()
	xml.Unmarshal(buf, v2)
}
