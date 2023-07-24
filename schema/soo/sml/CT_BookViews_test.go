package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_BookViewsConstructor(t *testing.T) {
	v := sml.NewCT_BookViews()
	if v == nil {
		t.Errorf("sml.NewCT_BookViews must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_BookViews should validate: %s", err)
	}
}

func TestCT_BookViewsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_BookViews()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_BookViews()
	xml.Unmarshal(buf, v2)
}
