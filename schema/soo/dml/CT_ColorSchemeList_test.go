package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ColorSchemeListConstructor(t *testing.T) {
	v := dml.NewCT_ColorSchemeList()
	if v == nil {
		t.Errorf("dml.NewCT_ColorSchemeList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ColorSchemeList should validate: %s", err)
	}
}

func TestCT_ColorSchemeListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ColorSchemeList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ColorSchemeList()
	xml.Unmarshal(buf, v2)
}
