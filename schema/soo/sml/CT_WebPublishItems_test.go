package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_WebPublishItemsConstructor(t *testing.T) {
	v := sml.NewCT_WebPublishItems()
	if v == nil {
		t.Errorf("sml.NewCT_WebPublishItems must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_WebPublishItems should validate: %s", err)
	}
}

func TestCT_WebPublishItemsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_WebPublishItems()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_WebPublishItems()
	xml.Unmarshal(buf, v2)
}
