package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DeletedFieldConstructor(t *testing.T) {
	v := sml.NewCT_DeletedField()
	if v == nil {
		t.Errorf("sml.NewCT_DeletedField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DeletedField should validate: %s", err)
	}
}

func TestCT_DeletedFieldMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DeletedField()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DeletedField()
	xml.Unmarshal(buf, v2)
}
