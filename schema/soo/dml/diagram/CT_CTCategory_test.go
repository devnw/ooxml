package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_CTCategoryConstructor(t *testing.T) {
	v := diagram.NewCT_CTCategory()
	if v == nil {
		t.Errorf("diagram.NewCT_CTCategory must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_CTCategory should validate: %s", err)
	}
}

func TestCT_CTCategoryMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_CTCategory()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_CTCategory()
	xml.Unmarshal(buf, v2)
}
