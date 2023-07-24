package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_CxnListConstructor(t *testing.T) {
	v := diagram.NewCT_CxnList()
	if v == nil {
		t.Errorf("diagram.NewCT_CxnList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_CxnList should validate: %s", err)
	}
}

func TestCT_CxnListMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_CxnList()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_CxnList()
	xml.Unmarshal(buf, v2)
}
