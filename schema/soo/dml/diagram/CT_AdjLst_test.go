package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_AdjLstConstructor(t *testing.T) {
	v := diagram.NewCT_AdjLst()
	if v == nil {
		t.Errorf("diagram.NewCT_AdjLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_AdjLst should validate: %s", err)
	}
}

func TestCT_AdjLstMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_AdjLst()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_AdjLst()
	xml.Unmarshal(buf, v2)
}
