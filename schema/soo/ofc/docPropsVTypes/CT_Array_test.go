package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestCT_ArrayConstructor(t *testing.T) {
	v := docPropsVTypes.NewCT_Array()
	if v == nil {
		t.Errorf("docPropsVTypes.NewCT_Array must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.CT_Array should validate: %s", err)
	}
}

func TestCT_ArrayMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewCT_Array()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewCT_Array()
	xml.Unmarshal(buf, v2)
}
