package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_PresentationOfConstructor(t *testing.T) {
	v := diagram.NewCT_PresentationOf()
	if v == nil {
		t.Errorf("diagram.NewCT_PresentationOf must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_PresentationOf should validate: %s", err)
	}
}

func TestCT_PresentationOfMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_PresentationOf()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_PresentationOf()
	xml.Unmarshal(buf, v2)
}
