package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_HtmlPublishPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_HtmlPublishProperties()
	if v == nil {
		t.Errorf("pml.NewCT_HtmlPublishProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_HtmlPublishProperties should validate: %s", err)
	}
}

func TestCT_HtmlPublishPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_HtmlPublishProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_HtmlPublishProperties()
	xml.Unmarshal(buf, v2)
}
