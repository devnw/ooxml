package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BackgroundFormattingConstructor(t *testing.T) {
	v := dml.NewCT_BackgroundFormatting()
	if v == nil {
		t.Errorf("dml.NewCT_BackgroundFormatting must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_BackgroundFormatting should validate: %s", err)
	}
}

func TestCT_BackgroundFormattingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_BackgroundFormatting()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_BackgroundFormatting()
	xml.Unmarshal(buf, v2)
}
