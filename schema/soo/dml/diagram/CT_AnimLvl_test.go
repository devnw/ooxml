package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_AnimLvlConstructor(t *testing.T) {
	v := diagram.NewCT_AnimLvl()
	if v == nil {
		t.Errorf("diagram.NewCT_AnimLvl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_AnimLvl should validate: %s", err)
	}
}

func TestCT_AnimLvlMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_AnimLvl()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_AnimLvl()
	xml.Unmarshal(buf, v2)
}
