package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestEmptyConstructor(t *testing.T) {
	v := docPropsVTypes.NewEmpty()
	if v == nil {
		t.Errorf("docPropsVTypes.NewEmpty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Empty should validate: %s", err)
	}
}

func TestEmptyMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewEmpty()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewEmpty()
	xml.Unmarshal(buf, v2)
}
