package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionHeaderConstructor(t *testing.T) {
	v := sml.NewCT_RevisionHeader()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionHeader must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionHeader should validate: %s", err)
	}
}

func TestCT_RevisionHeaderMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionHeader()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionHeader()
	xml.Unmarshal(buf, v2)
}
