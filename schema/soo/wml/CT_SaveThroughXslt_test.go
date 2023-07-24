package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SaveThroughXsltConstructor(t *testing.T) {
	v := wml.NewCT_SaveThroughXslt()
	if v == nil {
		t.Errorf("wml.NewCT_SaveThroughXslt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SaveThroughXslt should validate: %s", err)
	}
}

func TestCT_SaveThroughXsltMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SaveThroughXslt()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SaveThroughXslt()
	xml.Unmarshal(buf, v2)
}
