package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_PtConstructor(t *testing.T) {
	v := diagram.NewCT_Pt()
	if v == nil {
		t.Errorf("diagram.NewCT_Pt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Pt should validate: %s", err)
	}
}

func TestCT_PtMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Pt()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Pt()
	xml.Unmarshal(buf, v2)
}
