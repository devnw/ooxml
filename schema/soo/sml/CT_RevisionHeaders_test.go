package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionHeadersConstructor(t *testing.T) {
	v := sml.NewCT_RevisionHeaders()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionHeaders must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionHeaders should validate: %s", err)
	}
}

func TestCT_RevisionHeadersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionHeaders()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionHeaders()
	xml.Unmarshal(buf, v2)
}
