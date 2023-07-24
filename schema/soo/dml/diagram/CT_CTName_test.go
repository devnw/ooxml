package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_CTNameConstructor(t *testing.T) {
	v := diagram.NewCT_CTName()
	if v == nil {
		t.Errorf("diagram.NewCT_CTName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_CTName should validate: %s", err)
	}
}

func TestCT_CTNameMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_CTName()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_CTName()
	xml.Unmarshal(buf, v2)
}
