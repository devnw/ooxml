package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ConstraintsConstructor(t *testing.T) {
	v := diagram.NewCT_Constraints()
	if v == nil {
		t.Errorf("diagram.NewCT_Constraints must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Constraints should validate: %s", err)
	}
}

func TestCT_ConstraintsMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Constraints()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Constraints()
	xml.Unmarshal(buf, v2)
}
