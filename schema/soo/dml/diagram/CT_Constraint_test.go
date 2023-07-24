package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ConstraintConstructor(t *testing.T) {
	v := diagram.NewCT_Constraint()
	if v == nil {
		t.Errorf("diagram.NewCT_Constraint must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Constraint should validate: %s", err)
	}
}

func TestCT_ConstraintMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Constraint()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Constraint()
	xml.Unmarshal(buf, v2)
}
