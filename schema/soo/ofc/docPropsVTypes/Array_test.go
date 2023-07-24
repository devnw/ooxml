package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestArrayConstructor(t *testing.T) {
	v := docPropsVTypes.NewArray()
	if v == nil {
		t.Errorf("docPropsVTypes.NewArray must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Array should validate: %s", err)
	}
}

func TestArrayMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewArray()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewArray()
	xml.Unmarshal(buf, v2)
}
