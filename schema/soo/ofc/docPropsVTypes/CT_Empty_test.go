package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestCT_EmptyConstructor(t *testing.T) {
	v := docPropsVTypes.NewCT_Empty()
	if v == nil {
		t.Errorf("docPropsVTypes.NewCT_Empty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.CT_Empty should validate: %s", err)
	}
}

func TestCT_EmptyMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewCT_Empty()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewCT_Empty()
	xml.Unmarshal(buf, v2)
}
