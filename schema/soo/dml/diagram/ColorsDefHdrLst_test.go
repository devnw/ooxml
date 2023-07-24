package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestColorsDefHdrLstConstructor(t *testing.T) {
	v := diagram.NewColorsDefHdrLst()
	if v == nil {
		t.Errorf("diagram.NewColorsDefHdrLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.ColorsDefHdrLst should validate: %s", err)
	}
}

func TestColorsDefHdrLstMarshalUnmarshal(t *testing.T) {
	v := diagram.NewColorsDefHdrLst()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewColorsDefHdrLst()
	xml.Unmarshal(buf, v2)
}
