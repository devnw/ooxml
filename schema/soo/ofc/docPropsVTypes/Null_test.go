package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestNullConstructor(t *testing.T) {
	v := docPropsVTypes.NewNull()
	if v == nil {
		t.Errorf("docPropsVTypes.NewNull must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Null should validate: %s", err)
	}
}

func TestNullMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewNull()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewNull()
	xml.Unmarshal(buf, v2)
}
