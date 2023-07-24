package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_BookViewConstructor(t *testing.T) {
	v := sml.NewCT_BookView()
	if v == nil {
		t.Errorf("sml.NewCT_BookView must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_BookView should validate: %s", err)
	}
}

func TestCT_BookViewMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_BookView()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_BookView()
	xml.Unmarshal(buf, v2)
}
