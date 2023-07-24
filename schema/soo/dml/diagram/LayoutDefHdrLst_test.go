package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestLayoutDefHdrLstConstructor(t *testing.T) {
	v := diagram.NewLayoutDefHdrLst()
	if v == nil {
		t.Errorf("diagram.NewLayoutDefHdrLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.LayoutDefHdrLst should validate: %s", err)
	}
}

func TestLayoutDefHdrLstMarshalUnmarshal(t *testing.T) {
	v := diagram.NewLayoutDefHdrLst()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewLayoutDefHdrLst()
	xml.Unmarshal(buf, v2)
}
