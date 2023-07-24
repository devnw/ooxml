package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestAG_ConstraintRefAttributesConstructor(t *testing.T) {
	v := diagram.NewAG_ConstraintRefAttributes()
	if v == nil {
		t.Errorf("diagram.NewAG_ConstraintRefAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.AG_ConstraintRefAttributes should validate: %s", err)
	}
}

func TestAG_ConstraintRefAttributesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewAG_ConstraintRefAttributes()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewAG_ConstraintRefAttributes()
	xml.Unmarshal(buf, v2)
}
