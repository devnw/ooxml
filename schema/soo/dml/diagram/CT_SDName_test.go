package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_SDNameConstructor(t *testing.T) {
	v := diagram.NewCT_SDName()
	if v == nil {
		t.Errorf("diagram.NewCT_SDName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_SDName should validate: %s", err)
	}
}

func TestCT_SDNameMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_SDName()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_SDName()
	xml.Unmarshal(buf, v2)
}
