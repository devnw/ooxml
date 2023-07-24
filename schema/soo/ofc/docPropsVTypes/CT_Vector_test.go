package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestCT_VectorConstructor(t *testing.T) {
	v := docPropsVTypes.NewCT_Vector()
	if v == nil {
		t.Errorf("docPropsVTypes.NewCT_Vector must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.CT_Vector should validate: %s", err)
	}
}

func TestCT_VectorMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewCT_Vector()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewCT_Vector()
	xml.Unmarshal(buf, v2)
}
