package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_SRgbColorConstructor(t *testing.T) {
	v := dml.NewCT_SRgbColor()
	if v == nil {
		t.Errorf("dml.NewCT_SRgbColor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_SRgbColor should validate: %s", err)
	}
}

func TestCT_SRgbColorMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_SRgbColor()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_SRgbColor()
	xml.Unmarshal(buf, v2)
}
