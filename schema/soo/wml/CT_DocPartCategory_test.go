package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartCategoryConstructor(t *testing.T) {
	v := wml.NewCT_DocPartCategory()
	if v == nil {
		t.Errorf("wml.NewCT_DocPartCategory must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocPartCategory should validate: %s", err)
	}
}

func TestCT_DocPartCategoryMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocPartCategory()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocPartCategory()
	xml.Unmarshal(buf, v2)
}
