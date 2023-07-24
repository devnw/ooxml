package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SingleXmlCellsConstructor(t *testing.T) {
	v := sml.NewCT_SingleXmlCells()
	if v == nil {
		t.Errorf("sml.NewCT_SingleXmlCells must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SingleXmlCells should validate: %s", err)
	}
}

func TestCT_SingleXmlCellsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SingleXmlCells()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SingleXmlCells()
	xml.Unmarshal(buf, v2)
}
