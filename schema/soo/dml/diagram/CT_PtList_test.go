package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_PtListConstructor(t *testing.T) {
	v := diagram.NewCT_PtList()
	if v == nil {
		t.Errorf("diagram.NewCT_PtList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_PtList should validate: %s", err)
	}
}

func TestCT_PtListMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_PtList()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_PtList()
	xml.Unmarshal(buf, v2)
}
