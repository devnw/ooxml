package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_OfficeStyleSheetConstructor(t *testing.T) {
	v := dml.NewCT_OfficeStyleSheet()
	if v == nil {
		t.Errorf("dml.NewCT_OfficeStyleSheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_OfficeStyleSheet should validate: %s", err)
	}
}

func TestCT_OfficeStyleSheetMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_OfficeStyleSheet()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_OfficeStyleSheet()
	xml.Unmarshal(buf, v2)
}
