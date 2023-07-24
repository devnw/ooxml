package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestColorsDefHdrConstructor(t *testing.T) {
	v := diagram.NewColorsDefHdr()
	if v == nil {
		t.Errorf("diagram.NewColorsDefHdr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.ColorsDefHdr should validate: %s", err)
	}
}

func TestColorsDefHdrMarshalUnmarshal(t *testing.T) {
	v := diagram.NewColorsDefHdr()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewColorsDefHdr()
	xml.Unmarshal(buf, v2)
}
