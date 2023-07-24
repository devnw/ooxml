package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ParameterConstructor(t *testing.T) {
	v := diagram.NewCT_Parameter()
	if v == nil {
		t.Errorf("diagram.NewCT_Parameter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Parameter should validate: %s", err)
	}
}

func TestCT_ParameterMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Parameter()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Parameter()
	xml.Unmarshal(buf, v2)
}
