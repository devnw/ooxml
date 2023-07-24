package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionCustomViewConstructor(t *testing.T) {
	v := sml.NewCT_RevisionCustomView()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionCustomView must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionCustomView should validate: %s", err)
	}
}

func TestCT_RevisionCustomViewMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionCustomView()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionCustomView()
	xml.Unmarshal(buf, v2)
}
