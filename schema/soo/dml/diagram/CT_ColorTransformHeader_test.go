package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ColorTransformHeaderConstructor(t *testing.T) {
	v := diagram.NewCT_ColorTransformHeader()
	if v == nil {
		t.Errorf("diagram.NewCT_ColorTransformHeader must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_ColorTransformHeader should validate: %s", err)
	}
}

func TestCT_ColorTransformHeaderMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_ColorTransformHeader()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_ColorTransformHeader()
	xml.Unmarshal(buf, v2)
}
