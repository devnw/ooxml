package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_CategoryConstructor(t *testing.T) {
	v := diagram.NewCT_Category()
	if v == nil {
		t.Errorf("diagram.NewCT_Category must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Category should validate: %s", err)
	}
}

func TestCT_CategoryMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Category()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Category()
	xml.Unmarshal(buf, v2)
}
