package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ClipboardStyleSheetConstructor(t *testing.T) {
	v := dml.NewCT_ClipboardStyleSheet()
	if v == nil {
		t.Errorf("dml.NewCT_ClipboardStyleSheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ClipboardStyleSheet should validate: %s", err)
	}
}

func TestCT_ClipboardStyleSheetMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ClipboardStyleSheet()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ClipboardStyleSheet()
	xml.Unmarshal(buf, v2)
}
