package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestVectorConstructor(t *testing.T) {
	v := docPropsVTypes.NewVector()
	if v == nil {
		t.Errorf("docPropsVTypes.NewVector must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Vector should validate: %s", err)
	}
}

func TestVectorMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewVector()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewVector()
	xml.Unmarshal(buf, v2)
}
