package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_HierBranchStyleConstructor(t *testing.T) {
	v := diagram.NewCT_HierBranchStyle()
	if v == nil {
		t.Errorf("diagram.NewCT_HierBranchStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_HierBranchStyle should validate: %s", err)
	}
}

func TestCT_HierBranchStyleMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_HierBranchStyle()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_HierBranchStyle()
	xml.Unmarshal(buf, v2)
}
