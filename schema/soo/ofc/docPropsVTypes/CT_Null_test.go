package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestCT_NullConstructor(t *testing.T) {
	v := docPropsVTypes.NewCT_Null()
	if v == nil {
		t.Errorf("docPropsVTypes.NewCT_Null must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.CT_Null should validate: %s", err)
	}
}

func TestCT_NullMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewCT_Null()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewCT_Null()
	xml.Unmarshal(buf, v2)
}
