package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestStyleDefHdrConstructor(t *testing.T) {
	v := diagram.NewStyleDefHdr()
	if v == nil {
		t.Errorf("diagram.NewStyleDefHdr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.StyleDefHdr should validate: %s", err)
	}
}

func TestStyleDefHdrMarshalUnmarshal(t *testing.T) {
	v := diagram.NewStyleDefHdr()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewStyleDefHdr()
	xml.Unmarshal(buf, v2)
}
