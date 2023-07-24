package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestCT_VstreamConstructor(t *testing.T) {
	v := docPropsVTypes.NewCT_Vstream()
	if v == nil {
		t.Errorf("docPropsVTypes.NewCT_Vstream must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.CT_Vstream should validate: %s", err)
	}
}

func TestCT_VstreamMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewCT_Vstream()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewCT_Vstream()
	xml.Unmarshal(buf, v2)
}
