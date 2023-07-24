package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_TextPropsConstructor(t *testing.T) {
	v := diagram.NewCT_TextProps()
	if v == nil {
		t.Errorf("diagram.NewCT_TextProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_TextProps should validate: %s", err)
	}
}

func TestCT_TextPropsMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_TextProps()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_TextProps()
	xml.Unmarshal(buf, v2)
}
