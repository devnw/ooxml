package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_CTCategoriesConstructor(t *testing.T) {
	v := diagram.NewCT_CTCategories()
	if v == nil {
		t.Errorf("diagram.NewCT_CTCategories must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_CTCategories should validate: %s", err)
	}
}

func TestCT_CTCategoriesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_CTCategories()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_CTCategories()
	xml.Unmarshal(buf, v2)
}
