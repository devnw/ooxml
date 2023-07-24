package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestAG_IteratorAttributesConstructor(t *testing.T) {
	v := diagram.NewAG_IteratorAttributes()
	if v == nil {
		t.Errorf("diagram.NewAG_IteratorAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.AG_IteratorAttributes should validate: %s", err)
	}
}

func TestAG_IteratorAttributesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewAG_IteratorAttributes()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewAG_IteratorAttributes()
	xml.Unmarshal(buf, v2)
}
