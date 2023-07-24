package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextUnderlineFillGroupWrapperConstructor(t *testing.T) {
	v := dml.NewCT_TextUnderlineFillGroupWrapper()
	if v == nil {
		t.Errorf("dml.NewCT_TextUnderlineFillGroupWrapper must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextUnderlineFillGroupWrapper should validate: %s", err)
	}
}

func TestCT_TextUnderlineFillGroupWrapperMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextUnderlineFillGroupWrapper()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextUnderlineFillGroupWrapper()
	xml.Unmarshal(buf, v2)
}
