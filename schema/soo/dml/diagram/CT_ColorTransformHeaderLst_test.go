package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ColorTransformHeaderLstConstructor(t *testing.T) {
	v := diagram.NewCT_ColorTransformHeaderLst()
	if v == nil {
		t.Errorf("diagram.NewCT_ColorTransformHeaderLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_ColorTransformHeaderLst should validate: %s", err)
	}
}

func TestCT_ColorTransformHeaderLstMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_ColorTransformHeaderLst()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_ColorTransformHeaderLst()
	xml.Unmarshal(buf, v2)
}
