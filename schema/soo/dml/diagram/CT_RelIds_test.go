package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_RelIdsConstructor(t *testing.T) {
	v := diagram.NewCT_RelIds()
	if v == nil {
		t.Errorf("diagram.NewCT_RelIds must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_RelIds should validate: %s", err)
	}
}

func TestCT_RelIdsMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_RelIds()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_RelIds()
	xml.Unmarshal(buf, v2)
}
