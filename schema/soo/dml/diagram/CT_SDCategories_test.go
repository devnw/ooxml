package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_SDCategoriesConstructor(t *testing.T) {
	v := diagram.NewCT_SDCategories()
	if v == nil {
		t.Errorf("diagram.NewCT_SDCategories must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_SDCategories should validate: %s", err)
	}
}

func TestCT_SDCategoriesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_SDCategories()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_SDCategories()
	xml.Unmarshal(buf, v2)
}
