package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_CxnConstructor(t *testing.T) {
	v := diagram.NewCT_Cxn()
	if v == nil {
		t.Errorf("diagram.NewCT_Cxn must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Cxn should validate: %s", err)
	}
}

func TestCT_CxnMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Cxn()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Cxn()
	xml.Unmarshal(buf, v2)
}
