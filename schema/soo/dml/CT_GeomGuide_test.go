package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GeomGuideConstructor(t *testing.T) {
	v := dml.NewCT_GeomGuide()
	if v == nil {
		t.Errorf("dml.NewCT_GeomGuide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GeomGuide should validate: %s", err)
	}
}

func TestCT_GeomGuideMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GeomGuide()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GeomGuide()
	xml.Unmarshal(buf, v2)
}
