package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_RatioConstructor(t *testing.T) {
	v := dml.NewCT_Ratio()
	if v == nil {
		t.Errorf("dml.NewCT_Ratio must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Ratio should validate: %s", err)
	}
}

func TestCT_RatioMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Ratio()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Ratio()
	xml.Unmarshal(buf, v2)
}
