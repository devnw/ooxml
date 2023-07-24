package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_AdjConstructor(t *testing.T) {
	v := diagram.NewCT_Adj()
	if v == nil {
		t.Errorf("diagram.NewCT_Adj must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Adj should validate: %s", err)
	}
}

func TestCT_AdjMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Adj()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Adj()
	xml.Unmarshal(buf, v2)
}
