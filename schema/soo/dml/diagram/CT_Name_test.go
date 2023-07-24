package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_NameConstructor(t *testing.T) {
	v := diagram.NewCT_Name()
	if v == nil {
		t.Errorf("diagram.NewCT_Name must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Name should validate: %s", err)
	}
}

func TestCT_NameMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Name()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Name()
	xml.Unmarshal(buf, v2)
}
