package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ThemeableLineStyleConstructor(t *testing.T) {
	v := dml.NewCT_ThemeableLineStyle()
	if v == nil {
		t.Errorf("dml.NewCT_ThemeableLineStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ThemeableLineStyle should validate: %s", err)
	}
}

func TestCT_ThemeableLineStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ThemeableLineStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ThemeableLineStyle()
	xml.Unmarshal(buf, v2)
}
