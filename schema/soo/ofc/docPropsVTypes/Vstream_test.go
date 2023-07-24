package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/docPropsVTypes"
)

func TestVstreamConstructor(t *testing.T) {
	v := docPropsVTypes.NewVstream()
	if v == nil {
		t.Errorf("docPropsVTypes.NewVstream must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Vstream should validate: %s", err)
	}
}

func TestVstreamMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewVstream()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewVstream()
	xml.Unmarshal(buf, v2)
}
