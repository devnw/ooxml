package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ChooseConstructor(t *testing.T) {
	v := diagram.NewCT_Choose()
	if v == nil {
		t.Errorf("diagram.NewCT_Choose must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Choose should validate: %s", err)
	}
}

func TestCT_ChooseMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Choose()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Choose()
	xml.Unmarshal(buf, v2)
}
