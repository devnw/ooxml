package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_DescriptionConstructor(t *testing.T) {
	v := diagram.NewCT_Description()
	if v == nil {
		t.Errorf("diagram.NewCT_Description must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Description should validate: %s", err)
	}
}

func TestCT_DescriptionMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Description()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Description()
	xml.Unmarshal(buf, v2)
}
