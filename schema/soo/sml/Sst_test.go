package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestSstConstructor(t *testing.T) {
	v := sml.NewSst()
	if v == nil {
		t.Errorf("sml.NewSst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Sst should validate: %s", err)
	}
}

func TestSstMarshalUnmarshal(t *testing.T) {
	v := sml.NewSst()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewSst()
	xml.Unmarshal(buf, v2)
}
