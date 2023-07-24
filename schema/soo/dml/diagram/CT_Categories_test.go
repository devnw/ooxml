package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_CategoriesConstructor(t *testing.T) {
	v := diagram.NewCT_Categories()
	if v == nil {
		t.Errorf("diagram.NewCT_Categories must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Categories should validate: %s", err)
	}
}

func TestCT_CategoriesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Categories()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Categories()
	xml.Unmarshal(buf, v2)
}
