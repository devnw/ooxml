package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestLayoutDefHdrConstructor(t *testing.T) {
	v := diagram.NewLayoutDefHdr()
	if v == nil {
		t.Errorf("diagram.NewLayoutDefHdr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.LayoutDefHdr should validate: %s", err)
	}
}

func TestLayoutDefHdrMarshalUnmarshal(t *testing.T) {
	v := diagram.NewLayoutDefHdr()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewLayoutDefHdr()
	xml.Unmarshal(buf, v2)
}
