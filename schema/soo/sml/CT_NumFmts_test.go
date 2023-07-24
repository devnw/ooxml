package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_NumFmtsConstructor(t *testing.T) {
	v := sml.NewCT_NumFmts()
	if v == nil {
		t.Errorf("sml.NewCT_NumFmts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_NumFmts should validate: %s", err)
	}
}

func TestCT_NumFmtsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_NumFmts()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_NumFmts()
	xml.Unmarshal(buf, v2)
}
