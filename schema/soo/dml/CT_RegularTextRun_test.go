package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_RegularTextRunConstructor(t *testing.T) {
	v := dml.NewCT_RegularTextRun()
	if v == nil {
		t.Errorf("dml.NewCT_RegularTextRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_RegularTextRun should validate: %s", err)
	}
}

func TestCT_RegularTextRunMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_RegularTextRun()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_RegularTextRun()
	xml.Unmarshal(buf, v2)
}
