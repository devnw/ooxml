package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CustomWorkbookViewConstructor(t *testing.T) {
	v := sml.NewCT_CustomWorkbookView()
	if v == nil {
		t.Errorf("sml.NewCT_CustomWorkbookView must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomWorkbookView should validate: %s", err)
	}
}

func TestCT_CustomWorkbookViewMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomWorkbookView()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomWorkbookView()
	xml.Unmarshal(buf, v2)
}
