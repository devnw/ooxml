package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ChildPrefConstructor(t *testing.T) {
	v := diagram.NewCT_ChildPref()
	if v == nil {
		t.Errorf("diagram.NewCT_ChildPref must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_ChildPref should validate: %s", err)
	}
}

func TestCT_ChildPrefMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_ChildPref()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_ChildPref()
	xml.Unmarshal(buf, v2)
}
