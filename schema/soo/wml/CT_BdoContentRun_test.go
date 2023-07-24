package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_BdoContentRunConstructor(t *testing.T) {
	v := wml.NewCT_BdoContentRun()
	if v == nil {
		t.Errorf("wml.NewCT_BdoContentRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_BdoContentRun should validate: %s", err)
	}
}

func TestCT_BdoContentRunMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_BdoContentRun()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_BdoContentRun()
	xml.Unmarshal(buf, v2)
}
