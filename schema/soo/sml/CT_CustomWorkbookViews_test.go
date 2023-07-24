package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CustomWorkbookViewsConstructor(t *testing.T) {
	v := sml.NewCT_CustomWorkbookViews()
	if v == nil {
		t.Errorf("sml.NewCT_CustomWorkbookViews must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomWorkbookViews should validate: %s", err)
	}
}

func TestCT_CustomWorkbookViewsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomWorkbookViews()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomWorkbookViews()
	xml.Unmarshal(buf, v2)
}
