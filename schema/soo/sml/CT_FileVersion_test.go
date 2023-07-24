package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FileVersionConstructor(t *testing.T) {
	v := sml.NewCT_FileVersion()
	if v == nil {
		t.Errorf("sml.NewCT_FileVersion must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FileVersion should validate: %s", err)
	}
}

func TestCT_FileVersionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FileVersion()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FileVersion()
	xml.Unmarshal(buf, v2)
}
