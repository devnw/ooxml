package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_AnimOneConstructor(t *testing.T) {
	v := diagram.NewCT_AnimOne()
	if v == nil {
		t.Errorf("diagram.NewCT_AnimOne must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_AnimOne should validate: %s", err)
	}
}

func TestCT_AnimOneMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_AnimOne()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_AnimOne()
	xml.Unmarshal(buf, v2)
}
