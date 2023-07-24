package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ChildMaxConstructor(t *testing.T) {
	v := diagram.NewCT_ChildMax()
	if v == nil {
		t.Errorf("diagram.NewCT_ChildMax must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_ChildMax should validate: %s", err)
	}
}

func TestCT_ChildMaxMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_ChildMax()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_ChildMax()
	xml.Unmarshal(buf, v2)
}
