package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_WholeE2oFormattingConstructor(t *testing.T) {
	v := dml.NewCT_WholeE2oFormatting()
	if v == nil {
		t.Errorf("dml.NewCT_WholeE2oFormatting must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_WholeE2oFormatting should validate: %s", err)
	}
}

func TestCT_WholeE2oFormattingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_WholeE2oFormatting()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_WholeE2oFormatting()
	xml.Unmarshal(buf, v2)
}
