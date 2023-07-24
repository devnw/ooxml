package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_SDCategoryConstructor(t *testing.T) {
	v := diagram.NewCT_SDCategory()
	if v == nil {
		t.Errorf("diagram.NewCT_SDCategory must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_SDCategory should validate: %s", err)
	}
}

func TestCT_SDCategoryMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_SDCategory()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_SDCategory()
	xml.Unmarshal(buf, v2)
}
