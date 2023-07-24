package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestStyleDefHdrLstConstructor(t *testing.T) {
	v := diagram.NewStyleDefHdrLst()
	if v == nil {
		t.Errorf("diagram.NewStyleDefHdrLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.StyleDefHdrLst should validate: %s", err)
	}
}

func TestStyleDefHdrLstMarshalUnmarshal(t *testing.T) {
	v := diagram.NewStyleDefHdrLst()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewStyleDefHdrLst()
	xml.Unmarshal(buf, v2)
}
