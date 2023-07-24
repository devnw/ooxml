package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ColorTransformConstructor(t *testing.T) {
	v := diagram.NewCT_ColorTransform()
	if v == nil {
		t.Errorf("diagram.NewCT_ColorTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_ColorTransform should validate: %s", err)
	}
}

func TestCT_ColorTransformMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_ColorTransform()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_ColorTransform()
	xml.Unmarshal(buf, v2)
}
